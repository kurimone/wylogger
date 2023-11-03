package auth

import (
	"errors"
	"fmt"
	"xjtlu-dorm-net-auth-helper/api"
	"xjtlu-dorm-net-auth-helper/conf"
)

func Login() error {
	params := api.LoginParams{
		Domain:   conf.Get().Domain,
		Username: conf.Get().Username,
		Password: conf.Get().Password,
	}

	result, err := api.Login(params)
	if err != nil {
		fmt.Println("[ERROR/AUTH] Failed on API call:", err)
		return errors.New("failed on API call")
	}

	if result.ReplyCode != 0 {
		fmt.Println("[ERROR/AUTH] Exceptional API Returns:", result.ReplyCode, result.ReplyMsg)
		return errors.New("exceptional API Returns")
	}

	fmt.Println("[DEBUG/AUTH] API call succeeded.", result)

	return nil
}
