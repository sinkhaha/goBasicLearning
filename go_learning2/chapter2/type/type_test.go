package type_change

import "testing"

// 为int64类型起别名
type MyInt int64

// 测试go不支持隐式类型转换
func TestTypeChange(t *testing.T) {
	var a int32 = 1
	var b int64 // 默认值0

	// 错误，不支持隐式类型转换
	// b = a
	// 正确，支持显示类型转换
	b = int64(a)

	var c MyInt
	// 错误，别名也不支持隐式类型转换
	// c = b
	// 正确
	c = MyInt(b)

	t.Log(a, b, c)
}

// go不支持指针运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a          // &a获取a的地址，类似c++
	var aPtr2 *int = &a // 等价于 aPtr2 := &a

	// 错误，不支持指针运算
	// aPtr = aPtr + 1

	t.Log(a, aPtr, aPtr2)              // 如 1 0xc00011c198 0xc00011c198
	t.Logf("%T %T %T", a, aPtr, aPtr2) // %T是获取变量的类型，输出int *int *int
}

// string默认初始化为空字符串
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*") //**
	t.Log(len(s))        // 0

	if s == "" {
		t.Log("string初始值为空字符串")
	}
}
