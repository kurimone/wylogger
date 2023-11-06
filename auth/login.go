package auth

import (
	"time"
	"xjtlu-dorm-net-auth-helper/api"
	"xjtlu-dorm-net-auth-helper/conf"
	"xjtlu-dorm-net-auth-helper/logger"
)

func LoginD() {
	maxRetries := 1
	retryDelay := 10 * time.Second
	retryTimer := time.NewTimer(0)
	defer retryTimer.Stop()

	for {
		<-retryTimer.C

		var err error
		for retries := 0; retries < maxRetries; retries++ {
			err = Login()
			if err != nil {
				logger.Error("Failed to login: %s", err)
				logger.Error("Login attempt %d/%d failed.", retries+1, maxRetries)
				if retries < maxRetries-1 {
					logger.Info("Retrying in %v...", retryDelay)
					time.Sleep(retryDelay)
				} else {
					logger.Error("Reached maximum number of retries (%d) for login", maxRetries)
					break
				}
			} else {
				break
			}
		}

		if err != nil {
			logger.Info("will retry after 15 minutes")
			retryTimer.Reset(15 * time.Minute)
		} else {
			logger.Info("Will ensure to keep alive in 30 minutes")
			retryTimer.Reset(30 * time.Minute)
		}
	}
}

func Login() error {
	params := api.LoginParams{
		Domain:   conf.Get().Domain,
		Username: conf.Get().Username,
		Password: conf.Get().Password,
	}

	result, err := api.Login(params)
	if err != nil {
		logger.Error("Failed on API call: %s", err)
		return err
	}

	if result.ReplyCode != 0 {
		logger.Error("Exceptional API response: %d, %s", result.ReplyCode, result.ReplyMsg)
		return err
	}

	logger.Debug("API call succeeded: %#v", result)
	logger.Info("Login successful")
	return nil
}
