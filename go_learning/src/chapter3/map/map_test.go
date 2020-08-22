package map_test

import "testing"

// 测试map的初始化
func TestMapInit(t *testing.T) {
   // 方式一：中括号里面string是key的类型，后面int是值的类型
   m1 := map[string]int{"one":1, "two":2, "three":3}
   t.Log(len(m1)) // 3

   // 方式二：
   m2 := map[string]int{}
   t.Log(len(m2)) // 0
   m2["1"] = 6
   t.Log(m2["1"], len(m2)) // 6 1

   // 方式三
   m3 := make(map[int]int, 10) // 10是cap，不是len
   t.Log(len(m3)) // 0
}

// 判断map的元素是否存在
func TestExistKey(t *testing.T) {
	a := map[int]int{}
	t.Log(a[1]) // 0  此时不存在key为1，但是是返回了0，而不是返回空

	a[2] = 0
	t.Log(a[2]) // 0 此时手动设置元素的值为0，也是返回0

	// 问题：怎么判断元素是否存在呢（不能通过是否返回nil来判断元素是否存在）
    if v,ok := a[3];ok {
		t.Log("key 3 存在", v)
	} else {
		t.Log("key 3 不存在", v)
	}
}

// map的遍历
func TestMap(t *testing.T) {
	m1 := map[int]int{ 1:1, 2:2, 3:3 }
	// 遍历的key是无顺序的
	for key,val := range m1 {
		t.Log(key, val)
	} 
}

