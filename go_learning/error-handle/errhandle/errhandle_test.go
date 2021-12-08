package errhandle_test

import (
	"errors"
	"fmt"
	"testing"
)

// panic和os.Exit的区别？
// panic用于不可恢复的错误；panic退出前会执行defer方法内的代码
// os.Exit退出时不会调用defer指定的函数；os.Exit退出时不会输出当前调用栈的信息，panic则相反

func TestPanicVsExit(t *testing.T) {
	defer func() {
		fmt.Println("panic发生错误了，defer也会执行")
	}()

	fmt.Println("start")

	// 会输出具体错误栈，panic退出前会执行defer中的代码
	panic(errors.New("手动错误"))
	// 如果改成Exit，则不会输出具体错误栈 退出时不会调用defer指定的函数
	// os.Exit(1)
}

// 测试recover捕获错误，类似try{}catch(){}
func TestRecover(t *testing.T) {
	// defer和recover必须在panic抛出异常之前定义
	// panic异常会随着函数的调用栈向外传递
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover捕获了错误", err) // recover捕获了错误 出错了
		}
	}()

	fmt.Println("start")

	// 发生panic
	errHappen()

	t.Log(333) // 不会执行到
}

func errHappen() {
	panic(errors.New("出错了"))
}

// 例如数组下标越界，空指针异常时，系统就会调用自己本身的panic函数
func TestArrayRangeErr(t *testing.T) {
	i := 10
	testArrayRange(i)
	fmt.Println("hello world")
}

func testArrayRange(i int) {
	var arr [10]int
	// 优先使用错误拦截 在错误出现之前进行拦截 在错误出现后进行错误捕获
	// 错误拦截必须配合defer使用 通过匿名函数使用
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err) // runtime error: index out of range [10] with length 10
		}
	}()

	arr[i] = 123     // err panic
	fmt.Println(arr) // 前面出错了，这里不会执行到
}
