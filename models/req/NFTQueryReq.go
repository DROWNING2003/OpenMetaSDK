package req

/*
 *@author jerry.zhao
 *date: 2022/5/31
 */

// NFTQueryReq 查询 NFT
type NFTQueryReq struct {
	Offset    string `json:"offset,omitempty"`     //游标，默认为 0
	Limit     string `json:"limit,omitempty"`      //每页记录数，默认为 10，上限为 50
	Id        string `json:"id,omitempty"`         //NFT ID
	ClassId   string `json:"class_id,omitempty"`   //NFT 类别 ID
	Owner     string `json:"owner,omitempty"`      //NFT 持有者地址
	TxHash    string `json:"tx_hash,omitempty"`    //创建 NFT 的 Tx Hash
	StartDate string `json:"start_date,omitempty"` //NFT 创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate   string `json:"end_date,omitempty"`   //NFT 创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy    string `json:"sort_by,omitempty"`    //排序规则：ID_ASC / ID_DESC / DATE_ASC / DATE_DESC
	Status    string `json:"status,omitempty"`     //NFT 状态：active / burned，默认为 active
}
