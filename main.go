package main

import (
	"devops/db"
	"devops/router"
)

func main() {
	db.MysqlInit()
	router.Router()
}
