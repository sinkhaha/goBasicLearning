package duotai_test

import (
	"fmt"
	"testing"
)

// 接口定义 interface关键字
type Person interface {
	Eat() string
}

// 接口实现
type Man struct {
}

// 结构体的Eat方法
func (m *Man) Eat() string {
	return "man吃东西"
}

// 接口实现
type WoMan struct {
}

// 结构体的Eat方法
func (m *WoMan) Eat() string {
	return "woman吃东西"
}

func Eat(p Person) {
	fmt.Println("哈哈哈", p.Eat())
}

// 测试多态
func TestInterfaceFunc(t *testing.T) {
	man1 := new(Man) // 实现和接口关联
	t.Log(&man1)
	woman2 := new(WoMan)
	t.Log(&woman2)

	Eat(man1) // 传参时就进行了和接口关联
	Eat(woman2)
}
