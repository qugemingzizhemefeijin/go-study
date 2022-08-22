package main

import (
	"fmt"
	"github.com/go-playground/validator"
)

type RegisterReq struct {
	// 字符串的 gt=0 表示长度必须 > 0，gt = greater than
	Username       string   `validate:"gt=0"`
	// 同上
	PasswordNew    string   `validate:"gt=0"`
	// eqfield 跨字段相等校验
	PasswordRepeat string   `validate:"eqfield=PasswordNew"`
	// 合法 email 格式校验
	Email          string   `validate:"email"`
}

var validate = validator.New()

func validateFunc(req RegisterReq) error {
	err := validate.Struct(req)
	if err != nil {
		fmt.Println("validator error")
		return err
	}
	return nil
}

func main() {
	var req = RegisterReq {
		Username       : "Xargin",
		PasswordNew    : "ohno",
		PasswordRepeat : "ohn",
		Email          : "alex@abc.com",
	}

	err := validateFunc(req)
	fmt.Println(err)

	// Key: 'RegisterReq.PasswordRepeat' Error:Field validation for
	// 'PasswordRepeat' failed on the 'eqfield' tag
}