package encapsulation_test

import (
	"fmt"
	"testing"
)

// 结构体定义
type Employee struct {
	Id   string // 不需要逗号
	Name string
	Age  int
}

// 实例化结构体的3中方式
func TestCreateEmployeeObj(t *testing.T) {
	// 方式一
	e1 := Employee{"0", "张三", 20}
	t.Log(e1) // {0 张三 20}
	e11 := &e1
	e11.Id = "11" // 改变了e1地址上Id的值，所以会影响到e1的Id的值
	t.Log(e1)     // {11 张三 20}
	t.Log(e11)    // &{11 张三 20}

	// 方式二
	e2 := Employee{Name: "李四", Age: 30}
	t.Log(e2) // { 李四 30}

	// 方式三 获取地址
	e3 := new(Employee) // 这里返回引用，等价于e3 := &Employee{}
	t.Log(e3)           // &{  0}

	e3.Id = "4"
	e3.Name = "老五"
	e3.Age = 40
	t.Log(e3) // &{4 老五 40}

	t.Logf("e的类型是 %T", e1)      // encapsulation_test.Employee
	t.Logf("e3的类型是指针类型 %T", e3) // *encapsulation_test.Employee
	t.Logf("e的类型是 %T", &e1)     // 加个取指符  *encapsulation_test.Employee
}

// 实例方法的定义(2种)
// 方法名前面加上结构体名，则说明这个方法是结构体的方法，结构体实例可以直接调用该方法
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

func TestEmployeeFunc(t *testing.T) {
	e1 := Employee{"0", "张三", 20}
	// 通过实例访问
	t.Log(e1.MyString()) // ID是0,名字是张三,年龄是20

	e2 := &Employee{"0", "李四", 30}
	// 通过结构体Employee的指针访问
	t.Log(e2.MyString2()) // ID是0,名字是李四,年龄是30
}
