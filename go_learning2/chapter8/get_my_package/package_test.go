package my_package_test

import (
	"testing"
	// 导入自己的包，模块名/子目录
	"go_learning2/chapter8/my_package"
)
 
// 测试导入自己的包
func TestMyPackage(t *testing.T) {
	// 调用其他包的函数
	t.Log(my_package.GetFibonacci(5))
	
	// 以下语句出错，不可以调其他包方法名首字母小写的方法
	// t.Log(my_package.testLowerCase())
}
