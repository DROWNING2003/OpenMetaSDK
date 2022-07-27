package service

import (
	"GoForAvataAPIDemo/utils"
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

func PostHeader(url string, param []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(param))
	if err != nil {
		return nil, err
	}
	//验签
	reqS := utils.SignRequest(req, viper.GetString("nft.apiKey"), viper.GetString("nft.apiSecret"))
	reqS.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(reqS)
	return resp, err
}

func GetHeader(url string, param map[string]string) (*http.Response, error) {
	client := &http.Client{}
	//提交请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()

	for k, v := range param {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	//验签
	reqS := utils.SignRequest(req, viper.GetString("nft.apiKey"), viper.GetString("nft.apiSecret"))

	resp, _ := client.Do(reqS)

	return resp, nil
}

func PatchHeader(url string, param []byte) (*http.Response, error) {
	client := &http.Client{}
	//提交请求
	req, err := http.NewRequest("PATCH", url, bytes.NewReader(param))
	if err != nil {
		return nil, err
	}

	//验签
	reqS := utils.SignRequest(req, viper.GetString("nft.apiKey"), viper.GetString("nft.apiSecret"))

	resp, _ := client.Do(reqS)

	return resp, nil
}

func DeleteHeader(url string, param []byte) (*http.Response, error) {
	client := &http.Client{}
	//提交请求
	req, err := http.NewRequest("DELETE", url, bytes.NewReader(param))
	if err != nil {
		return nil, err
	}

	//验签
	reqS := utils.SignRequest(req, viper.GetString("nft.apiKey"), viper.GetString("nft.apiSecret"))

	resp, _ := client.Do(reqS)

	return resp, nil
}
