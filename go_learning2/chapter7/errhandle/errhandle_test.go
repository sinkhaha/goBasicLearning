package errhandle_test

import (
	"errors"
	"fmt"
	"testing"
)

// 测试panic和os.Exit的区别
func TestPanicVsExit(t *testing.T) {
	defer func() {
		fmt.Println("panic发生错误了，defer也会执行")
	}()
	fmt.Println("start")
	// 会输出具体错误栈 panic退出前会执行defer中的代码
	// panic(errors.New("手动错误"))

	// 不会输出具体错误栈 退出时不会调用defer指定的函数
	// 	os.Exit(1)
}

// 测试recover捕获错误，类似try{}catch(){}
func TestRecover(t *testing.T) {
	// defer和recover必须在panic抛出异常之前定义
	// panic异常会随着函数的调用栈向外传递
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover捕获了错误", err)
		}
	}()
	fmt.Println("start")

	// 会输出具体错误栈
	errHandle()

	t.Log(333)
}

func errHandle() {
	panic(errors.New("出错了"))
}
