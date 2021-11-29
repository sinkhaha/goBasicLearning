package groutine

import (
	"testing"
	"time"
)

// 测试协程
func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		// go表示协程,这里是10个协程，在里面打印i的值
		go func(i int) {
			t.Log(i)
		}(i) // i传进函数，值传递
	}

	// 等待
	time.Sleep(time.Millisecond * 50)
}
