package req

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

// NFTClassCreateReq 创建NFT类别
type NFTClassCreateReq struct {
	Name        string `json:"name"`                  //链账户名称，支持 1-20 位汉字、大小写字母及数字组成的字符串
	Symbol      string `json:"symbol,omitempty"`      //标识
	Description string `json:"description,omitempty"` //描述
	Uri         string `json:"uri,omitempty"`         //链外数据链接
	UriHash     string `json:"uri_hash,omitempty"`    //链外数据 Hash
	Data        string `json:"data,omitempty"`        //自定义链上元数据
	Owner       string `json:"owner"`                 //NFT 类别权属者地址，支持任一文昌链合法链账户地址
	Tag         tag    `json:"tag,omitempty"`         //交易标签, 自定义 key：支持大小写英文字母和汉字和数字，长度6-12位，自定义 value：长度限制在64位字符，支持大小写字母和数字
	OperationId string `json:"operation_id"`          //操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

type tag struct {
	Key1 string  `json:"key1,omitempty"` //链账户地址
	Key2 string  `json:"key2,omitempty"` //链账户名称
	Key3 float64 `json:"key3,omitempty"` //Gas 余额
}
