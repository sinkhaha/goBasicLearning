package singleton_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

// once.Do实现单例
func GetSingletonObj() *Singleton {
	// once.Do实现只执行一次的方法
	once.Do(func() {
		fmt.Println("create obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

// 单例
func TestGetSingletonObj(t *testing.T) {
	var wg sync.WaitGroup

	// 开启10个协程处理任务
	for i := 0; i < 10; i++ {
		wg.Add(1)

		go func() {
			obj := GetSingletonObj()

			// 输出对象的地址，10个协程都是同一个对象所以是同一个地址，如1268f78
			fmt.Printf("%x\n", unsafe.Pointer(obj))

			wg.Done()
		}()
	}

	wg.Wait()
}
