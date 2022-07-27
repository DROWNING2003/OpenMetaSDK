package req

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

// AccountHistoryQueryReq 查询链账户操作记录
type AccountHistoryQueryReq struct {
	Offset    string `json:"offset,omitempty"`  //游标，默认为 0
	Limit     string `json:"limit,omitempty"`   //每页记录数，默认为 10，上限为 50
	Account   string `json:"account,omitempty"` //链账户地址
	Module    string `json:"module,omitempty"`  //功能模块：account / nft
	Operation string `json:"operation,omitempty"`
	// 操作类型，仅 module 不为空时有效，默认为 ”all“
	// module = account 时可选：add_gas
	// module = nft 时可选：transfer_class / mint / edit / transfer / burn / issue_class
	StartDate string `json:"start_date,omitempty"` //创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate   string `json:"end_date,omitempty"`   //创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy    string `json:"sort_by,omitempty"`    //排序规则：DATE_ASC / DATE_DESC
	TxHash    string `json:"tx_hash,omitempty"`    //Tx Hash
}
