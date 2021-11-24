package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 测试函数可以有2个返回值
func TestFn(t *testing.T) {
	a, b := ReturnMutilValues()
	t.Log(a, b)

	// 技巧，下划线_忽略返回值
	c, _ := ReturnMutilValues()
	t.Log(c)
}

// 函数的返回值有2个，所以返回类型用括号括起来，返回1个则不用
func ReturnMutilValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

// 一个计算函数耗时的函数
// inner func (op int) int传入一个函数，这个函数是返回int型
// 代表func (op int) int这个函数是返回一个函数
func TimeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent", time.Since(start).Seconds())
		return ret
	}
}

// 函数定义
// (1) func 函数名(形参 形参类型) 返回值类型 {}
// (2) func 函数名(形参 形参类型) (返回值1类型 返回值2类型) {}
func SlowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 测试函数的参数为函数功能
func TestInnerFun(t *testing.T) {
	tsSF := TimeSpent(SlowFun)
	t.Log(tsSF(10))
}

// 测试可变参数，相当于一个数组
func TestChangeFun(t *testing.T) {
	a := Sum(1, 2, 3)
	b := Sum(1, 2, 3, 4)
	t.Log(a, b) // 6,10
}

// 可变参数 注意省略号是跟在类型前，而不是变量前
func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

// 测试defer函数（可以用于释放一些锁什么的资源）
func TestDeferFun(t *testing.T) {
	// 在函数执行结束后才会执行，即使抛错了也会执行
	defer Clear()

	fmt.Println("start")
	panic("手动抛错误")     // panic内置
	fmt.Println("end") // 前面报错了，此句不会执行
}

func Clear() {
	fmt.Println("释放资源")
}
