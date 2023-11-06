package main

import (
	"xjtlu-dorm-net-logger/auth"
	"xjtlu-dorm-net-logger/conf"
	"xjtlu-dorm-net-logger/logger"
)

func main() {
	logger.Init(logger.INFO)
	logger.Info("Logger loaded")

	err := conf.Load("config.yml")
	if err != nil {
		logger.Fatal("Failed to load profile \"config.yml\": %s", err)
	}

	if conf.Get().Debug {
		logger.SetLevel(logger.DEBUG)
	}
	logger.Debug("DEBUG MODE ENABLED, BE AWARE OF YOUR SAFETY!")
	logger.Debug("[ENV] URL = %s", conf.Get().URL)
	logger.Debug("[ENV] Domain = %s", conf.Get().Domain)
	logger.Debug("[ENV] Username = %s", conf.Get().Username)
	logger.Debug("[ENV] Password = %s", conf.Get().Password)

	logger.Info("Profile \"config.yml\" loaded")

	auth.Login()
}
