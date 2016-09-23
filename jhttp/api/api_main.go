package api

import (
	"SoftWeb/db/mongodb"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMainHandleMap() {
	router.routerMap["/"] = MainHandle
	router.routerMap["/{user}"] = UserHandle
}

func MainHandle(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the home !")
	tmp := []string{"ssss"}
	re, err := mongodb.GetDataConfig("1", "1", tmp)
	if err == nil {
		for _, key := range re {
			value, ok := re[key]
			if ok {
				fmt.Fprintf(w, value)
			}
		}
	}

}

func UserHandle(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	mongodb.SetDataConfig("1", "1", vars)
}
