package main

import (
	"xjtlu-dorm-net-auth-helper/auth"
	"xjtlu-dorm-net-auth-helper/conf"
	"xjtlu-dorm-net-auth-helper/logger"
)

func main() {
	logger.Init(logger.INFO)
	logger.Info("Logger loaded")

	var err error

	err = conf.Load("config.yml")
	if err != nil {
		logger.Fatal("Failed to load profile \"config.yml\"")
	}

	if conf.Get().Debug {
		logger.SetLevel(logger.DEBUG)
	}
	logger.Debug("DEBUG MODE ENABLED, MAKE SURE YOU WANT IT!")
	logger.Debug("[ENV] URL = %s", conf.Get().URL)
	logger.Debug("[ENV] Domain = %s", conf.Get().Domain)
	logger.Debug("[ENV] Username = %s", conf.Get().Username)
	logger.Debug("[ENV] Password = %s", conf.Get().Password)

	logger.Info("Profile \"config.yml\" loaded")

	err = auth.Login()
	if err != nil {
		logger.Error("Failed to login: %s", err)
		return
	}
	logger.Info("Login successful")
}
