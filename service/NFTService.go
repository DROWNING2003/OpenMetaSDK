package service

import (
	"GoForAvataAPIDemo/models/dto"
	"GoForAvataAPIDemo/models/req"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"strings"
)

// NFTClassesCreate 创建 NFT 类别
func NFTClassesCreate(param *req.NFTClassCreateReq) (res dto.NFTClassCreateDTO, errDTO dto.ErrorDTO) {
	bytesData, err := json.Marshal(param)
	if err != nil {
		errDTO.Err = err
		return
	}
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.create_nft_classes")
	resp, err := PostHeader(url, bytesData)
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

// NFTClassesQuery 查询 NFT 类别
func NFTClassesQuery(param *req.NFTClassQueryReq) (res dto.NFTClassQueryDTO, errDTO dto.ErrorDTO) {
	p, _ := json.Marshal(&param)
	data := make(map[string]string)
	err := json.Unmarshal(p, &data)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.query_nft_classes")
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

// NFTClassesDetailQuery 查询 NFT 类别详情
func NFTClassesDetailQuery(classId string) (res dto.NFTClassDetailQueryDTO, errDTO dto.ErrorDTO) {
	url := viper.GetString("nft.domain") + strings.Replace(viper.GetString("nft.api.url.query_nft_classes_detail"), "{id}", classId, -1)
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

// NFTClassesTransfer 转让 NFT 类别
func NFTClassesTransfer(param *req.NFTClassTransferReq) (res dto.NFTClassTransferDTO, errDTO dto.ErrorDTO) {
	if param.PathParam.ClassId == "" {
		errDTO.Err = errors.New("classId mast be provide")
		return
	}
	if param.PathParam.Owner == "" {
		errDTO.Err = errors.New("owner mast be provide")
		return
	}
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.transfers_nft_classes")
	url = strings.Replace(url, "{class_id}", param.PathParam.ClassId, -1)
	url = strings.Replace(url, "{owner}", param.PathParam.Owner, -1)
	bytesData, err := json.Marshal(param.RequestBody)
	if err != nil {
		errDTO.Err = err
		return
	}

	resp, err := PostHeader(url, bytesData)
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

// NFTCreate 发行 NFT
func NFTCreate(param *req.NFTCreateReq) (res dto.NFTCreateDTO, errDTO dto.ErrorDTO) {
	if param.PathParam.ClassId == "" {
		errDTO.Err = errors.New("classId mast be provide")
		return
	}
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.create_nft")
	url = strings.Replace(url, "{class_id}", param.PathParam.ClassId, -1)
	bytesData, err := json.Marshal(param.RequestBody)
	if err != nil {
		errDTO.Err = err
		return
	}

	resp, err := PostHeader(url, bytesData)
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

// NFTTransfer 转让 NFT
func NFTTransfer(param *req.NFTTransferReq) (res dto.NFTTransferDTO, errDTO dto.ErrorDTO) {
	if param.PathParam.ClassId == "" {
		errDTO.Err = errors.New("classId mast be provide")
		return
	}
	if param.PathParam.Owner == "" {
		errDTO.Err = errors.New("owner mast be provide")
		return
	}
	if param.PathParam.NftId == "" {
		errDTO.Err = errors.New("nft_id mast be provide")
		return
	}
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.transfers_nft")
	url = strings.Replace(url, "{class_id}", param.PathParam.ClassId, -1)
	url = strings.Replace(url, "{owner}", param.PathParam.Owner, -1)
	url = strings.Replace(url, "{nft_id}", param.PathParam.NftId, -1)
	bytesData, err := json.Marshal(param.RequestBody)
	if err != nil {
		errDTO.Err = err
		return
	}

	resp, err := PostHeader(url, bytesData)
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

// NFTEdit 转让 NFT
func NFTEdit(param *req.NFTEditReq) (res dto.NFTEditDTO, errDTO dto.ErrorDTO) {
	if param.PathParam.ClassId == "" {
		errDTO.Err = errors.New("classId mast be provide")
		return
	}
	if param.PathParam.Owner == "" {
		errDTO.Err = errors.New("owner mast be provide")
		return
	}
	if param.PathParam.NftId == "" {
		errDTO.Err = errors.New("nft_id mast be provide")
		return
	}
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.edit_nft")
	url = strings.Replace(url, "{class_id}", param.PathParam.ClassId, -1)
	url = strings.Replace(url, "{owner}", param.PathParam.Owner, -1)
	url = strings.Replace(url, "{nft_id}", param.PathParam.NftId, -1)
	bytesData, err := json.Marshal(param.RequestBody)
	if err != nil {
		errDTO.Err = err
		return
	}

	resp, err := PatchHeader(url, bytesData)
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

// NFTDelete 转让 NFT
func NFTDelete(param *req.NFTDeleteReq) (res dto.NFTDeleteDTO, errDTO dto.ErrorDTO) {
	if param.PathParam.ClassId == "" {
		errDTO.Err = errors.New("classId mast be provide")
		return
	}
	if param.PathParam.Owner == "" {
		errDTO.Err = errors.New("owner mast be provide")
		return
	}
	if param.PathParam.NftId == "" {
		errDTO.Err = errors.New("nft_id mast be provide")
		return
	}
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.edit_nft")
	url = strings.Replace(url, "{class_id}", param.PathParam.ClassId, -1)
	url = strings.Replace(url, "{owner}", param.PathParam.Owner, -1)
	url = strings.Replace(url, "{nft_id}", param.PathParam.NftId, -1)
	bytesData, err := json.Marshal(param.RequestBody)
	if err != nil {
		errDTO.Err = err
		return
	}

	resp, err := DeleteHeader(url, bytesData)
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

// NFTQuery 查询 NFT
func NFTQuery(param *req.NFTQueryReq) (res dto.NFTQueryDTO, errDTO dto.ErrorDTO) {
	p, _ := json.Marshal(&param)
	data := make(map[string]string)
	err := json.Unmarshal(p, &data)
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.query_nft")
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

// NFTDetailQuery 查询 NFT 详情
func NFTDetailQuery(classId string, nftId string) (res dto.NFTDetailQueryDTO, errDTO dto.ErrorDTO) {

	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.query_nft_detail")
	url = strings.Replace(url, "{class_id}", classId, -1)
	url = strings.Replace(url, "{nft_id}", nftId, -1)
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

// NFTHistoryQuery 查询 NFT 操作记录
func NFTHistoryQuery(param *req.NFTHistoryQueryReq) (res dto.NFTHistoryQueryDTO, errDTO dto.ErrorDTO) {
	if param.PathParam.ClassId == "" {
		errDTO.Err = errors.New("classId mast be provide")
	}
	if param.PathParam.NftId == "" {
		errDTO.Err = errors.New("nftId mast be provide")
	}
	url := viper.GetString("nft.domain") + viper.GetString("nft.api.url.query_nft_history")
	url = strings.Replace(url, "{class_id}", param.PathParam.ClassId, -1)
	url = strings.Replace(url, "{nft_id}", param.PathParam.NftId, -1)

	p, _ := json.Marshal(&param.RequestBody)
	data := make(map[string]string)
	err := json.Unmarshal(p, &data)

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
