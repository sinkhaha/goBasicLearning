package err_test

import (
	"errors"
	"fmt"
	"testing"
)

// 自己定义错误类型,errors.New创建错误实例子
var LessTwoErr = errors.New("n小于2")
var LargerHundredErr = errors.New("大于100")

func GetFibonacci(n int) ([]int, error) {
	// 添加错误机制
	if n < 2 {
		// return nil,errors.New("n小于2")
		// 改成返回同一个错误实例
		return nil, LessTwoErr
	}
	if n > 100 {
		// return nil,errors.New("大于100")
		return nil, LargerHundredErr
	}
	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

// 测试错误处理
func TestGetFibonacci(t *testing.T) {
	// 处理返回的错误
	if result, err := GetFibonacci(0); err != nil {
		// 判断错误类型
		if err == LessTwoErr {
			fmt.Println("这是小于2的错误", err)
		}
		if err == LargerHundredErr {
			fmt.Println("这是大于100的错误", err)
		}
		t.Error(err)
	} else {
		t.Log(result)
	}
}
