package mysql

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
const (
	userName = "root"
	password = "123456"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "simple_chat"
)

// DB : 数据库实例
var DB *sql.DB

func init() {
	initDB()
}

func initDB() {
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8mb4&parseTime=true"}, "")
	var err error
	DB, err = sql.Open("mysql", path)
	if err != nil {
		panic(err)
	}
	DB.SetConnMaxLifetime(100)
	DB.SetConnMaxIdleTime(10)
	// 验证
	if err = DB.Ping(); err != nil {
		panic(err)
	}
	log.Println("database connected successfully!!!")
}
