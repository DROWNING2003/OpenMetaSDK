package dto

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

type nftClassTransfer struct {
	TaskId      string `json:"task_id"`      //Task ID
	OperationId string `json:"operation_id"` //操作 ID
}

type NFTClassTransferDTO struct {
	Status     string           `json:"status"`
	StatusCode int              `json:"statusCode"`
	Data       nftClassTransfer `json:"data"`
}
