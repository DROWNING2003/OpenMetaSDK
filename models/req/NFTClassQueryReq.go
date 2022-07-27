package req

/*
 *@author jerry.zhao
 *date: 2022/5/31
 */

// NFTClassQueryReq 查询 NFT 类别
type NFTClassQueryReq struct {
	Offset    string `json:"offset,omitempty"`     //游标，默认为 0
	Limit     string `json:"limit,omitempty"`      //每页记录数，默认为 10，上限为 50
	Id        string `json:"id,omitempty"`         //NFT 类别 ID
	Name      string `json:"name,omitempty"`       //NFT 类别名称，支持模糊查询
	Owner     string `json:"owner,omitempty"`      //NFT 类别权属者地址
	TxHash    string `json:"tx_hash,omitempty"`    //创建 NFT 类别的 Tx Hash
	StartDate string `json:"start_date,omitempty"` //NFT 类别创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate   string `json:"end_date,omitempty"`   //NFT 类别创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy    string `json:"sort_by,omitempty"`    //排序规则：DATE_ASC / DATE_DESC
}
