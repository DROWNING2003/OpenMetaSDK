package dto

/*
 *@author jerry.zhao
 *date: 2022/5/30
 */

type nftClasses struct {
	Id        string `json:"id"`        //NFT 类别 ID
	Name      string `json:"name"`      //NFT 类别名称
	Symbol    string `json:"symbol"`    //NFT 类别标识
	NftCount  int    `json:"nft_count"` //NFT 类别包含的 NFT 总量
	Uri       string `json:"uri"`       //链外数据链接
	Owner     string `json:"owner"`     //NFT 类别权属者地址
	TxHash    string `json:"tx_hash"`   //创建 NFT 类别的 Tx Hash
	Timestamp string `json:"timestamp"` //创建 NFT 类别的时间戳（UTC 时间）
}

type nftClassQueryData struct {
	Offset     int          `json:"offset"`      //游标
	Limit      int          `json:"limit"`       //每页记录数
	TotalCount int          `json:"total_count"` //总记录数
	Classes    []nftClasses `json:"classes"`     //操作记录列表
}

type NFTClassQueryDTO struct {
	Status     string            `json:"status"`
	StatusCode int               `json:"statusCode"`
	Data       nftClassQueryData `json:"data"`
}
