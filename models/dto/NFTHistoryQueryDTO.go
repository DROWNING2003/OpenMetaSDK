package dto

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

type nftHistoryOR struct {
	TxHash    string `json:"tx_hash"`   //NFT 操作的 Tx Hash
	Operation string `json:"operation"` //NFT 操作类型 Enum: "mint" "edit" "transfer" "burn"
	Signer    string `json:"signer"`    //Tx 签名者地址
	Recipient string `json:"recipient"` //NFT 接收者地址
	Timestamp string `json:"timestamp"` //NFT 操作时间戳（UTC 时间）
}

type nftHistoryQuery struct {
	Offset           int            `json:"offset"`            //游标
	Limit            int            `json:"limit"`             //每页记录数
	TotalCount       int            `json:"total_count"`       //总记录数
	OperationRecords []nftHistoryOR `json:"operation_records"` //NFT 列表
}

type NFTHistoryQueryDTO struct {
	Status     string          `json:"status"`
	StatusCode int             `json:"statusCode"`
	Data       nftHistoryQuery `json:"data"`
}
