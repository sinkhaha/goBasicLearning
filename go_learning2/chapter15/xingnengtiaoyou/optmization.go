package profiling

import (
	"encoding/json"
	"strconv"
	"strings"
)

// 构造一个请求
func createRequest() string {
	payload := make([]int, 100, 100)
	for i := 0; i < 100; i++ {
		payload[i] = i
	}
	req := Request{"demo_transaction", payload}
	v, err := json.Marshal(&req)
	if err != nil {
		panic(err)
	}
	return string(v)
}

// 新的处理请求方式（easyjson处理）
func processRequest(reqs []string) []string {
	reps := []string{}
	for _, req := range reqs {
		reqObj := &Request{}
		reqObj.UnmarshalJSON([]byte(req))
		//	json.Unmarshal([]byte(req), reqObj)

		var buf strings.Builder
		for _, e := range reqObj.PayLoad {
			buf.WriteString(strconv.Itoa(e))
			buf.WriteString(",")
		}
		repObj := &Response{reqObj.TransactionID, buf.String()}
		repJson, err := repObj.MarshalJSON()
		//repJson, err := json.Marshal(&repObj)
		if err != nil {
			panic(err)
		}
		reps = append(reps, string(repJson))
	}
	return reps
}

// 旧的处理请求方式（内置json处理）
func processRequestOld(reqs []string) []string {
	reps := []string{}
	for _, req := range reqs {
		reqObj := &Request{}
		json.Unmarshal([]byte(req), reqObj)
		ret := ""
		for _, e := range reqObj.PayLoad {
			ret += strconv.Itoa(e) + ","
		}
		repObj := &Response{reqObj.TransactionID, ret}
		repJson, err := json.Marshal(&repObj)
		if err != nil {
			panic(err)
		}
		reps = append(reps, string(repJson))
	}
	return reps
}
