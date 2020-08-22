package client

import "testing"
// 导入自己的包，src下的目录开始写
import "chapter8/myPackage"

// 测试导入自己的包
func TestMyPackage(t *testing.T) {
	// 调用其他包的函数
	t.Log(myPackage.GetFibonacci(5))
}
