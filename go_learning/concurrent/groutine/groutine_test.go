package groutine

import (
	"fmt"
	"testing"
	"time"
)

// 协程，go后面加方法表示协程
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// 循环产生10个协程，此处是个匿名方法，并调用方法，在里面打印i的值
		go func(i int) {
			fmt.Println(i)
		}(i) // i传进函数，值传递
	}

	// 等待
	time.Sleep(time.Millisecond * 50)
}

func TestGroutine2(t *testing.T) {
	for i := 0; i < 10; i++ {
		go test(i)
	}
	time.Sleep(time.Millisecond * 50)
}

func test(i int) {
	fmt.Println(i)
}
