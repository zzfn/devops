package controller

import (
	"devops/db"
	"devops/model"
	"devops/util"
	"encoding/json"
	"log"
	"net/http"
)

func Download(w http.ResponseWriter, r *http.Request) {
	var projectList []model.Project
	rows, _ := db.MySql.Query("select name,repo from t_project")
	defer rows.Close()
	for rows.Next() {
		var projectItem model.Project
		if err := rows.Scan(&projectItem.Name, &projectItem.Repo); err != nil {
			log.Fatal(err)
		}
		projectList = append(projectList, projectItem)
	}
	util.Clone(projectList[0].Repo, projectList[0].Name)
	res, _ := json.Marshal(projectList)
	w.Write(res)
}

func Build(w http.ResponseWriter, r *http.Request) {
	util.Build()
}
func Send(w http.ResponseWriter, r *http.Request) {
	util.Send()
}
func Oss(w http.ResponseWriter, r *http.Request) {
	util.PutOss()
}
