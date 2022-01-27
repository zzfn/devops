package controller

import (
	"devops/db"
	"fmt"
	"log"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.MySql.Query("select title from t_article")
	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			log.Fatal(err)
		}
		fmt.Println(title)
	}
	w.Write([]byte("dasd"))
}
