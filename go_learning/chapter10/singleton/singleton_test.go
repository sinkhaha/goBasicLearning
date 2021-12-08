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

// 实例化
func GetSingletonObj() *Singleton {
	// 只执行一次的方法once.Do()
	once.Do(func() {
		fmt.Println("create obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

// 测试单例
func TestGetSingletonObje(t *testing.T) {
	var wg sync.WaitGroup
	// 开启10个协程处理任务
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj()
			// 输出对象的地址，都是同一个对象所以是同一个地址，如1268f78
			fmt.Printf("%x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}
