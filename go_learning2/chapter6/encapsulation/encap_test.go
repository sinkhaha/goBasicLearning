package encapsulation_test

import (
	"testing"
	"fmt"
)

// 结构体定义
type Employee struct {
	Id string // 不需要逗号
	Name string
	Age int
}

// 测试实例化结构体
func TestCreateEmplyeeObj(t *testing.T) {
	// 实例方式一
	e1 := Employee{"0", "张三", 20}
	t.Log(e1) // {0 张三 20}

	// 实例方式二
	e2 := Employee{Name: "李四", Age: 30}
	t.Log(e2) // { 李四 30}

	// 实例方式三
	e3 := new(Employee) // 这里返回引用，相当于e := &Employee{}
	t.Log(e3) // &{  0}
	e3.Id = "4"
	e3.Name = "老五"
	e3.Age = 40
	t.Log(e3) // &{4 老五 40}

	t.Logf("e的类型是 %T", e1) // encapsulation_test.Employee
	t.Logf("e3的类型是指针类型 %T", e3) // *encapsulation_test.Employee
	t.Logf("e的类型是 %T", &e1) // 加个取指符  *encapsulation_test.Employee
}

// 行为(方法)的定义一，此方式会复制e对象，此时e是一个新的地址
func (e Employee) MyString() string {
	// 这样直接可以获取该对象，(e Employee)说明这个方法是结构体Employee的方法
	return fmt.Sprintf("ID是%s,名字是%s,年龄是%d", e.Id, e.Name, e.Age)
}

// 行为(方法)的定义二(推荐)，此方式不会复制e对象，地址不变
func (e *Employee) MyString2() string {
	// 这样直接可以获取该对象，(e Employee)说明这个方法是结构体Employee的方法
	return fmt.Sprintf("ID是%s,名字是%s,年龄是%d", e.Id, e.Name, e.Age)
}

func TestEmplyeeFunc(t *testing.T) {
	e1 := Employee{"0", "张三", 20}
	// 通过实例访问
	t.Log(e1.MyString()) // ID是0,名字是张三,年龄是20

	e2 := &Employee{"0", "李四", 30}
	// 通过结构体Employee的指针访问
	t.Log(e2.MyString2()) // ID是0,名字是李四,年龄是30
}
