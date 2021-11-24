package map_test

import "testing"

// 测试map的声明和初始化
func TestInitMap(t *testing.T) {
	// 方式一：中括号里面string是key的类型，后面int是值的类型
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])                 // 4
	t.Logf("len m1=%d", len(m1)) // 3

	// 方式二
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2)) //1

	// 方式三 10表示容量为10，map不能用cap函数
	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3)) // 0
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1]) // 默认初始化为0

	m1[2] = 0

	t.Log(m1[2]) // 值本来就是0，所以直接判断是否等于0是无法判断存不存在2这个key的
	m1[3] = 0

	// 判断元素是否存在的正确方法
	if v, ok := m1[3]; ok {
		t.Logf("Key 3's value is %d", v)
	} else {
		t.Log("key 3 is not existing.")
	}
}

// map的遍历
func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	for k, v := range m1 {
		t.Log(k, v)
	}
}
