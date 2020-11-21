package mysql

import (
	"fmt"
	"log"
	"strconv"
)

// Insert : 插入数据
func Insert(username string, password string) {
	insert_query := fmt.Sprintf(
		"insert into userinfo (username, password) values(‘%s‘,‘%s’)",
		username,
		password,
	)
	log.Println("sql:", insert_query)
	rows, err := DB.Exec(insert_query)
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
