package server

import (
	"github.com/astaxie/beego/logs"
	"github.com/infoepoch/dingtalk-open/utils"
)

type DingOpenResult struct {
	Ding_open_errcode int    `json:"ding_open_errcode"`
	Error_msg         string `json:" error_msg"`
	Success           bool   `json:" success"`
	Task_id           int    `json:"  task_id"`
}

// dingtalk.corp.message.corpconversation.asyncsend (企业会话消息异步发送)
// 企业可以主动发消息给员工，此接口是/message/send的异步化版本，性能更好，开发者要迁移到此接口上。 目前支持text、image、voice、file、link、OA消息类型。 当前ISV给同一用户发消息一天不得超过50次。企业开发者每分钟最多可调用1500次，ISV开发者每分钟最多可调用200次。

//curl -X POST "https://eco.taobao.com/router/rest" \
//-H 'Content-Type:application/x-www-form-urlencoded;charset=utf-8' \
//-d 'format=json' \
//-d 'method=dingtalk.corp.message.corpconversation.asyncsendbycode' \
//-d 'session=385d9c05-4dc1-41b6-9fe6-e3fda6dea06f' \
//-d 'timestamp=2017-08-19+12%3A15%3A26' \
//-d 'v=2.0' \
//-d 'agent_id=1234' \
//-d 'code=7dsf78d78sf87sd' \
//-d 'dept_id_list=123%2C456' \
//-d 'msgcontent=' \
//-d 'msgtype=oa' \
//-d 'to_all_user=false' \
//-d 'user_id_list=zhangsan%2Clisi'

func Asyncsend(access_token string, method string, msgtype string, agent_id string, user_id_list string, dept_id_list string, to_all_user string, msgcontent string, code string) (DingOpenResult, error) {
	var result DingOpenResult
	req := utils.GetEcoReq(access_token, method)
	req.Param("msgtype", msgtype)
	req.Param("agent_id", agent_id)
	req.Param("user_id_list", user_id_list)
	req.Param("dept_id_list", dept_id_list)
	req.Param("to_all_user", to_all_user)
	req.Param("msgcontent", msgcontent)
	req.Param("code", code)
	str, err := req.String()
	if err != nil {
		logs.Error(err.Error())
	} else {
		err = utils.TransStrToJSON(str, &result)
		if err != nil {
			logs.Error(err.Error())
		}
	}
	return result, err
}
