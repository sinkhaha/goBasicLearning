package obj_pool

import (
	"fmt"
	"testing"
	"time"
)

// 测试对象池
func TestObjPool(t *testing.T) {
	// 创建10个channel的对象池
	pool := NewObjPool(10)

	// 尝试放置超出池大小的对象
	// if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	// t.Error(err)
	// }

	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			fmt.Printf("对象类型是 %T\n", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}

	fmt.Println("Done")
}
