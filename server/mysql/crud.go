package mysql

import (
	"log"
	"strconv"
)

// Insert : 插入数据
func Insert(username string, password string) {
	insertQuery := "insert into userinfo (username, password) values(?,?)"
	log.Println("sql:", insertQuery)
	rows, err := DB.Exec(insertQuery, username, password)
	checkErr(err)
	num, err := rows.RowsAffected()
	checkErr(err)
	log.Println("Effected rows' number: " + strconv.Itoa(int(num)))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
