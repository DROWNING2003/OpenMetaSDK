package dto

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

// Accounts 账户
type accounts struct {
	Account     string  `json:"account"`      //链账户地址
	Name        string  `json:"name"`         //链账户名称
	Gas         float64 `json:"gas"`          //Gas 余额
	Business    float64 `json:"business"`     //业务费余额
	OperationId string  `json:"operation_id"` //操作 ID
	Status      int     `json:"status"`       //链账户的授权状态，0 未授权；1 已授权
}

// AccountQueryDTO 查询链账户DTO
type accountQuery struct {
	Offset     int        `json:"offset"`      //游标
	Limit      int        `json:"limit"`       //每页记录数
	TotalCount int        `json:"total_count"` //总记录数
	Accounts   []accounts `json:"accounts"`    //链账户列表
}

type AccountQueryDTO struct {
	Status     string       `json:"status"`
	StatusCode int          `json:"statusCode"`
	Data       accountQuery `json:"data"`
}
