package auth

import (
	"xjtlu-dorm-net-helper/api"
	"xjtlu-dorm-net-helper/conf"
)

func Login() error {
	params := api.LoginParams{
		Domain:   conf.Get().Domain,
		Username: conf.Get().Username,
		Password: conf.Get().Password,
	}

	err := api.Login(params)
	if err != nil {
		return err
	}

	return nil
}
