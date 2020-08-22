package my_package // 自己定义的包，自己定义的包里不能有main函数,使用了mod，所以包名要和当前目录名一样

import "fmt"

// 创建package里面的init方法
func init() {
    fmt.Println("init1")
}

func init() {
    fmt.Println("init2")
}

// 首字母必须大写，小写外部访问不了
func GetFibonacci(n int) []int{
	fibList := []int{1, 1}
	for i:=2; i<n; i++ {
       fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList
}

// 测试方法名首字母小写不可以被外部调用,外部引用会报错
func testLowerCase(){
	fmt.Println("首字母小写不可被外部调用")
}
