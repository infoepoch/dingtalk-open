package utils

import (
	"github.com/infoepoch/dingtalk-open/config"
	"github.com/infoepoch/dingtalk-open/httplib"
)

// ECO 请求基本参数封装

//时间戳，格式为yyyy-MM-dd HH:mm:ss，时区为GMT+8，例如：2015-01-01 12:00:00。淘宝API服务端允许客户端请求最大时间误差为10分钟。
func GetTimestamp() string {
	return ""
}

//响应格式。默认为xml格式，可选值：xml，json。
func GetFormat() string {
	return "json"
}

// API协议版本，可选值：2.0。
func GetV() string {
	return "2.0"
}

// 返回通用请求
func GetEcoReq(access_token string, method string) *httplib.BeegoHTTPRequest {
	req := httplib.Get(config.EcoUrl)
	req.Header("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	req.Param("fromat", GetFormat())
	req.Param("method", method)
	req.Param("session", access_token)
	req.Param("timestamp", GetTimestamp())
	req.Param("v", GetV())
	return req
}
