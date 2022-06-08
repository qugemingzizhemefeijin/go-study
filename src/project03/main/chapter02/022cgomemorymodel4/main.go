package main

/*
#include<stdio.h>


extern char* NewGoString(char*);
extern void FreeGoString(char*);
extern void PrintGoString(char*);

static void printString(const char* s) {
	char* gs = NewGoString(s);
	PrintGoString(gs);
	FreeGoString(gs);
}
 */
import "C"

import (
	"sync"
	"unsafe"
)

//export NewGoString
func NewGoString(s *C.char) *C.char {
	gs := C.GoString(s)
	id := NewObjectId(gs)
	return (*C.char)(unsafe.Pointer(uintptr(id)))
}

//export FreeGoString
func FreeGoString(p *C.char) {
	id := ObjectId(uintptr(unsafe.Pointer(p)))
	id.Free()
}

//export PrintGoString
func PrintGoString(p *C.char) {
	id := ObjectId(uintptr(unsafe.Pointer(p)))
	gs := id.Get().(string)
	print(gs)
}

// C⻓期持有Go指针对象
// Go语⾔禁⽌在C语⾔函数中⻓期持有Go指针对象， 但是这种需求是切实存在的。 如果需要在C语⾔中访问Go语⾔内存对象，
// 我们可以将Go语⾔内存对象在Go语⾔空间映射为⼀个int类型的id， 然后通过此id来间接访问和控制Go语⾔对象
// Go对象映射为整数类型的ObjectId， ⽤完之后需要⼿⼯调⽤free⽅法释放该对象ID

// 我们通过⼀个map来管理Go语⾔对象和id对象的映射关系。 其中NewObjectId⽤于创建⼀个和对象绑定的id，
// ⽽id对象的⽅法可⽤于解码出原始的Go对象， 也可以⽤于结束id和原始Go对象的绑定。
type ObjectId int32

var refs struct {
	sync.Mutex
	objs map[ObjectId]interface{}
	next ObjectId
}

func init() {
	refs.Lock()
	defer refs.Unlock()

	refs.objs = make(map[ObjectId]interface{})
	refs.next = 1000
}

func NewObjectId(obj interface{}) ObjectId {
	refs.Lock()
	defer refs.Unlock()

	id := refs.next
	refs.next++

	refs.objs[id] = obj
	return id
}

func (id ObjectId) IsNil() bool {
	return id == 0
}

func (id ObjectId) Get() interface{} {
	refs.Lock()
	defer refs.Unlock()

	return refs.objs[id]
}

func (id *ObjectId) Free() interface{} {
	refs.Lock()
	defer refs.Unlock()

	obj := refs.objs[*id]
	delete(refs.objs, *id)
	*id = 0

	return obj
}

func main() {
	// 在printString函数中， 我们通过NewGoString创建⼀个对应的Go字符串对象， 返回的其实是⼀个id， 不能直接使⽤。
	// 我们借助PrintGoString函数将id解析为Go语⾔字符串后打印。 该字符串在C语⾔函数中完全跨越了Go语⾔的内存管理，
	// 在PrintGoString调⽤前即使发⽣了栈伸缩导致的Go字符串地址发⽣变化也依然可以正常⼯作，
	// 因为该字符串对应的id是稳定的， 在Go语⾔空间通过id解码得到的字符串也就是有效的。

	// 此案例未通过编译
	C.printString("hello")
}
