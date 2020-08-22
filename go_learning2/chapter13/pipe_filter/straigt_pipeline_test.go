package pipe_filter

import "testing"

// 测试pipe-filter,几个filter串行处理
func TestStraightPipeline(t *testing.T) {
	// 实例化各个filter过滤器
	spliter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()

	// 实例化pipeline，传入各个filter
	sp := NewStraightPipeline("p1", spliter, converter, sum)
	// 调用pipeline的处理方法
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if ret != 6 {
		t.Fatalf("The expected is 6, but the actual is %d", ret)
	}
}
