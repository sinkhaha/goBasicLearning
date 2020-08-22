package channel_close

import (
	"testing"
	"fmt"
	"sync"
)

// 生产者
func dataProducer(ch chan int, wg *sync.WaitGroup) {
	// 固定放10个数据
    go func() {
        for i := 0; i < 10; i++ {
			ch <- i
		}
		// 当数据发送完了，关闭channel，通知接收者
		close(ch)
		// 往关闭的channel发消息，会发生panic错误
        // ch <- 11
		wg.Done()
	}()
}

// 接收者
// 接收者此时不清楚发送者里有几个消息，所以得通过判断channel是否关闭的机制来处理数据
func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
        for {
			// true表示正常接收数据，false表示通道关闭, 
			// 如果不判断通道是否关闭，此时从通道取到的数据是通道里数据类型的默认值，如此时是0
			if data,ok := <-ch; ok {
				fmt.Println(data)
			} else {
				fmt.Println("通道关闭了，退出")
				break
			}          
		}
		wg.Done()
	}()
}

// 测试发送者关闭channel通知接收者
func TestCloseChannel(t *testing.T) {
   var wg sync.WaitGroup
   // 声明一个channel
   ch := make(chan int)
   wg.Add(1)
   dataProducer(ch, &wg)

   // 此时开了两个接收者，两个接收者共同处理同一个channel的数据
   wg.Add(1)
   dataReceiver(ch, &wg)
   wg.Add(1)
   dataReceiver(ch, &wg)

   wg.Wait()
}
