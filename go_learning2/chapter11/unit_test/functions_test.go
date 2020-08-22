package unit_test_testing

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert" // 断言库
)

// 测试t.Error方法
func TestErrorInCode(t *testing.T) {
	fmt.Println("start====")
	t.Error("error")
	// 会继续执行
	fmt.Println("end====")
}

// 测试t.Fatal方法
func TestFatalInCode(t *testing.T) {
	fmt.Println("start----")
	t.Fatal("fatal")
	// 不会继续执行
	fmt.Println("end----")
}

// 使用断言
func TestSquareWithAssert(t *testing.T) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		// 断言
		assert.Equal(t, expected[i], ret)
	}
}
