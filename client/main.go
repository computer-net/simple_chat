package main

import (
	"client/handler"
	"fmt"
)

func main() {
	var order byte
	fmt.Println("Register(0) or Log in(1)?")
	fmt.Scanf("%d", &order)
	switch order {
	case 0:
		username, password := getUserinfo()
		handler.Register(username, password)
	case 1:
		username, password := getUserinfo()
		handler.Login(username, password)
	}
}

func getUserinfo() (username string, password string) {
	fmt.Println("Please input username:")
	fmt.Scanln(&username)
	fmt.Println("Please input password:")
	fmt.Scanln(&password)
	return username, password
}
