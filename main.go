package main

import (
	"fmt"
	"xjtlu-dorm-net-auth-helper/auth"
	"xjtlu-dorm-net-auth-helper/conf"
)

func main() {
	var err error
	err = conf.Load("config.yml")
	if err != nil {
		panic(err)
	}
	fmt.Println("[DEBUG/MODE] DEBUG MODE ENABLED, BEWARE OF YOUR SECURITY!")
	fmt.Println("[DEBUG/ENV] URL =", conf.Get().URL)
	fmt.Println("[DEBUG/ENV] Domain =", conf.Get().Domain)
	fmt.Println("[DEBUG/ENV] Username =", conf.Get().Username)
	fmt.Println("[DEBUG/ENV] Password =", conf.Get().Password)
	fmt.Println("[INFO/ENV] Profile \"config.yml\" loaded.")

	err = auth.Login()
	if err != nil {
		fmt.Println("[ERROR/MAIN] Failed to login:", err)
	}
	println("[INFO/MAIN] Login Successful.")
}
