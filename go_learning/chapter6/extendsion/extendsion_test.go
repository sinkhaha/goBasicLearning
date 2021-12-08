package extendsion_test

import (
	"fmt"
	"testing"
)

// 相当于父类
type Pet struct {
}

// 父类的方法
func (p *Pet) Speak() {
	fmt.Print("...")
}
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

// 实现组合
type Dog struct {
	p *Pet // 扩展Pet的功能，此处是指针类型
}

func (d *Dog) Speak() {
	// 组合，调的Pet的方法
	d.p.Speak() // ...
}

func (d *Dog) SpeakTo(host string) {
	// 组合,调的Pet的方法
	d.p.SpeakTo(host) // ... dog
}

// 测试组合复用
func TestDog(t *testing.T) {
	dog := new(Dog)
	dog.Speak()        // ...
	dog.SpeakTo("dog") // ... dog
}

/** =======================================*/
// 虽然看起来类似继承，但是子类的方法不能覆写父类的方法，子类型无法转换成父类型
type Cat struct {
	Pet // 内嵌的形式
}

func (c *Cat) Speak() {
	fmt.Println("miao")
}

// func (c *Cat) SpeakTo(a string) {
// 	fmt.Println("miao SpeakTo")
// }

// 测试继承
func TestCat(t *testing.T) {
	cat := new(Cat)

	// cat实例可以调Pet的SpeakTo方法，类似继承，如果cat有SpeakTo，则调cat自己的
	cat.SpeakTo("cat") // ... cat,注意不是miao cat, SpeakTo里的Speak调的还是Pet父类的方法
}
