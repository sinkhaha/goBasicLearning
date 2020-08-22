package const_test

import "testing"

// 常量定义,第1种
const (
	Monday = 1 + iota // 初始值为1，后面变量的增量为1
	Tuesday
	Wednesday
)
// 第2种，也可以使用此类声明，不过对于连续的增量还是推荐上面的声明
// const (
// 	Monday = 1 
// 	Tuesday = 2
// 	Wednesday = 3
// )

// 测试常量
func TestConst1(t *testing.T) {
	// 输出1 2 3
	t.Log(Monday, Tuesday, Wednesday)
	
}

const (
	Readable = 1 << iota // 表示增量为向左移1, 
	Writable
	Executable
)
func TestConst2(t *testing.T) {
    // 输出1 2 4  (01, 10, 100)
	t.Log(Readable, Writable, Executable)
	a := 7 // 0111
	// &与运算，同为1才为1
	t.Log(Readable&a == Readable, Writable&a == Writable, Executable&a == Executable)
}