package dto

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

type nftClassCreate struct {
	TaskId      string `json:"task_id"`      //Task ID
	OperationId string `json:"operation_id"` //操作 ID
}

type NFTClassCreateDTO struct {
	Status     string         `json:"status"`
	StatusCode int            `json:"statusCode"`
	Data       nftClassCreate `json:"data"`
}
