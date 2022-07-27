package main

import (
	"github.com/DROWNING2003/OpenMetaSDK/config"
	"github.com/DROWNING2003/OpenMetaSDK/service"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	var (
		cfg = pflag.StringP("config", "c", "", "./config.yaml")
	)
	//geta("eqwer")
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// 上链交易结果查询
	transferQuery()
}

//上链交易结果查询
func transferQuery() {
	var taskId = "safd23gvddaV9ewedcx4247"
	result, err := service.TransferQuery(taskId)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	byteResultStr, _ := json.Marshal(result)
	logrus.Println(string(byteResultStr))
	byteErrStr, _ := json.Marshal(err)
	logrus.Error(string(byteErrStr))
}
