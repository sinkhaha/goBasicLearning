package easyjson

import (
	"encoding/json"
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
	  	"name":"Mike",
		"age":30
	},
	"job_info":{
		"skills":["Java","Go","C"]
	}
}	`

// 测试go内置的json解析
func TestEmbeddedJson(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}

}

// 测试使用easyjson进行解析json
func TestEasyJson(t *testing.T) {
	e := Employee{}

	// json => 对象
	e.UnmarshalJSON([]byte(jsonStr))
	fmt.Println("对象是：", e)

	// 对象=>json
	if v, err := e.MarshalJSON(); err != nil {
		t.Error(err)
	} else {
		fmt.Println("json字符串是：", string(v))
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

// 运行 go test -bench=.
// 性能测试使用easyjson进行json解析
func BenchmarkEasyJson(b *testing.B) {
	b.ResetTimer()
	e := Employee{}
	for i := 0; i < b.N; i++ {
		err := e.UnmarshalJSON([]byte(jsonStr))
		if err != nil {
			b.Error(err)
		}
		if _, err = e.MarshalJSON(); err != nil {
			b.Error(err)
		}
	}
}
