package api

import "xjtlu-dorm-net-helper/request"

type LoginParams struct {
	Domain   string `schema:"domain"`
	Username string `schema:"username"`
	Password string `schema:"password"`
}

type LoginReturns struct {
	ReplyCode int `json:"reply_code"`
}

func Login(params LoginParams) error {
	var result LoginReturns
	err := request.Do("/login", "POST", params, &result)
	if err != nil {
		return err
	}

	return nil
}
