package userinfo

import (
	"log"
	"net"
)

// GetUsrinfo : 获取用户信息
func GetUsrinfo(conn net.Conn) string {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if n < 1 || err != nil {
		log.Fatalln(err)
	}
	return string(buf[:n-1])
}
