package router

import (
	"net/http"

	"github.com/mumushuiding/go-simple-web-demo/controller"
)

// Mux 路由
var Mux = http.NewServeMux()

func init() {
	setMux()
}
func setMux() {
	Mux.HandleFunc("/api/v1/workdiary/index", controller.Index)
}
