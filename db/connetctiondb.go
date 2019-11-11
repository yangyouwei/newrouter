package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var SqlDB *sql.DB

//初始化方法
func init() {
	var err error
	SqlDB, err = sql.Open("sqlite3", "./sqlite3.db")
	if err != nil {
		log.Fatal(err.Error())
	}

	//连接检测
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}