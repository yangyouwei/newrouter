package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/yangyouwei/newrouter/conf"
	"log"
)

var SqlDB *sql.DB

//初始化方法
func init() {
	var err error
	SqlDB, err = sql.Open("sqlite3", conf.SQLITE3)
	if err != nil {
		log.Fatal(err.Error())
	}

	//连接检测
	err = SqlDB.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
}