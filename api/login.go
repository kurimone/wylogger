package api

import (
	"errors"
	"xjtlu-dorm-net-auth-helper/logger"
	"xjtlu-dorm-net-auth-helper/request"
)

type LoginParams struct {
	Domain   string `json:"domain"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginReturns struct {
	ReplyCode int    `json:"reply_code"`
	ReplyMsg  string `json:"reply_msg"`
}

func Login(params LoginParams) (LoginReturns, error) {
	var result LoginReturns
	err := request.Do("/login", "POST", params, &result)
	if err != nil {
		logger.Error("Failed HTTP request: %s", err)
		return result, errors.New("failed HTTP reqeust")
	}
	logger.Debug("HTTP request succeeded")

	return result, nil
}
