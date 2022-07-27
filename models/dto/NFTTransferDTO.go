package dto

/*
 *@author jerry.zhao
 *date: 2022/5/31
 */

type nftTransfer struct {
	TaskId      string `json:"task_id"`      //Task ID
	OperationId string `json:"operation_id"` //操作 ID
}

type NFTTransferDTO struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Data       nftTransfer `json:"data"`
}
