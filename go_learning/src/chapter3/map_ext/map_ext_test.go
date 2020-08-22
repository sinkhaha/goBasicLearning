package map_ext_test

import "testing"

// map的值可以是一个方法
func TestMap(t *testing.T) {
	// 此时map的值是一个方法
	m1 := map[int]func(op int) int {}

	m1[1] = func(op int) int { return op }
	m1[2] = func(op int) int { return op*op }
	m1[3] = func(op int) int { return op*op*op }

	// 2 4 8
	t.Log(m1[1](2), m1[2](2), m1[3](2))
}

// map实现set
func TestMapForSet(t *testing.T) {
    mySet := map[int]bool{}
	mySet[1] = true
	n := 1
	// 判断元素1是否存在
	if mySet[n] {
        t.Logf("%d 存在", n)
	} else {
        t.Logf("%d 不存在", n)
	}

	mySet[3] = true
	t.Log(len(mySet)) // 2

	// 删除元素1
	delete(mySet, 1)
	if mySet[n] {
        t.Logf("%d 存在", n)
	} else {
        t.Logf("%d 不存在", n)
	}
}
