package any_response

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(i int) string {
	time.Sleep(10 * time.Millisecond)
	return fmt.Sprintf("the result is from %d", i)
}

func FirstResponse() string {
	numberOfRunner := 10

	// 实例buffer channel
	ch := make(chan string, numberOfRunner)

	// 10个协程
	for i := 0; i < numberOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	// 当有任何一个任务完成后，就返回，即从channel取数据，由于是buffer channel，所以其他发送者不会阻塞
	return <-ch
}

// 测试仅需任意一个任务完成
func TestFirstResponse(t *testing.T) {
	// 协程数
	t.Log("before:", runtime.NumGoroutine()) // before: 2

	t.Log(FirstResponse()) // the result is from 8

	time.Sleep(1 * time.Second)

	t.Log("after:", runtime.NumGoroutine()) // after: 2
}
