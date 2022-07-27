package dto

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

type nftData struct {
	Id          string `json:"id"`           //NFT ID
	Name        string `json:"name"`         //NFT 名称
	ClassId     string `json:"class_id"`     //NFT 类别 ID
	ClassName   string `json:"class_name"`   //NFT 类别 名称
	ClassSymbol string `json:"class_symbol"` //NFT 类别标识
	Uri         string `json:"uri"`          //链外数据链接
	Owner       string `json:"owner"`        //NFT 持有者地址
	Status      string `json:"status"`       //NFT 状态：Active; Burned;
	TxHash      string `json:"tx_hash"`      //NFT 发行 Tx Hash
	Timestamp   string `json:"timestamp"`    //NFT 发行时间戳（UTC 时间）
}

type nftQuery struct {
	Offset     int       `json:"offset"`      //游标
	Limit      int       `json:"limit"`       //每页记录数
	TotalCount int       `json:"total_count"` //总记录数
	Nfts       []nftData `json:"nfts"`        //NFT 列表
}

type NFTQueryDTO struct {
	Status     string   `json:"status"`
	StatusCode int      `json:"statusCode"`
	Data       nftQuery `json:"data"`
}
