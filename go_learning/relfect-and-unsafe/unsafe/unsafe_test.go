package unsafe_test

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type Customer struct {
	Name string
	Age  int
}

// 1. 测试不安全编程
func TestUnsafe(t *testing.T) {
	i := 10
	// 输出指针的值
	t.Log(unsafe.Pointer(&i)) // 0xc0000120e8
	// 强制转换成浮点值(非常危险，平常不推荐)
	f := *(*float64)(unsafe.Pointer(&i)) // 5e-323
	t.Log(f)
}

type MyInt int

// 2. 测试不安全编程进行合理的类型转换
func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b) // [1 2 3 4]
}

// 3. 测试原子类型操作，共享buffer安全到读写
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer

	// 写：写100个数到切片，然后用原子操作写到共享buffer到指针上
	writeDataFn := func() {
		data := []int{}
		for i := 0; i < 100; i++ {
			data = append(data, i)
		}
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}

	// 读：从buffer指针读取数据
	readDataFn := func() {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println(data, *(*[]int)(data))
	}

	var wg sync.WaitGroup
	writeDataFn()

	// 开启10次写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 10; i++ {
				writeDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()

		wg.Add(1)
		// 开启10次读操作
		go func() {
			for i := 0; i < 10; i++ {
				readDataFn()
				time.Sleep(time.Microsecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
