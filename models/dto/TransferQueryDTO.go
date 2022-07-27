package dto

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

type tag struct {
	Key1 string  `json:"key1"` //链账户地址
	Key2 string  `json:"key2"` //链账户名称
	Key3 float64 `json:"key3"` //Gas 余额
}

type transferQuery struct {
	Type        string `json:"type"`         //用户操作类型 Enum: "issue_class" "mint_nft" "edit_nft" "burn_nft" "transfer_class" "transfer_nft"
	TxHash      string `json:"tx_hash"`      //交易哈希
	Status      int    `json:"status"`       //交易状态, 0 处理中; 1 成功; 2 失败; 3 未处理
	ClassId     string `json:"class_id"`     //类别 ID
	NftId       string `json:"nft_id"`       //NFT ID
	Message     string `json:"message"`      //错误描述
	BlockHeight int    `json:"block_height"` //交易上链的区块高度
	Timestamp   string `json:"timestamp"`    //交易上链时间（UTC 时间）
	Tag         tag    `json:"tag"`          //交易标签, 自定义 key：支持大小写英文字母和汉字和数字，长度6-12位，自定义 value：长度限制在64位字符，支持大小写字母和数字

}

// TransferQueryDTO 上链交易查询DTO
type TransferQueryDTO struct {
	Status     string        `json:"status"`
	StatusCode int           `json:"statusCode"`
	Data       transferQuery `json:"data"`
}
