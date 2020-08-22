package duotai_test

import (
	"testing"
	"fmt"
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

// 接口实现
type WoMan struct {

}

// 实现接口的Eat方法
func (m *WoMan) Eat() string {
	return "woman吃东西"
}

func Eat(p Person) {
    fmt.Println("哈哈哈", p.Eat());
}

// 测试多态
func TestInterfaceFunc(t *testing.T) {
	man1 := new(Man) // 实现和接口关联
	woman2 := new(WoMan)

	Eat(man1) // 传参时就进行了和接口关联
	Eat(woman2)
}
