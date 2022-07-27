package req

/*
 *@author jerry.zhao
 *date: 2022/5/29
 */

// AccountQuery 查询链账户 QUERY PARAMETERS
type AccountQuery struct {
	Offset      string `json:"offset,omitempty"`       //游标，默认为 0
	Limit       string `json:"limit,omitempty"`        //每页记录数，默认为 10，上限为 50
	Account     string `json:"account,omitempty"`      //链账户地址
	StartDate   string `json:"start_date,omitempty"`   //创建日期范围 - 开始，yyyy-MM-dd（UTC 时间）
	EndDate     string `json:"end_date,omitempty"`     //创建日期范围 - 结束，yyyy-MM-dd（UTC 时间）
	SortBy      string `json:"sort_by,omitempty"`      //排序规则：DATE_ASC / DATE_DESC
	OperationId string `json:"operation_id,omitempty"` //操作 ID
}
