package auth

import (
	"fmt"
	"time"
	"xjtlu-dorm-net-logger/api"
	"xjtlu-dorm-net-logger/conf"
	"xjtlu-dorm-net-logger/logger"
)

func Login() {
	maxRetries := 10
	retryDelay := 10 * time.Second
	retryTimer := time.NewTimer(0)
	defer retryTimer.Stop()

	for {
		<-retryTimer.C

		var err error
		for retries := 0; retries < maxRetries; retries++ {
			err = login()
			if err != nil {
				logger.Error("Failed to login: %s", err)
				logger.Info("Login attempt %d/%d failed", retries+1, maxRetries)
				if retries < maxRetries-1 {
					logger.Info("Retrying in %v...", retryDelay)
					time.Sleep(retryDelay)
				} else {
					logger.Info("Reached maximum number of retries (%d) for login", maxRetries)
					break
				}
			} else {
				break
			}
		}

		if err != nil {
			logger.Info("Will retry after 15 minutes")
			retryTimer.Reset(15 * time.Minute)
		} else {
			logger.Info("Will ensure to keep alive in 30 minutes")
			retryTimer.Reset(30 * time.Minute)
		}
	}
}

func login() error {
	params := api.LoginParams{
		Domain:   conf.Get().Domain,
		Username: conf.Get().Username,
		Password: conf.Get().Password,
	}

	result, err := api.Login(params)
	if err != nil {
		return err
	}

	if result.ReplyCode != 0 {
		return fmt.Errorf("exceptional API response: %d, %s", result.ReplyCode, result.ReplyMsg)
	}

	logger.Debug("API call succeeded: %#v", result)
	logger.Info("Login successful")
	return nil
}
