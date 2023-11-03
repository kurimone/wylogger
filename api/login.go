package api

import (
	"errors"
	"fmt"
	"xjtlu-dorm-net-auth-helper/request"
)

type LoginParams struct {
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReturns struct {
	ReplyCode int `json:"reply_code"`
	ReplyMsg  int `json:"reply_msg"`
}

func Login(params LoginParams) (LoginReturns, error) {
	var result LoginReturns
	err := request.Do("/login", "POST", params, &result)
	if err != nil {
		fmt.Println("[ERROR/API] Failed HTTP request:", err)
		return result, errors.New("failed HTTP reqeust")
	}
	fmt.Println("[DEBUG/API] HTTP request succeeded.")

	return result, nil
}
