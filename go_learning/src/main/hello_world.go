package main // 包，表示代码所在的包，执行入口是main包，不是文件名，即文件名的main可以修改

import "fmt"
import "os"

// import (
// 	"fmt"
// 	"os"
// )

// 执行入口main方法
func main() {
	// 获取命令行参数
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		fmt.Println("hello world", os.Args[1])
	}

	// fmt.Println("hello world")
	// 返回退出状态，不支持return
	os.Exit(-1)
}
