package response

import "net/http"

type output struct {
	Code      uint        `json:"code"`
	Message   string      `json:"message"`
	Result    interface{} `json:"result,omitempty"`
	ErrorCode uint        `json:"ErrorCode,omitempty"`
	ErrorMsg  string      `json:"ErrorMsg,omitempty"`
}

// RespOK -- response output OK
func RespOK(data interface{}) interface{} {
	out := output{
		Code:    http.StatusOK,
		Message: "OK",
		Result:  data,
	}

	return out
}

// RespBad -- response bad request
func RespBad(errCode uint, errMsg string) interface{} {
	out := output{
		Code:      http.StatusBadRequest,
		Message:   "Bad Requests",
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
	}

	return out
}

// RespNotFound -- response 404
func RespNotFound(errCode uint, errMsg string) interface{} {
	out := output{
		Code:      http.StatusNotFound,
		Message:   "Not Found",
		ErrorCode: errCode,
		ErrorMsg:  errMsg,
	}
	return out
}
