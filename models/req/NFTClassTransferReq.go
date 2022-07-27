package req

/*
 *@author jerry.zhao
 *date: 2022/5/31
 */

type nftClassTransferPathParam struct {
	ClassId string `json:"class_id"` //NFT 类别 ID
	Owner   string `json:"owner"`    //NFT 类别权属者地址
}

type nftClassTransferRequestBody struct {
	Recipient   string `json:"recipient"`     //游标，默认为 0
	OperationId string `json:"operation_id"`  //操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
	Tag         tag    `json:"tag,omitempty"` //交易标签, 自定义 key：支持大小写英文字母和汉字和数字，长度6-12位，自定义 value：长度限制在64位字符，支持大小写字母和数字
}

// NFTClassTransferReq 转让 NFT 类别
type NFTClassTransferReq struct {
	PathParam   nftClassTransferPathParam   `json:"pathParam"`   //PATH PARAMETERS
	RequestBody nftClassTransferRequestBody `json:"requestBody"` //REQUEST BODY
}
