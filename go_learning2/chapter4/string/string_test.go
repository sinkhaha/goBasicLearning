package string_test

import (
	"strconv"
	"strings"
	"testing"
)

// 测试string的初始化
func TestStringInit(t *testing.T) {
	var s string
	t.Log(s) // 初始化为默认值""

	s = "hello"
	t.Log(s)      // hello
	t.Log(len(s)) // 5 指5个byte数，并不是字符

	s = "中"
	t.Log(len(s)) // 3 指3个byte数，并不是字符数

	s = "\xE4\xB8\xA5" // 十六进制串
	t.Log(s)           // 严
	t.Log(s[1])        // 184
}

// 测试strings包的函数
func TestStringFn(t *testing.T) {
	str := "a,b,c"

	parts := strings.Split(str, ",")
	for _, v := range parts {
		t.Log(v)
	}

	t.Log(strings.Join(parts, "-")) // a-b-c

	// 数字10转字符串
	s := strconv.Itoa(10)
	t.Log("str" + s)

	// 字符串10转数字10
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i) // 20
	}
}
