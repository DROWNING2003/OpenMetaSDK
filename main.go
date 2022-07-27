package main

import (
	"github.com/DROWNING2003/OpenMetaSDK/config"
	"github.com/DROWNING2003/OpenMetaSDK/models/req"
	"github.com/DROWNING2003/OpenMetaSDK/service"
	"github.com/DROWNING2003/OpenMetaSDK/utils"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

func main() {
	var (
		cfg = pflag.StringP("config", "c", "", "./config.yaml")
	)
	fmt.Printf("cfg: %v\n", *cfg)
	//geta("eqwer")
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	// accountCreate()
	// accountQuery()
	// 创建 NFT 类别
	// nftClassCreate()
	//查询 NFT 类别
	// nftClassQuery()
	//查询 NFT 类别详情
	// nftClassDetailQuery()
	//转让 NFT 类别
	// nftClassTransfer()
	//发行 NFT
	// nftCreate()
	transferQuery()
	// nftQuery()
	//转让 NFT
	// nftTransfer()
	//编辑 NFT
	//nftEdit()
	//销毁 NFT
	//nftDelete()
	//查询 NFT
	// nftQuery()
	//查询 NFT 详情
	// nftDetailQuery()
	//查询 NFT 操作记录
	// nftHistoryQuery()
}

//上链交易结果查询
func transferQuery() {
	var taskId = "e7f53506-6794-4afa-8eba-1ea80fb753f5"
	result, err := service.TransferQuery(taskId)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	fmt.Printf("上链交易结果查询: %v\n", result)
	fmt.Printf("上链交易结果查询: %v\n", result.Data.NftId)
	fmt.Printf("上链交易结果查询: %v\n", result.Data.ClassId)
	fmt.Printf("上链交易结果查询: %v\n", result.Data.Status)
	// byteResultStr, _ := json.Marshal(result)
	// logrus.Println(string(byteResultStr))
	// byteErrStr, _ := json.Marshal(err)
	// logrus.Error(string(byteErrStr))
}

// 创建链账户
func accountCreate() {
	param := new(req.AccountCreateReq)
	param.Name = "ly"
	param.OperationId = utils.GenerateOperationId()
	result, err := service.AccountCreate(param)

	if err.Err != nil {
		logrus.Error(err.Err)
	}
	fmt.Printf("创建链账户: %v\n", result)
	// byteResultStr, _ := json.Marshal(result)
	// logrus.Println(string(byteResultStr))
	// byteErrStr, _ := json.Marshal(err)
	// logrus.Error(string(byteErrStr))
}

// 查询链账户
func accountQuery() {
	param := new(req.AccountQuery)
	param.Account = "iaa19zvgqv5lwr5mdk6a6uqngavsgfj9460vzz7n2j"
	// param.OperationId = "8662b870-cbbc-4834-992b-1407665c7100"
	result, err := service.AccountQuery(param)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	fmt.Printf("查询链账户: %v\n", result)
	// byteResultStr, _ := json.Marshal(result)
	// logrus.Println(string(byteResultStr))
	// byteErrStr, _ := json.Marshal(err)
	// logrus.Error(string(byteErrStr))
}

// nftClassCreate 创建 NFT 类别
func nftClassCreate() {
	param := new(req.NFTClassCreateReq)
	param.Name = "openmate"
	param.Owner = "iaa19zvgqv5lwr5mdk6a6uqngavsgfj9460vzz7n2j"
	param.OperationId = utils.GenerateOperationId()
	result, err := service.NFTClassesCreate(param)

	if err.Err != nil {
		logrus.Error(err.Err)
	}
	fmt.Printf("创建 NFT 类别: %v\n", result)
	// byteResultStr, _ := json.Marshal(result)
	// logrus.Println(string(byteResultStr))
	// byteErrStr, _ := json.Marshal(err)
	// logrus.Error(string(byteErrStr))
}

// nftClassQuery 查询 NFT 类别
func nftClassQuery() {
	param := new(req.NFTClassQueryReq)
	//param.Owner = "iaa12raqplxcer2lqcp0j8s9s8mq6sfyhxzkhtqtqe"
	param.Owner = "iaa19zvgqv5lwr5mdk6a6uqngavsgfj9460vzz7n2j"
	result, err := service.NFTClassesQuery(param)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	fmt.Printf("查询 NFT 类别: %v\n", result)
	// byteResultStr, _ := json.Marshal(result)
	// logrus.Println(string(byteResultStr))
	// byteErrStr, _ := json.Marshal(err)
	// logrus.Error(string(byteErrStr))
}

// nftClassDetailQuery 查询 NFT 类别详情
func nftClassDetailQuery() {
	classId := "avataqoqpkwbqr7wtunmufqmcrcd6ojc"
	result, err := service.NFTClassesDetailQuery(classId)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	fmt.Printf("查询 NFT 类别详情: %v\n", result)
	// byteResultStr, _ := json.Marshal(result)
	// logrus.Println(string(byteResultStr))
	// byteErrStr, _ := json.Marshal(err)
	// logrus.Error(string(byteErrStr))
}

// nftClassTransfer 转让 NFT 类别
func nftClassTransfer() {
	param := new(req.NFTClassTransferReq)
	param.PathParam.ClassId = "avata232b584c954d6ef5376c4dfade289891b1b2b987a1c1aa94a108b4a4113d11ee"
	param.PathParam.Owner = "iaa12raqplxcer2lqcp0j8s9s8mq6sfyhxzkhtqtqe"
	param.RequestBody.Recipient = "iaa1mg0ullv2vmceqzy83dv2ts7y782tvcxvcpf90l"
	param.RequestBody.OperationId = "safd23gvddaV9ewedcx4238"
	result, err := service.NFTClassesTransfer(param)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	fmt.Printf("转让 NFT 类别: %v\n", result)
	// byteResultStr, _ := json.Marshal(result)
	// logrus.Println(string(byteResultStr))
	// byteErrStr, _ := json.Marshal(err)
	// logrus.Error(string(byteErrStr))
}

//发行 NFT
func nftCreate() {
	param := new(req.NFTCreateReq)
	param.PathParam.ClassId = "avataqoqpkwbqr7wtunmufqmcrcd6ojc"
	param.RequestBody.Name = "openmate"
	param.RequestBody.Uri = "http://www.4399.com"
	param.RequestBody.OperationId = utils.GenerateOperationId()
	result, err := service.NFTCreate(param)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	byteResultStr, _ := json.Marshal(result)
	logrus.Println(string(byteResultStr))
	byteErrStr, _ := json.Marshal(err)
	logrus.Error(string(byteErrStr))
}

//转让 NFT
func nftTransfer() {
	param := new(req.NFTTransferReq)
	param.PathParam.ClassId = "avataqoqpkwbqr7wtunmufqmcrcd6ojc"
	param.PathParam.Owner = "iaa1c3duhxfa67ujmc9gp25u6eddg4gp5nvrmh9z4q"
	param.PathParam.NftId = "avataiph1xftj6pzv2fofda5vsbpbfln"
	param.RequestBody.Recipient = "iaa19zvgqv5lwr5mdk6a6uqngavsgfj9460vzz7n2j"
	param.RequestBody.OperationId = utils.GenerateOperationId()
	result, err := service.NFTTransfer(param)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	fmt.Printf("转让 NFT: %v\n", result)
	// byteResultStr, _ := json.Marshal(result)
	// logrus.Println(string(byteResultStr))
	// byteErrStr, _ := json.Marshal(err)
	// logrus.Error(string(byteErrStr))
}

//编辑 NFT
func nftEdit() {
	param := new(req.NFTEditReq)
	param.PathParam.ClassId = "avata232b584c954d6ef5376c4dfade289891b1b2b987a1c1aa94a108b4a4113d11ee"
	param.PathParam.Owner = "iaa16ykw9e8m3kr6uhejxwryt7q2glfq0ldutn8vsh"
	param.PathParam.NftId = "avata150019a88d8f5089da056da58ebcca410a0101c47d71cff8597d906867de9c68"
	param.RequestBody.Name = "ntf.name.edit.01"
	param.RequestBody.OperationId = "safd23gvddaV9ewedcx4245"
	result, err := service.NFTEdit(param)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	fmt.Printf("编辑 NFT: %v\n", result)
	// byteResultStr, _ := json.Marshal(result)
	// logrus.Println(string(byteResultStr))
	// byteErrStr, _ := json.Marshal(err)
	// logrus.Error(string(byteErrStr))
}

//销毁 NFT
func nftDelete() {
	param := new(req.NFTDeleteReq)
	param.PathParam.ClassId = "avata232b584c954d6ef5376c4dfade289891b1b2b987a1c1aa94a108b4a4113d11ee"
	param.PathParam.Owner = "iaa16ykw9e8m3kr6uhejxwryt7q2glfq0ldutn8vsh"
	param.PathParam.NftId = "avata150019a88d8f5089da056da58ebcca410a0101c47d71cff8597d906867de9c68"
	param.RequestBody.OperationId = "safd23gvddaV9ewedcx4246"
	result, err := service.NFTDelete(param)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	fmt.Printf("销毁 NFT: %v\n", result)
	// byteResultStr, _ := json.Marshal(result)
	// logrus.Println(string(byteResultStr))
	// byteErrStr, _ := json.Marshal(err)
	// logrus.Error(string(byteErrStr))
}

// nftQuery 查询 NFT
func nftQuery() {
	param := new(req.NFTQueryReq)
	// param.Id = "avata150019a88d8f5089da056da58ebcca410a0101c47d71cff8597d906867de9c68"
	param.Owner = "iaa19zvgqv5lwr5mdk6a6uqngavsgfj9460vzz7n2j"
	result, err := service.NFTQuery(param)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	fmt.Printf("nftQuery 查询 NFT: %v\n", result)
	// byteResultStr, _ := json.Marshal(result)
	// logrus.Println(string(byteResultStr))
	// byteErrStr, _ := json.Marshal(err)
	// logrus.Error(string(byteErrStr))
}

func nftDetailQuery() {
	classId := "avataqoqpkwbqr7wtunmufqmcrcd6ojc"
	nftId := "avataiph1xftj6pzv2fofda5vsbpbfln"
	result, err := service.NFTDetailQuery(classId, nftId)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	fmt.Printf("result: %v\n", result)
	// byteResultStr, _ := json.Marshal(result)
	// logrus.Println(string(byteResultStr))
	// byteErrStr, _ := json.Marshal(err)
	// logrus.Error(string(byteErrStr))
}

func nftHistoryQuery() {
	param := new(req.NFTHistoryQueryReq)
	param.PathParam.ClassId = "avata232b584c954d6ef5376c4dfade289891b1b2b987a1c1aa94a108b4a4113d11ee"
	param.PathParam.NftId = "avata150019a88d8f5089da056da58ebcca410a0101c47d71cff8597d906867de9c68"
	param.RequestBody.SortBy = "DATE_ASC"
	result, err := service.NFTHistoryQuery(param)
	if err.Err != nil {
		logrus.Error(err.Err)
	}
	byteResultStr, _ := json.Marshal(result)
	logrus.Println(string(byteResultStr))
	byteErrStr, _ := json.Marshal(err)
	logrus.Error(string(byteErrStr))
}
