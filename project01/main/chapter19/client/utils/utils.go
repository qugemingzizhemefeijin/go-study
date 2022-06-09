package utils

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"go_code/project01/main/chapter19/common/message"
	"net"
)

//Transfer 这里将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte //这是传输时使用的缓冲
}

//ReadPkg ...
func (transfer *Transfer) ReadPkg() (mes message.Message, err error) {
	fmt.Println("读取客户端发送的数据...")
	//conn.Read 在conn没有被关闭的情况下，才会阻塞
	//如果客户端关闭了 conn 则，就不会阻塞了，会到处报错
	_, err = transfer.Conn.Read(transfer.Buf[:4])
	if err != nil {
		//err = errors.New("read pkg header error")
		return
	}
	//fmt.Println("读到的buf=", buf[:4])
	//根据buf[:4] 转成一个uint32类型
	var pkgLen uint32 = binary.BigEndian.Uint32(transfer.Buf[:4])

	//根据pkgLen读取消息内容
	n, err := transfer.Conn.Read(transfer.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		err = errors.New("read pkg body error")
		return
	}

	//把pkgLen 反序列化成 -> message.Message
	err = json.Unmarshal(transfer.Buf[:pkgLen], &mes)
	if err != nil {
		err = errors.New("read pkg json.Unmarshal error")
		return
	}

	return
}

//WritePkg ...
func (transfer *Transfer) WritePkg(data []byte) (err error) {
	//先放一个长度给对方
	var pkgLen uint32 = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(transfer.Buf[:4], pkgLen)

	// 发送长度
	n, err := transfer.Conn.Write(transfer.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail=", err)
		return
	}

	// 发送data本身
	n, err = transfer.Conn.Write(data)
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Write(data) fail=", err)
		return
	}

	return
}
