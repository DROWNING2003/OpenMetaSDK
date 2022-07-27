package dto

/*
 *@author jerry.zhao
 *date: 2022/5/30
 */

type operationRecords struct {
	TxHash      string                 `json:"tx_hash"`      //操作 Tx Hash
	Module      string                 `json:"module"`       //功能模块 Enum: "account" "nft"
	Operation   float64                `json:"operation"`    //操作类型 Enum: "add_gas" "transfer_class" "mint" "edit" "transfer" "burn"
	Signer      float64                `json:"signer"`       //Tx 签名者地址
	Timestamp   string                 `json:"timestamp"`    //操作时间戳（UTC 时间）
	GasFee      int                    `json:"gas_fee"`      //链上交易消耗的能量值
	BusinessFee int                    `json:"business_fee"` //DDC 链上交易消耗的业务费
	Message     map[string]interface{} `json:"message"`      //对应不同操作类型的消息体,下方的Key只作为展示用, 实际返回中不存在该Key, 只返回对应数据
}

type accountHistoryData struct {
	Offset           int                `json:"offset"`            //游标
	Limit            int                `json:"limit"`             //每页记录数
	TotalCount       int                `json:"total_count"`       //总记录数
	OperationRecords []operationRecords `json:"operation_records"` //操作记录列表
}

type AccountHistoryQueryDTO struct {
	Status     string             `json:"status"`
	StatusCode int                `json:"statusCode"`
	Data       accountHistoryData `json:"data"`
}
