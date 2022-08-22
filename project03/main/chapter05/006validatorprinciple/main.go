package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type Nested struct {
	Email string `validate:"email"`
}
type T struct {
	Age	int `validate:"eq=10"`
	Nested Nested
}

func validateEmail(input string) bool {
	if pass, _ := regexp.MatchString(
		`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, input,
	); pass {
		return true
	}
	return false
}

func validate(v interface{}) (bool, string) {
	validateResult := true
	errmsg := "success"
	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)
	for i := 0; i <vv.NumField(); i++ {
		fieldVal := vv.Field(i)
		tagContent := vt.Field(i).Tag.Get("validate")
		k := fieldVal.Kind()

		switch k {
		case reflect.Int:
			val := fieldVal.Int()
			tagValStr := strings.Split(tagContent, "=")
			tagVal, _ := strconv.ParseInt(tagValStr[1], 10, 64)
			if val != tagVal {
				errmsg = "validate int failed, tag is:"+ strconv.FormatInt(
					tagVal, 10,
				)
				validateResult = false
			}
		case reflect.String:
			val := fieldVal.String()
			tagValStr := tagContent
			switch tagValStr {
			case "email":
				nestedResult := validateEmail(val)
				if nestedResult == false {
					errmsg = "validate mail failed, field val is:"+ val
					validateResult = false
				}
			}
		case reflect.Struct:
			// 如果有内嵌的 struct，那么深度优先遍历
			// 就是一个递归过程
			valInter := fieldVal.Interface()
			nestedResult, msg := validate(valInter)
			if nestedResult == false {
				validateResult = false
				errmsg = msg
			}
		}
	}
	return validateResult, errmsg
}

// 这里我们简单地对 eq=x 和 email 这两个 tag 进行了支持
// 原理很简单，就是用反射对结构体进行树形遍历。
// 读者这时候可能会产生一个问题，我们对结构体进行校验时大量使用了反射，而 Go 的反射在性能上不太出众，有时甚至会影响到我们程序的性能。
// 这样的考虑确实有一些道理，但需要对结构体进行大量校验的场景往往出现在 Web 服务，这里并不一定是程序的性能瓶颈所在，
// 实际的效果还是要从 pprof 中做更精确的判断。

// 如果基于反射的校验真的成为了你服务的性能瓶颈怎么办？现在也有一种思路可以避免反射：
// 使用 Go 内置的 Parser 对源代码进行扫描，然后根据结构体的定义生成校验代码。我们可以将所有需要校验的结构体放在单独的包内。
func main() {
	var a = T{Age: 10, Nested: Nested{Email: "abc@abc.com"}}

	validateResult, errmsg := validate(a)
	fmt.Println(validateResult, errmsg)
}
