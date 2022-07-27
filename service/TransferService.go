package service

import (
	"GoForAvataAPIDemo/models/dto"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"strings"
)

// TransferQuery 上链交易结果查询
func TransferQuery(taskId string) (res dto.TransferQueryDTO, errDTO dto.ErrorDTO) {
	//data := make(map[string]string)
	//data["task_id"] = taskId
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.query_task")
	url = strings.Replace(url, "{task_id}", taskId, -1)
	//url = strings.Replace(url, "{task_id}", taskId, -1)
	resp, err := GetHeader(url, nil)
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
