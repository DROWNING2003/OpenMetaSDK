package dto

/*
 *@author jerry.zhao
 *date: 2022/5/30
 */

type accountsCreate struct {
	Accounts []string `json:"accounts"` //链账户地址列表
}

type AccountsCreateDTO struct {
	Status     string         `json:"status"`
	StatusCode int            `json:"statusCode"`
	Data       accountsCreate `json:"data"`
}
