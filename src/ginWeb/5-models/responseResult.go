package models

type Result struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var Results *Result

func SuccessResult(data string) *Result {
	return &Result{
		Msg:  "success",
		Data: data,
	}
}

func ErrorResult(data string) *Result {
	return &Result{
		Msg:  "failure",
		Data: data,
	}
}
