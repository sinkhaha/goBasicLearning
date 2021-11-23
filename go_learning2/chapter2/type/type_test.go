package type_change

import "testing"

// 为int64类型起别名
type MyInt int64

// 测试go不支持隐式类型转换
func TestTypeChange(t *testing.T) {
	var a int32 = 1
	var b int64
	// 错误，不支持隐式类型转换
	// b = a
	// 正确，显示类型转换
	b = int64(a)

	var c MyInt
	// 错误，别名也不支持隐式类型转换
	// c = b
	// 正确,显示类型转化
	c = MyInt(b)

	t.Log(a, b, c)
}

// go不支持指针运算
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a // &a获取a的地址
	// 错误，不支持指针运算
	// aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr) // 格式化获取变量的类型， %T 获取变量的类型
}

// go把string初始化为空字符串
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*") //**
	t.Log(len(s)) // 0
	if s == "" {
		t.Log("string初始值为空字符串")
	}
}