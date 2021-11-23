package first_test

import "testing"

// 单元测试，Test开头，t.Log替代fmt.Println打印
// 运行 go test -v执行
func TestFirstTry(t *testing.T) {
	t.Log("my first try")
}
