package share_mem

import (
	"sync"
	"testing"
	"time"
)

// 共享内存
func TestShareMemory(t *testing.T) {
	counter := 0

	for i := 0; i < 5000; i++ {
		// 5000个协程操作同一个变量去自增
		go func() {
			counter++
		}()
	}

	// 强制等待1秒 防止程序执行完了 协程还没执行完导致结果错误
	time.Sleep(time.Second * 1)
	// 输出 counter = 4685，并不是5000，因为不是线程安全的
	t.Logf("counter = %d", counter)
}

// 线程安全的实现1：Mutex互斥锁
func TestShareMemorySafe(t *testing.T) {
	var mut sync.Mutex

	counter := 0
	for i := 0; i < 5000; i++ {
		go func() {
			// defer实现有错误时解锁
			defer func() {
				mut.Unlock()
			}()

			// 先加锁，然后再自增
			mut.Lock()
			counter++
		}()
	}

	// 强制等待1秒 防止程序执行完了，协程还没执行完导致结果错误
	time.Sleep(time.Second * 1)

	// 输出counter = 5000，是线程安全的
	t.Logf("counter = %d", counter)
}

// 线程安全的实现1：Mutex和WaitGroup
func TestCounterWaitGroup(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup

	counter := 0
	for i := 0; i < 5000; i++ {
		wg.Add(1) // 1、增加要等待的量，加一个等待的任务
		go func() {
			// defer实现有错误时解锁
			defer func() {
				mut.Unlock()
			}()

			// 2、加锁
			mut.Lock()
			counter++

			// 3、任务执行结束告诉wg任务结束
			wg.Done()
		}()
	}

	// 4、wait等待所有任务完成，会阻塞等待所有执行完才继续执行
	wg.Wait()

	// counter = 5000，是线程安全的
	t.Logf("counter = %d", counter)
}
