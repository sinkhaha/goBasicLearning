package profiling

import "testing"

func TestCreateRequest(t *testing.T) {
	str := createRequest()
	t.Log("构造后的请求信息：", str)
}

func TestProcessRequest(t *testing.T) {
	reqs := []string{}
	reqs = append(reqs, createRequest())
	reps := processRequest(reqs)
	t.Log("处理请求后的响应信息：", reps[0])
}

// 性能测试新方式的请求处理
func BenchmarkProcessRequest(b *testing.B) {

	reqs := []string{}
	reqs = append(reqs, createRequest())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processRequest(reqs)
	}
	b.StopTimer()

}

// 性能测试旧方式的请求处理
func BenchmarkProcessRequestOld(b *testing.B) {

	reqs := []string{}
	reqs = append(reqs, createRequest())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = processRequestOld(reqs)
	}
	b.StopTimer()

}
