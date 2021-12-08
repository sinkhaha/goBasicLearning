package const_test

import "testing"

// 常量定义，第1种方式，对于连续的增量推荐使用此方式
const (
	Monday = 1 + iota // 初始值为1，后面变量的增量为1
	Tuesday
	Wednesday
)

// 第2，单独赋值
const (
	Thursday = 4
	Friday   = 5
	Saturday = 6
)

const (
	Readable = 1 << iota // 表示增量为向左移1
	Writable
	Executable
)

func TestConst(t *testing.T) {
	t.Log(Monday, Tuesday, Wednesday) // 1 2 3
	t.Log(Thursday, Friday, Saturday) // 4 5 6

	// 输出1 2 4  因为是01, 10, 100
	t.Log(Readable, Writable, Executable)
	a := 7 // 0111
	// &与运算，同为1才为1
	t.Log(Readable&a == Readable, Writable&a == Writable, Executable&a == Executable) // true true true
}
