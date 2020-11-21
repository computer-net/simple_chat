package userinfo

import (
	"bufio"
	"log"
	"net"
	"strings"
)

// GetUserinfo : 获取用户信息
func GetUserinfo(conn net.Conn) (string, string) {
	reader := bufio.NewReader(conn)
	content, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	info := strings.Split(content, "\t")
	return info[0], info[1]
}
