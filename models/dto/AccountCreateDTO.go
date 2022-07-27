package dto

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

type accountCreate struct {
	Account     string `json:"account"`      //链账户地址
	Name        string `json:"name"`         //链账户名称
	OperationId string `json:"operation_id"` //操作 ID
}

type AccountCreateDTO struct {
	Status     string        `json:"status"`
	StatusCode int           `json:"statusCode"`
	Data       accountCreate `json:"data"`
}
