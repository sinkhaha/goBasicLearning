package interface_test

import (
	"testing"
)

// go接口为非入侵性，实现不依赖于接口定义
// 所有接口的定义可以包含在接口使用者包内

// 接口定义 interface关键字
type Person interface {
	Eat() string
}

// 接口实现(类)
type Man struct {
}

// 实现接口的Eat方法（是结构体Man的一个方法）
func (m *Man) Eat() string {
	return "man吃东西"
}

// 测试接口
func TestInterfaceFunc(t *testing.T) {
	var p Person
	p = new(Man) // 实现和接口关联，不能用p = Man{}， 也可以用p = &Man{}  有点类似java的向上转型

	t.Log(p.Eat())
}
