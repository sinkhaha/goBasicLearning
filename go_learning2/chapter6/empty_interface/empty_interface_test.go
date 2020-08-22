package empty_interface_test

import (
	"testing"
	"fmt"
)

// 测试空接口
func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}

// 空接口类型
func DoSomething(p interface{}) {
	// 可以用if
	// if i,ok := p.(int); ok {
	// 	fmt.Println("interger", i)
	// 	return
	// }
	// if s,ok := p.(string); ok {
	// 	fmt.Println("string", s)
	// 	return
	// }
	// fmt.Println("unknow type")
	// 也可以用switch
	switch v := p.(type) {
	case int:
		fmt.Println("interger", v)
	case string:
		fmt.Println("string", v)	
	default:
		fmt.Println("unknow type")		
	}
}

