package dto

/*
 *@author jerry.zhao
 *date: 2022/5/30
 */

type nftClassDetailQueryData struct {
	Id          string `json:"id"`          //NFT 类别 ID
	Name        string `json:"name"`        //NFT 类别名称
	Symbol      string `json:"symbol"`      //NFT 类别标识
	Description string `json:"description"` //NFT 类别描述
	NftCount    int    `json:"nft_count"`   //NFT 类别包含的 NFT 总量
	Uri         string `json:"uri"`         //链外数据链接
	UriHash     string `json:"uri_hash"`    //链外数据 Hash
	Data        string `json:"data"`        //自定义链上元数据
	Owner       string `json:"owner"`       //NFT 类别权属者地址
	TxHash      string `json:"tx_hash"`     //创建 NFT 类别的 Tx Hash
	Timestamp   string `json:"timestamp"`   //创建 NFT 类别的时间戳（UTC 时间）
}

type NFTClassDetailQueryDTO struct {
	Status     string                  `json:"status"`
	StatusCode int                     `json:"statusCode"`
	Data       nftClassDetailQueryData `json:"data"`
}
