package operator_test

import "testing"

// 测试数组直接用==比较
func TestOperator(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	// c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	e := [...]int{1, 3, 2, 4}
	// false
	t.Log(a == b)

	// 数组长度不一样，出现编译错误
	// t.Log(a == c)

	// true
	t.Log(a == d)
	
	// false 元素顺序不一样
	t.Log(a == e)
}

// 测试按位置零 &^ 运算符
func TestOpe(t *testing.T) {
	a := 7 // 0111
	a = a &^ 1 // 0111 &^ 01 等于 0110 即 6
	t.Log(a)
}