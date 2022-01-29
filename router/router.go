package router

import (
	"devops/controller"
	"net/http"
)

func Router() {
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("web"))))
	http.HandleFunc("/test", controller.GetList)
	http.HandleFunc("/repos", controller.Download)
	http.HandleFunc("/build", controller.Build)
	http.HandleFunc("/send", controller.Send)
	http.HandleFunc("/oss", controller.Oss)
	http.ListenAndServe(":8088", nil)
}
