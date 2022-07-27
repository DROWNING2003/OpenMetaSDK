package dto

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

// ErrorData 错误信息
type errorData struct {
	Code      string `json:"code"`
	CodeSpace string `json:"code_space"`
	Message   string `json:"message"`
}

// ErrorDTO 错误返回值
type ErrorDTO struct {
	Err        error     `json:"err"`
	Status     string    `json:"status"`
	StatusCode int       `json:"statusCode"`
	Error      errorData `json:"error"`
}
