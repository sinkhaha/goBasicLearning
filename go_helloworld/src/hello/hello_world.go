package main // 表示代码所在的包，执行入口是main包，文件名不一定是main.go

import (
	"fmt"
	"os"
)

// import (
// 	"fmt"
// 	"os"
// )

// 必须是main方法
func main() {
	// 获取命令行参数
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	} else {
		fmt.Println("hello world")
	}
	// 返回退出状态，不支持return
	os.Exit(-1)
}
