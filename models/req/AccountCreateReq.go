package req

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

// AccountCreateReq 创建链账户
type AccountCreateReq struct {
	/*
	 *链账户名称，支持 1-20 位汉字、大小写字母及数字组成的字符串
	 */
	Name string `json:"name"`
	/*
	 *操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
	 */
	OperationId string `json:"operation_id"`
}
