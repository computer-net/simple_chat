package handler

import (
	"bufio"
	"fmt"
	"net"
	"server/userinfo"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
	users    []userinfo.Client
)

// Broadcaster : 广播给所有用户，同时控制用户的在线情况
func Broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// HandleChat : 聊天的句柄
func HandleChat(conn net.Conn) {
	defer conn.Close()
	// 读取用户名和密码
	username := userinfo.GetUsrinfo(conn)
	fmt.Println(username)
	password := userinfo.GetUsrinfo(conn)
	fmt.Println(password)
	users = append(users, userinfo.Client{
		conn,
		username,
	})
	// TODO: 验证用户名和密码
	ch := make(chan string)
	go clientWriter(conn, ch)
	ch <- "You are " + username
	messages <- username + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	// TODO: 设置关键字，实现私聊
	for input.Scan() {
		messages <- username + ": " + input.Text()
	}
	leaving <- ch
	messages <- username + " has left"
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
