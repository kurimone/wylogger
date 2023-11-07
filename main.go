package main

import (
	"flag"
	"fmt"
	"runtime"
	"wylogger/auth"
	"wylogger/conf"
	"wylogger/logger"
)

const version = "0.1.5"

var (
	confPath     string
	printVersion bool
)

func main() {
	praseFlag()

	if printVersion {
		fmt.Println("Version:", version)
		return
	}

	logger.Init(logger.INFO)
	logger.Info("Logger loaded")

	err := conf.Load(confPath)
	if err != nil {
		logger.Fatal("Failed to load configuration \"%s\": %s", confPath, err)
	}
	logger.Info("Configuration \"%s\" loaded", confPath)

	if conf.Get().Debug {
		setDebug()
	}

	auth.Login()
}

func praseFlag() {
	flag.StringVar(&confPath, "c", getDefaultConfPath(), "Path to config file")
	flag.StringVar(&confPath, "config", getDefaultConfPath(), "Path to config file")
	flag.BoolVar(&printVersion, "v", false, "Show version")
	flag.BoolVar(&printVersion, "version", false, "Show version")

	flag.Parse()
}

func setDebug() {
	logger.SetLevel(logger.DEBUG)
	logger.Debug("DEBUG MODE ENABLED, BE AWARE OF YOUR SAFETY!")
	logger.Debug("[ENV] URL = %s", conf.Get().URL)
	logger.Debug("[ENV] Domain = %s", conf.Get().Domain)
	logger.Debug("[ENV] Username = %s", conf.Get().Username)
	logger.Debug("[ENV] Password = %s", conf.Get().Password)
}

func getDefaultConfPath() string {
	if runtime.GOOS == "windows" {
		return "config.yml"
	}
	return "/etc/wylogger/config.yml"
}
