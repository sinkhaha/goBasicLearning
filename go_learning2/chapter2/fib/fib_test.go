package fib

import "testing"

// fibonacci数列实现
func TestFibList(t *testing.T) {
	// 第1种赋值方式
	var a int = 1
	var b int = 1
	// 第2种
	// var ( 
	// 	a int = 1  // int可以省略
	// 	b = 1
	// )
	// 第3种
	// a := 1
	// b := 1
	t.Log(a)
	// fmt.Print(a)
	for i:=0;i<5;i++ {
		t.Log(" ", b)
		// fmt.Print(" ", b)
		tmp := a
		a = b 
		b = tmp + a
	}
}

// 交换两个数的值
func TestExchangeTwo(t *testing.T) {
	a := 1
	b := 2
	// tmp := a
	// a = b
	// b = tmp
	// 可以在一个赋值语句中对多个变量进行同时赋值,不需要引入第3变量
	a,b = b,a
	t.Log(a, b);
}
