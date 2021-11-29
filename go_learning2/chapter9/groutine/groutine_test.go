package groutine

import (
	"fmt"
	"testing"
	"time"
)

// 测试协程
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// go后面加个方法表示协程，这里是10个协程，在里面打印i的值
		// 此处是个匿名方法，并调用方法
		go func(i int) {
			fmt.Println(i)
		}(i) // i传进函数，值传递
	}

	// 等待
	time.Sleep(time.Millisecond * 50)
}

func TestGroutine2(t *testing.T) {
	for i := 0; i < 10; i++ {
		// go后面加个方法表示协程，这里是10个协程，在里面打印i的值
		go test(i) // i传进函数，值传递
	}
	// 等待
	time.Sleep(time.Millisecond * 50)
}

func test(i int) {
	fmt.Println(i)
}
