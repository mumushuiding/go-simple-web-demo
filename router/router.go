package router

import (
	"net/http"

	"github.com/mumushuiding/go-simple-web-demo/config"

	"github.com/mumushuiding/go-simple-web-demo/controller"
)

// Mux 路由
var Mux = http.NewServeMux()
var conf = *config.Config

func crossOrigin(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", conf.AccessControlAllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", conf.AccessControlAllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", conf.AccessControlAllowHeaders)
		h(w, r)
	}
}
func init() {
	setMux()
}
func setMux() {
	Mux.HandleFunc("/api/v1/test/index", crossOrigin(controller.Index))
}
