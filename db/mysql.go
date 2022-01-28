package db

import (
	"database/sql"
	"devops/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var MySql *sql.DB

func MysqlInit() {
	dsn := config.GetConfig().Mysql
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	MySql = db
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
}
