package all_response

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

func AllResponse() string {
	numberOfRunner := 10
	// 实例buffer channel
	ch := make(chan string, numberOfRunner)
	for i := 0; i < numberOfRunner; i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}

	// 从channel取所有结果相加后返回
	finalRet := ""
	for j := 0; j < numberOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

// 测试所有任务完成后才返回
func TestAllResponse(t *testing.T) {
	// 协程数
	t.Log("before:", runtime.NumGoroutine()) // 2

	t.Log(AllResponse())

	time.Sleep(1 * time.Second)

	t.Log("after:", runtime.NumGoroutine()) // 2
}
