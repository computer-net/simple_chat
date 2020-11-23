package handler

import (
	"bufio"
	"fmt"
	"net"
	"server/userinfo"
	"strings"
	"sync"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
	users    sync.Map
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
	username, password := userinfo.GetUserinfo(conn)
	fmt.Println("username: " + username)
	fmt.Println("password: " + password)
	// 单点登录
	if _, ok := users.Load(username); ok {
		fmt.Fprintln(conn, username+"is online!")
		conn.Close()
		return
	} else { // TODO: 验证用户名和密码
		users.Store(username, userinfo.Client{
			conn,
			password,
		})
	}
	ch := make(chan string)
	go clientWriter(conn, ch)
	ch <- "You are " + username
	messages <- username + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)
	// TODO: 设置关键字，实现私聊
	for input.Scan() {
		content := input.Text()
		if strings.HasPrefix(content, "@@") { // 命令模式，私聊
			res := strings.SplitN(content, " ", 2)
			if v, ok := users.Load(res[0][2:]); ok {
				fmt.Println(len(content))
				if len(content) == 2 {
					fmt.Fprintln(v.(userinfo.Client).Connection, content[1])
				} else {
					fmt.Fprintln(v.(userinfo.Client).Connection, username+"提到了你")
				}
			}
		} else if strings.HasPrefix(content, "@q") { // 命令模式，断开连接
			break
		} else { // 正常的群聊
			messages <- username + ": " + content
		}
	}
	leaving <- ch
	messages <- username + " has left"
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
