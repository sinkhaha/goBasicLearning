package jsontest

import (
	"encoding/json"
	"fmt"
	"testing"
)

// json字符串
var jsonStr = `{
	"basic_info":{
	  	"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}	`

// 测试json解析
func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)

	// 字符串到空对象的解析
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("填充后e的信息是:", *e) // 填充后e的信息是: {{Mike 30} {[Java Go C]}}

	// 把对象转换成json字符串
	if v, err := json.Marshal(e); err == nil {
		fmt.Println("e转成json串是:", string(v)) // e转成json串是: {"basic_info":{"name":"Mike","age":30},"job_info":{"skills":["Java","Go","C"]}}
	} else {
		t.Error(err)
	}
}

// 运行 go test -bench=.
// 性能测试使用内置go进行json解析
func BenchmarkEmbeddedJson(b *testing.B) {
	b.ResetTimer()
	e := new(Employee)
	for i := 0; i < b.N; i++ {

		err := json.Unmarshal([]byte(jsonStr), e)
		if err != nil {
			b.Error(err)
		}
		if _, err = json.Marshal(e); err != nil {
			b.Error(err)
		}
	}
}
