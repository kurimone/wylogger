package api

import (
	"wylogger/logger"
	"wylogger/request"
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
	logger.Debug("Requesting API")
	err := request.Do("/login", "POST", params, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
