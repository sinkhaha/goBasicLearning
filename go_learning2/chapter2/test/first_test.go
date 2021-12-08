package first_test

import "testing"

// 单元测试，Test开头，Test后面跟的字符首字母要大写， t.Log替代fmt.Println打印
// 运行 go test -v执行，-v表示输出详细信息
func TestFirstTry(t *testing.T) {
	t.Log("my first test")
}
