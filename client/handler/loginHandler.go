package handler

import (
	"io"
	"log"
	"net"
	"os"
)

// Login : For each user to login
func Login(username string, password string) {
	// 建立 TCP 连接
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// TODO: 验证用户名和密码
	conn.Write([]byte(username + "\t" + password + "\n"))
	// 开始聊天
	startChat(conn)
}

func startChat(conn net.Conn) {
	done := make(chan struct{})
	defer conn.Close()
	go func() {
		mustCopy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	go mustCopy(conn, os.Stdin)
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
