package dto

/*
 *@author jerry.zhao
 *date: 2022/5/31
 */

type nftDetailQuery struct {
	Id          string `json:"id"`           //NFT ID
	Name        string `json:"name"`         //NFT 名称
	ClassId     string `json:"class_id"`     //NFT 类别 ID
	ClassName   string `json:"class_name"`   //NFT 类别 名称
	ClassSymbol string `json:"class_symbol"` //NFT 类别标识
	Uri         string `json:"uri"`          //链外数据链接
	UriHash     string `json:"uri_hash"`     //链外数据 Hash
	Data        string `json:"data"`         //自定义链上元数据
	Owner       string `json:"owner"`        //NFT 持有者地址
	Status      string `json:"status"`       //NFT 状态：Active; Burned;
	TxHash      string `json:"tx_hash"`      //NFT 发行 Tx Hash
	Timestamp   string `json:"timestamp"`    //NFT 发行时间戳（UTC 时间）
}

type NFTDetailQueryDTO struct {
	Status     string         `json:"status"`
	StatusCode int            `json:"statusCode"`
	Data       nftDetailQuery `json:"data"`
}
