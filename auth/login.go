package auth

import (
	"errors"
	"xjtlu-dorm-net-auth-helper/api"
	"xjtlu-dorm-net-auth-helper/conf"
	"xjtlu-dorm-net-auth-helper/logger"
)

func Login() error {
	params := api.LoginParams{
		Domain:   conf.Get().Domain,
		Username: conf.Get().Username,
		Password: conf.Get().Password,
	}

	result, err := api.Login(params)
	if err != nil {
		logger.Error("Failed on API call: %s", err)
		return errors.New("failed on API call")
	}

	if result.ReplyCode != 0 {
		logger.Error("Exceptional API Returns: %d, %s", result.ReplyCode, result.ReplyMsg)
		return errors.New("exceptional API Returns")
	}

	logger.Debug("API call succeeded: %v", result)

	return nil
}
