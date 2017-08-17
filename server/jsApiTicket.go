package server

import (
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/infoepoch/dingtalk-open/config"
	"github.com/infoepoch/dingtalk-open/utils"
)

//获取jsapi_ticket
//
//企业在使用微应用中的JS API时，需要先从钉钉开放平台接口获取jsapi_ticket生成签名数据，并将最终签名用的部分字段及签名结果返回到H5中，JS API底层将通过这些数据判断H5是否有权限使用JS API。
//
//注意：1.在企业应用中，在jsticket未过期的时候通过get_jsapi_ticket获取到的都是一个全新的jsticket（和旧的jsticket值不同），这个全新的jsticket过期时间是2小时。2.在ISV应用中，在jsticket未过期的时候，也就是两小时之内通过get_jsapi_ticket获取到的jsticket和老的jsticket值相同，只是过期时间延长到2小时。3.jsticket是以一个企业对应一个，所以在使用的时候需要将jsticket以企业为维度进行缓存下来（设置缓存过期时间2小时），并不需要每次都通过接口拉取。
//
//由于很多开发者对该接口的使用原理不了解,所以带来了很多问题。
//请务必查看:JSAPI鉴权的正确使用方式
//
//请求说明
//
//Https请求方式：GET
//
//https://oapi.dingtalk.com/get_jsapi_ticket?access_token=ACCESS_TOKE

//{
//"errcode": 0,
//"errmsg": "ok",
//"ticket": "dsf8sdf87sd7f87sd8v8ds0vs09dvu09sd8vy87dsv87",
//"expires_in": 7200
//}

type JsApiToken struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Ticket    string `json:"ticket"`
	ExpiresIn int    `json:"expires_in"`
}

// 获取 js token
func GetJsApiToken(access_token string) (JsApiToken, error) {
	req := httplib.Get(config.BaseUrl + "/get_jsapi_ticket?access_token=" + access_token)
	str, err := req.String()
	var token JsApiToken
	if err != nil {
		logs.Error(err.Error())
	} else {
		err = utils.TransStrToJSON(str, &token)
		if err != nil {
			logs.Error(err.Error())
		}
	}
	return token, err
}
