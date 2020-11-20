package userinfo

import "net"

// Client : 存储用户的连接和名称
type Client struct {
	Connection net.Conn // 用户 websocket 连接
	Name       string   // 用户名称
}
