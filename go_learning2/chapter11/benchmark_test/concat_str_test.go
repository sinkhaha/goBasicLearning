package benchmark_test

import (
	"bytes"
	"testing"
)

// 测试benchmark,运行go test -bench=方法名 运行指定的benchmark方法
func BenchmarkConcatStringByAdd(b *testing.B) {

	elems := []string{"1", "2", "3", "4", "5"}
	// 与性能测试无关的代码
	b.ResetTimer()
	// 测试代码
	for i := 0; i < b.N; i++ {
		ret := ""
		for _, elem := range elems {
			ret += elem
		}
	}
	b.StopTimer()
	// 与性能测试无关的代码
}

func BenchmarkConcatStringByBytesBuffer(b *testing.B) {
	elems := []string{"1", "2", "3", "4", "5"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer

		for _, elem := range elems {
			buf.WriteString(elem)

		}
	}
	b.StopTimer()
}