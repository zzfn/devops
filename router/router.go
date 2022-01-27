package router

import (
	"devops/controller"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controller.GetList)
}
