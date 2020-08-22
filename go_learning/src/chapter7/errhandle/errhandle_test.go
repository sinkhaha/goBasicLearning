package errhandle_test

import (
	"testing"
	"errors"
	"fmt"
)

// 测试panic和os.Exit的区别
func TestPanicVsExit(t *testing.T) {
	defer func() {
		fmt.Println("panic发生错误了，defer也会执行")
	}()
	fmt.Println("start")
	// 会输出具体错误栈
	// panic(errors.New("手动错误"))
	
	// 不会输出具体错误栈
    // os.Exit(1)
}

// 测试recover捕获错误，类似try{}catch(){}
func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover捕获了错误", err)
		}
	}()
	fmt.Println("start")
	// 会输出具体错误栈
	panic(errors.New("出错了"))
}
