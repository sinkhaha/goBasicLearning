package csp

import (
	"fmt"
	"testing"
	"time"
)

// 睡眠函数
func service() string {
	time.Sleep(time.Millisecond * 50)
	return "done"
}

func otherTask() {
	fmt.Println("haha")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("task is done")
}

// 测试普通同步执行的情况，用于对比
func TestService(t *testing.T) {
	// 输出
	// done
	// haha
	// task is done
	fmt.Println(service())
	otherTask()
}

// 异步 channel阻塞、bufferchannel不阻塞
func AsyncService() chan string {
	// 声明一个channel，string表示这个channel只能放string类型的数据
	retCh := make(chan string)

	// 开启协程执行service方法
	go func() {
		ret := service()

		fmt.Println("return result")

		// 当有结果了往channel里放数据，等接收者接收了才会执行输出“service exit”
		// 因为是channel，所以这里会阻塞这个协程，可以改为不阻塞的，声明channel时使用bufferchannel，此时容量为1，retCh := make(chan string, 1)
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

	// 往channel取数据并打印出来，打印done，然后协程会继续执行
	fmt.Println(<-retCh) // done

	fmt.Println("finally")
	time.Sleep(time.Second * 1)
}
