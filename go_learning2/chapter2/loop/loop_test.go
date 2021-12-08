package loop_test

import "testing"

// 测试for代替while循环
func TestLoopWhile(t *testing.T) {
	/** 相当于while(n<5)*/
	n := 0
	for n < 5 {
		t.Log(n)
		n++
	}
}

// 测试for代替while的无限循环
func TestLoopWhileTrue(t *testing.T) {
	/** while(true)*/
	n := 0
	for {
		t.Log(n)
		if n == 4 {
			break
		}
		n++
	}
}

// if语句的赋值支持声明变量
func TestIf(t *testing.T) {
	if a := 1; a == 1 {
		t.Log("1==1")
	}
}

// 测试switch的case可以匹配多个值，不需要手写break
func TestSwitchMuitlCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2: // 匹配多个值
			t.Log("偶数", i)
		case 1, 3:
			t.Log("奇数", i)
		default:
			t.Log("不是0-3")
		}
	}
}

// 测试switch实现if else的效果
func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch { // 直接中括号，类似if-else if-else
		case i%2 == 0:
			t.Log("偶数")
		case i%2 == 1:
			t.Log("奇数")
		default:
			t.Log("不是0-3")
		}
	}
}

// 测试switch可以定义变量
func TestSwitchCaseDecalar(t *testing.T) {
	switch j := 2; j { // 可以定义j
	case 1:
		t.Log("1")
	case 2:
		t.Log("2")
	default:
		t.Log("默认")
	}
}
