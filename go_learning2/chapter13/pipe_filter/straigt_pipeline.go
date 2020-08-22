package pipe_filter

// StraightPipeline is composed of the filters, and the filters are piled as a straigt line.
// 定义一个pipeline, 里面有很多重filter
type StraightPipeline struct {
	Name    string
	Filters *[]Filter
}

// NewStraightPipeline create a new StraightPipelineWithWallTime
// 实例化
func NewStraightPipeline(name string, filters ...Filter) *StraightPipeline {
	return &StraightPipeline{
		Name:    name,
		Filters: &filters,
	}
}

// Process is to process the coming data by the pipeline
func (f *StraightPipeline) Process(data Request) (Response, error) {
	var ret interface{}
	var err error
	// 处理pipeline里面的各个filter
	for _, filter := range *f.Filters {
		ret, err = filter.Process(data)
		if err != nil {
			return ret, err
		}
		data = ret
	}
	return ret, err
}
