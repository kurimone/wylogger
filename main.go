package main

import (
	"fmt"
	"xjtlu-dorm-net-helper/auth"
	"xjtlu-dorm-net-helper/conf"
)

func main() {
	var err error
	err = conf.Load("config.yml")
	if err != nil {
		panic(err)
	}

	fmt.Println("[ENV] URL = ", conf.Get().URL)
	fmt.Println("[ENV] Domain = ", conf.Get().Domain)
	fmt.Println("[ENV] Username = ", conf.Get().Username)
	fmt.Println("[ENV] Password = ", conf.Get().Password)

	err = auth.Login()
	if err != nil {
		panic(err)
	}
}
