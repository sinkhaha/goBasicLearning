package select_test

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	// 500比100毫秒长，实现超时
	time.Sleep(time.Millisecond * 500)
	return "done"
}

// 异步 channel阻塞
func AsyncService() chan string {
	// 声明一个channel，string表示这个channel只能放string类型的数据
	retCh := make(chan string)
	// 在协程里执行
	go func() {
		ret := service()
		fmt.Println("return result")
		retCh <- ret
		fmt.Println("service exit")
	}()
	return retCh
}

// 测试select实现 超时机制
func TestSelect(t *testing.T) {
	// case后面跟的是一种阻塞事件，case不能保证执行顺序，哪个先准备好处于非阻塞状态则哪个先执行
	select {
	// 从channel等待一个消息
	case ret := <-AsyncService():
		t.Log(ret)
	// 超时100ms，比上面的500ms短，所以会先执行
	case <-time.After(time.Millisecond * 100):
		t.Error("time out")
		// 当执行select时其他事件都没准备好，此时会直接执行default，没有default则阻塞在case事件中
		// default:
		//	t.Error("都没准备好")
	}
}
