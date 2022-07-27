package service

import (
	"github.com/DROWNING2003/OpenMetaSDK/models/dto"
	"github.com/DROWNING2003/OpenMetaSDK/models/req"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
)

// AccountCreate 创建链账户
func AccountCreate(param *req.AccountCreateReq) (res dto.AccountCreateDTO, errDTO dto.ErrorDTO) {
	p, err := json.Marshal(&param)
	if err != nil {
		errDTO.Err = err
		return
	}
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.create_account")
	resp, err := PostHeader(url, p)
	if err != nil {
		errDTO.Err = err
		logrus.Error(err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err.Error())
		}
	}(resp.Body)
	if err != nil {
		errDTO.Err = err
		logrus.Error(err.Error())
		return
	}
	if resp.StatusCode == 400 {
		err = json.Unmarshal(body, &errDTO)
	} else if resp.StatusCode == 200 {
		err = json.Unmarshal(body, &res)
	}
	// 设置状态和状态码
	errDTO.Status = resp.Status
	errDTO.StatusCode = resp.StatusCode
	res.Status = resp.Status
	res.StatusCode = resp.StatusCode
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	return res, errDTO
}

// AccountsCreate 批量创建链账户
func AccountsCreate(param *req.AccountsCreateReq) (res dto.AccountsCreateDTO, errDTO dto.ErrorDTO) {
	p, err := json.Marshal(&param)
	if err != nil {
		errDTO.Err = err
		return
	}
	var url = viper.GetString("nft.domain") + viper.GetString("nft.api.url.create_account_batch")
	resp, err := PostHeader(url, p)

	if err != nil {
		errDTO.Err = err
		logrus.Error(err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err.Error())
		}
	}(resp.Body)
	if err != nil {
		errDTO.Err = err
		logrus.Error(err.Error())
		return
	}
	if resp.StatusCode == 400 {
		err = json.Unmarshal(body, &errDTO)
	} else if resp.StatusCode == 200 {
		err = json.Unmarshal(body, &res)
	}
	// 设置状态和状态码
	errDTO.Status = resp.Status
	errDTO.StatusCode = resp.StatusCode
	res.Status = resp.Status
	res.StatusCode = resp.StatusCode
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	return res, errDTO
}

// AccountQuery 查询链账户
func AccountQuery(param *req.AccountQuery) (res dto.AccountQueryDTO, errDTO dto.ErrorDTO) {
	p, _ := json.Marshal(&param)
	data := make(map[string]string)
	err := json.Unmarshal(p, &data)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.query_account")
	resp, err := GetHeader(url, data)
	if err != nil {
		errDTO.Err = err
		logrus.Error(err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err.Error())
		}
	}(resp.Body)
	if err != nil {
		errDTO.Err = err
		logrus.Error(err.Error())
		return
	}
	if resp.StatusCode == 400 {
		err = json.Unmarshal(body, &errDTO)
	} else if resp.StatusCode == 200 {
		err = json.Unmarshal(body, &res)
	}
	// 设置状态和状态码
	errDTO.Status = resp.Status
	errDTO.StatusCode = resp.StatusCode
	res.Status = resp.Status
	res.StatusCode = resp.StatusCode
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	return res, errDTO
}

// AccountHistoryQuery 查询链账户操作记录
func AccountHistoryQuery(param *req.AccountHistoryQueryReq) (res dto.AccountHistoryQueryDTO, errDTO dto.ErrorDTO) {
	p, _ := json.Marshal(&param)
	data := make(map[string]string)
	err := json.Unmarshal(p, &data)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.query_account_history")
	resp, err := GetHeader(url, data)
	if err != nil {
		errDTO.Err = err
		logrus.Error(err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			logrus.Error(err.Error())
		}
	}(resp.Body)
	if err != nil {
		errDTO.Err = err
		logrus.Error(err.Error())
		return
	}
	if resp.StatusCode == 400 {
		err = json.Unmarshal(body, &errDTO)
	} else if resp.StatusCode == 200 {
		err = json.Unmarshal(body, &res)
	}
	// 设置状态和状态码
	errDTO.Status = resp.Status
	errDTO.StatusCode = resp.StatusCode
	res.Status = resp.Status
	res.StatusCode = resp.StatusCode
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	return res, errDTO
}
