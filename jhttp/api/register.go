package api

import "net/http"

type HandleFunc func(http.ResponseWriter, *http.Request)

var router struct {
	routerMap map[string]HandleFunc
}

func GetRouterMap() (result map[string]HandleFunc) {
	return router.routerMap
}

func init() {
	router.routerMap = map[string]HandleFunc{}
	GetMainHandleMap()
}
