package jsontest

// 使用反射结合struct的tag实现json的解析
type BasicInfo struct {
	Name string `json:"name"` // struct的tag实现Name映射到name
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}
