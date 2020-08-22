package pipe_filter

import "errors"

var SumFilterWrongFormatError = errors.New("input data should be []int")

type SumFilter struct {
}

// 实例化
func NewSumFilter() *SumFilter {
	return &SumFilter{}
}

// 求和处理实现方法
func (sf *SumFilter) Process(data Request) (Response, error) {
	elems, ok := data.([]int)
	if !ok {
		return nil, SumFilterWrongFormatError
	}
	ret := 0
	for _, elem := range elems {
		ret += elem
	}
	return ret, nil
}
