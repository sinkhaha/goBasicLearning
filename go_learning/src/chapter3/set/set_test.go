package set_test

import "testing"

// 测试set的初始化
func TestSetInit(t *testing.T) {
   // 方式一：中括号里面string是key的类型，后面int是值的类型
   m1 := map[string]int{"one":1, "two":2, "three":3}
   t.Log(len(m1)) // 3

}

