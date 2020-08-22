package interface_test

import (
	"testing"
)

// 接口定义 interface关键字
type Person interface {
	Eat() string
}

// 接口实现
type Man struct {

}

// 实现接口的Eat方法
func (m *Man) Eat() string {
	return "man吃东西"
}

// 测试接口
func TestInterfaceFunc(t *testing.T) {
	var p Person
	p = new(Man) // 实现和接口关联，不能用p = Man{}, 也可以用p = &Man{}

    t.Log(p.Eat())
}
