package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "done"
}

func otherTask() {
	fmt.Println("haha")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task is done")
}

// 普通情况测试，同步执行，用于对比
func TestService(t *testing.T) {
	// 输出
	// done
	// haha
	// task is done
	fmt.Println(service())
	otherTask()
}

// 异步 channel阻塞和bufferchannel不阻塞
func AsyncService() chan string {
	// 声明一个channel，string表示这个channel只能放string类型的数据
	retCh := make(chan string)
	// 在协程里执行
	go func() {
		ret := service()
		fmt.Println("return result")
		// 当有结果了往channel里放数据,等接收者接收了才会执行service exit,
		// 所以这里会阻塞这个协程，改进方式，改为不阻塞
		// 声明channel时使用bufferchannle,此时容量为1，retCh := make(chan string, 1)
		retCh <- ret
		fmt.Println("service exit")
	}()
	return retCh
}

// 测试异步，等其他同步程序处理完了再去异步拿结果
func TestAsyncService(t *testing.T) {
	// 输出如下：
	// haha
	// return result
	// task is done
	// done
	// finally
	// service exit
	retCh := AsyncService() // 此时可能还没执行完成
	otherTask()

	// 往channel取数据并打印出来
	fmt.Println(<-retCh)

	fmt.Println("finally")
	time.Sleep(time.Second * 1)
}
