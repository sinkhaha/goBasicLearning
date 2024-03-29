package main

import (
	"fmt"
	"net/http" // 内置http模块
	"time"
)

// 启动http服务 go run hello_http.go
// 请求http://localhost:8080/和http://localhost:8080/time/
// 注意访问http://localhost:8080/time/111也会走到/time/路由
func main() {
	// 定义请求路由
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World! go go go")
	})

	// 注意time是以/结尾，需要了解go默认的路由规则
	http.HandleFunc("/time/", func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		timeStr := fmt.Sprintf("{\"time\": \"%s\"}", t)
		w.Write([]byte(timeStr))
	})

	// 监听8080端口
	http.ListenAndServe(":8080", nil)
}
