package sync_pool_obj_cache

import (
	"fmt"
	"sync"
	"runtime"
	"testing"
)

// 测试sync.pool缓存对象的获取和放回
func TestSyncPool(t *testing.T) {
	// 创建实例化一个池
	pool := &sync.Pool{
		// 如果所有子池是空的，就用此自定义的方法产生一个新的对象返回，并不会存在池中，需要用put放入
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 100
		},
	}

	// 获取对象，此时池是空的，所以会先new创建，返回100
	v := pool.Get().(int)
	fmt.Println("获取到的对象v是", v)

	// 放回对象
	pool.Put(3)
	// 此时池里面有值3了，所以不会去创建了,返回3
	v1, _ := pool.Get().(int)
	fmt.Println("获取到的对象v1是", v1)

	// 测试gc
	// 先往池中放一个4，然后执行GC清除sync.pool中缓存的对象,gc时间不确定，可能出现还没gc掉池的对象
	pool.Put(4)
	runtime.GC()
	// 此时池是空的，所以会去new创建,返回100
	v2, _ := pool.Get().(int)
	fmt.Println("获取到的对象v2是", v2)
}


// 测试多个协程操作对象池
func TestSyncPoolInMultiGroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Create a new object.")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		// 开10个协程获取对象
		wg.Add(1)
		go func(id int) {
			fmt.Println("获取到的对象是", pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}