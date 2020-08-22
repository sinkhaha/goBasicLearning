package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var employeeDB map[string]*Employee

// init在main之前执行，进行一些值的初始化
func init() {
	employeeDB = map[string]*Employee{}
	employeeDB["Mike"] = &Employee{"e-1", "Mike", 35}
	employeeDB["Rose"] = &Employee{"e-2", "Rose", 45}
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func GetEmployeeByName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	qName := ps.ByName("name")
	var (
		ok       bool
		info     *Employee
		infoJson []byte
		err      error
	)
	if info, ok = employeeDB[qName]; !ok {
		w.Write([]byte("{\"error\":\"Not Found\"}"))
		return
	}
	if infoJson, err = json.Marshal(info); err != nil {
		w.Write([]byte(fmt.Sprintf("{\"error\":,\"%s\"}", err)))
		return
	}

	w.Write(infoJson)
}

// 测试restful服务
func main() {
	router := httprouter.New()
	router.GET("/", Index)
	// 测试获取某个名字的信息
	router.GET("/employee/:name", GetEmployeeByName)

	log.Fatal(http.ListenAndServe(":8080", router))
}
