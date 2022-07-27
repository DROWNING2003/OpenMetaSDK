package req

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

// AccountsCreateReq 批量创建链账户
type AccountsCreateReq struct {
	Count       int    `json:"count,omitempty"` //批量创建链账户的数量 Default: 1
	OperationId string `json:"operation_id"`
	//操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、
	//针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}
