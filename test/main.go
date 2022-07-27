package main

import (
	"GoForAvataAPIDemo/config"
	"GoForAvataAPIDemo/models/req"
	"GoForAvataAPIDemo/service"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	var (
		cfg = pflag.StringP("config", "c", "", "./config.yaml")
	)
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	// 创建链账户
	// accountCreate()
	// 批量创建链账户
	//accountsCreate()
	// 查询链账户
	// accountQuery()
	// 查询链账户操作记录
	//accountHistoryQuery()

}

// 创建链账户
func accountCreate() {
	param := new(req.AccountCreateReq)
	param.Name = "dr"
	param.OperationId = "safd23gvddaV9ewedcx4233"
	result, err := service.AccountCreate(param)
	fmt.Printf("result: %v\n", result)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	byteResultStr, _ := json.Marshal(result)
	logrus.Println(string(byteResultStr))
	byteErrStr, _ := json.Marshal(err)
	logrus.Error(string(byteErrStr))
}

// 批量创建链账户
func accountsCreate() {
	param := new(req.AccountsCreateReq)
	param.OperationId = "safd23gvddaV9ewedcx4223"
	param.Count = 2
	result, err := service.AccountsCreate(param)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	byteResultStr, _ := json.Marshal(result)
	logrus.Println(string(byteResultStr))
	byteErrStr, _ := json.Marshal(err)
	logrus.Error(string(byteErrStr))
}

// 查询链账户
func accountQuery() {
	param := new(req.AccountQuery)
	param.OperationId = ""
	result, err := service.AccountQuery(param)
	if err.Err != nil {
		fmt.Println("69err")
		logrus.Error(err.Err)
	}
	fmt.Printf("result: %v\n", result)
	byteResultStr, _ := json.Marshal(result)
	logrus.Println(string(byteResultStr))
	byteErrStr, _ := json.Marshal(err)
	logrus.Error(string(byteErrStr))
}

// 查询链账户操作记录
func accountHistoryQuery() {
	param := new(req.AccountHistoryQueryReq)
	param.Account = "iaa1d64yzft6saw9a83gjywklalkzh44gyxfdpc6l2"
	result, err := service.AccountHistoryQuery(param)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	byteResultStr, _ := json.Marshal(result)
	logrus.Println(string(byteResultStr))
	byteErrStr, _ := json.Marshal(err)
	logrus.Error(string(byteErrStr))
}
