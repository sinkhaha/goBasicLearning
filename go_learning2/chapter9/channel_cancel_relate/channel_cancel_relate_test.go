package channel_cancel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCancelled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

// 测试用context来取消关联任务
func TestCancleTack(t *testing.T) {
	// 创建根context
	partentCtx := context.Background()
	// 创建子context
	childCtx, cancel := context.WithCancel(partentCtx)

	// 开启5个协程处理任务
	for i := 0; i < 5; i++ {
		go func(i int, childCtx context.Context) {
			for {
				if isCancelled(childCtx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "cancelled")
		}(i, childCtx)
	}

	// 取消任务
	cancel()

	time.Sleep(time.Second * 1)
}
