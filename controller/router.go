package controller

import (
	"errors"

	"github.com/mumushuiding/go-simple-web-demo/model"
)

// RouteFunction 根据路径指向方法
type RouteFunction func(*model.Container) error

// RouterMap 路由
var RouterMap map[string]RouteFunction

// SetRouterMap 设置函数路由
func SetRouterMap() {
	RouterMap = make(map[string]RouteFunction)
	// RouterMap["visit/editor/flow"] = service.FindAllEditorFlowPaged

}

// GetRoute 获取执行函数
func GetRoute(route string) (func(*model.Container) error, error) {
	f := RouterMap[route]
	if f == nil {
		return nil, errors.New("method:" + route + ",不存在")
	}
	return f, nil
}
