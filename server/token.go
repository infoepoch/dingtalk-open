package server

import (
	"github.com/astaxie/beego/logs"
	"github.com/infoepoch/dingtalk-open/config"
	"github.com/infoepoch/dingtalk-open/httplib"
	"github.com/infoepoch/dingtalk-open/utils"
)

//获取AccessToken
//
//AccessToken是企业访问钉钉开放平台接口的全局唯一票据，调用接口时需携带AccessToken。
//
//AccessToken需要用CorpID和CorpSecret来换取，不同的CorpSecret会返回不同的AccessToken。正常情况下AccessToken有效期为7200秒，有效期内重复获取返回相同结果，并自动续期。
//
//请求说明
//
//Https请求方式: GET
//https://oapi.dingtalk.com/gettoken?corpid=id&corpsecret=secrect
//
//参数说明
//
//参数	参数类型	必须	说明
//corpid	String	是	企业Id
//corpsecret	String	是	企业应用的凭证密钥
//返回说明
//
//a)正确的Json返回结果:
//
//{
//"errcode": 0,
//"errmsg": "ok",
//"access_token": "fw8ef8we8f76e6f7s8df8s"
//}
//参数	说明
//errcode	错误码
//errmsg	错误信息
//access_token	获取到的凭证
//b)错误的Json返回示例:
//
//
//{
//"errcode": 43003,
//"errmsg": "require https"
//}

type AccessToken struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
}

// 获取 access_token
func GetAccessToken(corpId string, corpSecret string) (AccessToken, error) {
	var accessToken AccessToken
	req := httplib.Get(config.BaseUrl + "/gettoken?corpid=" + corpId + "&corpsecret=" + corpSecret)
	str, err := req.String()
	if err != nil {
		logs.Error(err.Error())
	} else {
		err = utils.TransStrToJSON(str, &accessToken)
		if err != nil {
			logs.Error(err.Error())
		}
	}
	return accessToken, err
}
