package controller

import (
	"devops/db"
	"devops/model"
	"encoding/json"
	"log"
	"net/http"
)

func GetList(w http.ResponseWriter, r *http.Request) {
	var serverList []model.Server
	rows, _ := db.MySql.Query("select name,host,password,username from t_server")
	defer rows.Close()
	for rows.Next() {
		var doc2 model.Server
		if err := rows.Scan(&doc2.Name, &doc2.Host, &doc2.Password, &doc2.Username); err != nil {
			log.Fatal(err)
		}
		serverList = append(serverList, doc2)
	}
	res, _ := json.Marshal(serverList)
	w.Write(res)
}
