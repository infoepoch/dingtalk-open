package webhook

import (
	"github.com/infoepoch/dingtalk-open/httplib"
	"log"
)

/*
发送文字
{
    "msgtype": "text",
    "text": {
        "content": "我就是我, 是不一样的烟火"
    },
    "at": {
        "atMobiles": [
            "156xxxx8827",
            "189xxxx8325"
        ],
        "isAtAll": false
    }
}

link 类型
{
    "msgtype": "link",
    "link": {
        "text": "这个即将发布的新版本，创始人陈航（花名“无招”）称它为“红树林”。
而在此之前，每当面临重大升级，产品经理们都会取一个应景的代号，这一次，为什么是“红树林”？",
        "title": "时代的火车向前开",
        "picUrl": "",
        "messageUrl": "https://mp.weixin.qq.com/s?__biz=MzA4NjMwMTA2Ng==&mid=2650316842&idx=1&sn=60da3ea2b29f1dcc43a7c8e4a7c97a16&scene=2&srcid=09189AnRJEdIiWVaKltFzNTw&from=timeline&isappinstalled=0&key=&ascene=2&uin=&devicetype=android-23&version=26031933&nettype=WIFI"
    }
}
 */

const API string = "https://oapi.dingtalk.com/robot/send?access_token="

// 具体的消息数字题
type MsgText struct {
	MsgType string `json:"msgtype"`
	Text    Text `json:"text"`
	At      At `json:"at"`
}

// 文本内容主体
type Text struct {
	Content string `json:"content"`
}

// @人的操作
type At struct {
	AtMobiles []string `json:"atMobiles"`
	IsAtAll   bool `json:"isAtAll"`
}

// link类型
type MsgLink struct {
	MsgType string `json:"msgtype"`
	Link    Link `json:"link"`
}

// link 类型详细
type Link struct {
	Text       string `json:"text"`
	Tile       string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

// 发送文本内容
func SendText(content string, access_token string) {

	var msgText MsgText
	msgText.MsgType = "text"
	msgText.Text.Content = content

	req, err := httplib.Post(API + access_token).JSONBody(msgText)
	log.Println(msgText)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(req.String())
}
