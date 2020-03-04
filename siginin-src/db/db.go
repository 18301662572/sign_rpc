package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//数据库实例

var (
	db *sqlx.DB
)

//Init初始化
func InitDB(dsn string) (err error) {
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect sqlx failed,err:%v\n", err)
		return err
	}
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	return nil
}
