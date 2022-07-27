package req

/*
 *@author jerry.zhao
 *date: 2022/5/31
 */

type nftHistoryPathParam struct {
	ClassId string `json:"class_id"` //NFT 类别 ID
	NftId   string `json:"nft_id"`   //NFT ID
}

type nftHistoryRequestBody struct {
	Offset    string `json:"offset,omitempty"`     //游标，默认为 0
	Limit     string `json:"limit,omitempty"`      //每页记录数，默认为 10，上限为 50
	Signer    string `json:"signer,omitempty"`     //Tx 签名者地址
	TxHash    string `json:"tx_hash,omitempty"`    //NFT 操作 Tx Hash
	Operation string `json:"operation"`            //操作类型：mint / edit / transfer / burn
	StartDate string `json:"start_date,omitempty"` //NFT 类别创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate   string `json:"end_date,omitempty"`   //NFT 类别创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy    string `json:"sort_by,omitempty"`    //排序规则：DATE_ASC / DATE_DESC
}

// NFTHistoryQueryReq 编辑NFT
type NFTHistoryQueryReq struct {
	PathParam   nftHistoryPathParam   `json:"pathParam"`   //PATH PARAMETERS
	RequestBody nftHistoryRequestBody `json:"requestBody"` //REQUEST BOD
}
