package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xujiajun/godbal"
	"github.com/xujiajun/godbal/driver/mysql"
)

func main() {
	database, err := godbal.NewMysql("root:123@tcp(127.0.0.1:3306)/test?charset=utf8").Open()

	if err != nil {
		panic(err)
	}

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	queryBuilder := mysql.NewQueryBuilder(database)

	sql := queryBuilder.Insert("userinfo").Value("username", "johnny").Value("departname", "tec").Value("created", "1521010136").GetSQL()

	fmt.Print(sql)

	lastInsertId, err := queryBuilder.PrepareAndExecute()

	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(lastInsertId)
}
