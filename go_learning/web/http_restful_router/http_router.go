package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome! restful \n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// 获取路由中的参数值
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

// 使用第三方库实现restful服务
func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}
