package dto

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

type nftCreate struct {
	TaskId      string `json:"task_id"`      //Task ID
	OperationId string `json:"operation_id"` //操作 ID
}

type NFTCreateDTO struct {
	Status     string    `json:"status"`
	StatusCode int       `json:"statusCode"`
	Data       nftCreate `json:"data"`
}
