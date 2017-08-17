package server

import (
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"github.com/infoepoch/dingtalk-open/config"
	"github.com/infoepoch/dingtalk-open/utils"
)

//{
//    "errcode": 0,
//    "errmsg": "ok",
//    "userid": "USERID",
//    "deviceId":"DEVICEID",
//    "is_sys": true,
//    "sys_level": 0|1|2
//}

type CodeUserInfo struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	Userid   string `json:"userid"`
	DeviceId string `json:"deviceId"`
	IsSys    bool   `json:"is_sys"`
	SysLevel int    `json:"sys_level"`
}

//{
//"errcode": 0,
//"errmsg": "ok",
//"userid": "zhangsan",
//"name": "张三",
//"tel": "010-123333",
//"workPlace":"",
//"remark": "",
//"mobile": "13800000000",
//"email": "dingding@aliyun.com",
//"active": true,
//"orderInDepts": "{1:10, 2:20}",
//"isAdmin": false,
//"isBoss": false,
//"dingId": "WsUDaq7DCVIHc6z1GAsYDSA",
//"unionid": "cdInjDaq78sHYHc6z1gsz",
//"isLeaderInDepts": "{1:true, 2:false}",
//"isHide": false,
//"department": [1, 2],
//"position": "工程师",
//"avatar": "dingtalk.com/abc.jpg",
//"jobnumber": "111111"
//}

type UserInfo struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	Userid    string `json:"userid"`
	Name      string `json:"name"`
	Tel       string `json:"tel"`
	WorkPlace string `json:"workPlace"`
	Remark    string `json:"remark"`
	Mobile    string `json:"mobile"`
	Email     string `json:"email"`
	Active    bool   `json:"active"`
	IsAdmin   bool   `json:"isAdmin"`
	IsBoss    bool   `json:"isBoss"`
	DingId    string `json:"dingId"`
	Unionid   string `json:"unionid"`
	Position  string `json:"position"`
	Avatar    string `json:"avatar"`
	Jobnumber string `json:"jobnumber"`
}

// 通过CODE换取用户身份
func GetBaseByCode(access_token string, code string) (CodeUserInfo, error) {
	var codeUser CodeUserInfo
	req := httplib.Get(config.BaseUrl + "/user/getuserinfo?access_token=" + access_token + "&code=" + code)
	str, err := req.String()
	if err != nil {
		logs.Error(err.Error())
	} else {
		err = utils.TransStrToJSON(str, &codeUser)
		if err != nil {
			logs.Error(err.Error())
		}
	}
	return codeUser, err
}

// 获取成员详情
func GetUserByUserid(access_token string, userid string) (UserInfo, error) {
	var user UserInfo
	req := httplib.Get(config.BaseUrl + "/user/get?access_token=" + access_token + "&userid=" + userid)
	str, err := req.String()
	if err != nil {
		logs.Error(err.Error())
	} else {
		err = utils.TransStrToJSON(str, &user)
		if err != nil {
			logs.Error(err.Error())
		}
	}
	return user, err
}
