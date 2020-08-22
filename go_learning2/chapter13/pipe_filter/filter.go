// Package pipefilter is to define the interfaces and the structures for pipe-filter style implementation
package pipe_filter

// Request is the input of the filter
type Request interface{}

// Response is the output of the filter
type Response interface{}

// Filter interface is the definition of the data processing components
// Pipe-Filter structure
// 过滤器接口，接收请求，返回响应，在这里实际没什么用
type Filter interface {
	Process(data Request) (Response, error)
}
