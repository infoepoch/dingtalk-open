package webhook

import (
	"github.com/infoepoch/dingtalk-open/httplib"
	"log"
)

var AccessToken = ""

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
整体跳转ActionCard类型
{
    "actionCard": {
        "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
        "text": "![screenshot](@lADOpwk3K80C0M0FoA)
 ### 乔布斯 20 年前想打造的苹果咖啡厅
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
        "hideAvatar": "0",
        "btnOrientation": "0",
        "singleTitle" : "阅读全文",
        "singleURL" : "https://www.dingtalk.com/"
    },
    "msgtype": "actionCard"
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

// 整体跳转ActionCard类型
type MsgActionCard struct {
	MsgType    string `json:"msgtype"`
	ActionCard ActionCard `json:"actionCard"`
}

type ActionCard struct {
	Text           string `json:"text"`
	Tile           string `json:"title"`
	HideAvatar     string `json:"hideAvatar"`
	BtnOrientation string `json:"btnOrientation"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
}

// 发送文本内容
func SendText(content string) (string, error) {

	var msgText MsgText
	msgText.MsgType = "text"
	msgText.Text.Content = content

	req, err := httplib.Post(API + AccessToken).JSONBody(msgText)
	log.Println(msgText)
	if err != nil {
		log.Fatal(err)
	}
	str, err1 := req.String()
	log.Println(str)
	return str, err1
}

// 发送链接地址
func SendLink(text string, title string, picUrl string, messageUrl string) (string, error) {
	var msgLink MsgLink
	msgLink.MsgType = "link"
	msgLink.Link.Text = text
	msgLink.Link.Tile = title
	msgLink.Link.PicUrl = picUrl
	msgLink.Link.MessageUrl = messageUrl
	req, err := httplib.Post(API + AccessToken).JSONBody(msgLink)
	log.Println(msgLink)
	if err != nil {
		log.Fatal(err)
	}

	str, err1 := req.String()
	log.Println(str)
	return str, err1
}

// 发送链接地址
func SendActionCard(text string, title string, hideAvatar string, btnOrientation string, singleTitle string, singleURL string) (string, error) {
	var msgLink MsgActionCard
	msgLink.MsgType = "actionCard"
	msgLink.ActionCard.Text = text
	msgLink.ActionCard.Tile = title
	msgLink.ActionCard.HideAvatar = hideAvatar
	msgLink.ActionCard.BtnOrientation = btnOrientation
	msgLink.ActionCard.SingleTitle = singleTitle
	msgLink.ActionCard.SingleURL = singleURL

	req, err := httplib.Post(API + AccessToken).JSONBody(msgLink)
	log.Println(msgLink)
	if err != nil {
		log.Fatal(err)
	}

	str, err1 := req.String()
	log.Println(str)
	return str, err1
}
