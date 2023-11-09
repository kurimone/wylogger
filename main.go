package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"wylogger/auth"
	"wylogger/conf"
	"wylogger/logger"
)

var (
	version      = "dev"
	commit       = "none"
	buildDate    = "unknown"
	confPath     = getDefaultConfPath()
	printVersion = false
)

func main() {
	parseFlag()

	if printVersion {
		fmt.Printf("wylogger %s (%s %s)", version, commit, buildDate)
		return
	}

	logger.Init(logger.INFO)
	logger.Info("Logger loaded")

	err := conf.Load(confPath)
	if err != nil {
		logger.Fatal("Failed to load configuration from %s: %v", confPath, err)
	}
	logger.Info("Configuration loaded from %s", confPath)

	if conf.Get().Debug {
		setDebug()
	}

	auth.Login()
}

func getDefaultConfPath() string {
	e, _ := os.Executable()

	if runtime.GOOS == "windows" {
		return filepath.Dir(e) + "/config.yml"
	}
	return "/etc/wylogger/config.yml"
}

func parseFlag() {
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
