package channel_cancel

import (
	"testing"
	"fmt"
	"time"
)

// 判断是否取消了任务，当接收channel里的消息被阻塞时，会执行default，说明channel还没有被取消
func isCancelled(ch chan struct{}) bool {
    select {
	case <-ch:
		return true
	default:
		return false	
	}
}

// 关闭channel
func cancelChannel(ch chan struct{}) {
    close(ch)
}

// 测试取消协程里的任务
func TestCancleTack(t *testing.T) {
	// 通道里是放的是结构体数据
	cancelChan := make(chan struct{}, 0) 
	
	// 开启5个协程处理任务
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
            for {
				if isCancelled(cancelCh) {
                    break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "cancelled")
		}(i, cancelChan)
	}

	// 关闭channel
	cancelChannel(cancelChan)
	
	time.Sleep(time.Second * 1)
}
