package main

import "xjtlu-dorm-net-helper/auth"

func main() {
	err := auth.Login()
	if err != nil {
		panic(err)
	}
}
