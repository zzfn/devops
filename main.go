package main

import (
	"devops/db"
	"devops/router"
	"devops/util"
)

func main() {
	util.InitBasePath()
	db.MysqlInit()
	router.Router()
}
