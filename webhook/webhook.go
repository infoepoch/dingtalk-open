package webhook

import (
	"github.com/infoepoch/dingtalk-open/httplib"
	"log"
)

/*
官方参考文档：
https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.8JWX03&treeId=257&articleId=105735&docType=1

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

type Btns struct {
	Title     string `json:"title"`
	ActionURL string `json:"actionURL"`
}

// link类型
type MsgLink struct {
	MsgType string `json:"msgtype"`
	Link    Link `json:"link"`
}

// link 类型详细
type Link struct {
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

// 整体跳转ActionCard类型
type MsgActionCard struct {
	MsgType    string `json:"msgtype"`
	ActionCard ActionCard `json:"actionCard"`
}

// 整体跳转ActionCard类型
type ActionCard struct {
	Text           string `json:"text"`
	Tile           string `json:"title"`
	HideAvatar     string `json:"hideAvatar"`
	BtnOrientation string `json:"btnOrientation"`
	SingleTitle    string `json:"singleTitle"`
	SingleURL      string `json:"singleURL"`
}

type MsgActionCardBtns struct {
	MsgType    string `json:"msgtype"`
	ActionCard ActionCardBtns `json:"actionCard"`
}

// 独立跳转ActionCard类型
type ActionCardBtns struct {
	Text           string `json:"text"`
	Tile           string `json:"title"`
	HideAvatar     string `json:"hideAvatar"`
	BtnOrientation string `json:"btnOrientation"`
	Btns           []Btns `json:"btns"`
}

// FeedCard类型
type MsgFeedCard struct {
	MsgType      string `json:"msgtype"`
	FeedCardLink FeedCardLink `json:"feedCard"`
}

type FeedCardLink struct {
	Links []Link `json:"links"`
}

// 发送文本内容
func SendText(accessToken string, content string) (string, error) {

	var msgText MsgText
	msgText.MsgType = "text"
	msgText.Text.Content = content

	req, err := httplib.Post(API + accessToken).JSONBody(msgText)
	log.Println(msgText)
	if err != nil {
		log.Fatal(err)
	}
	str, err1 := req.String()
	log.Println(str)
	return str, err1
}

// 发送链接地址
func SendLink(accessToken string, text string, title string, picUrl string, messageUrl string) (string, error) {
	var msgLink MsgLink
	msgLink.MsgType = "link"
	msgLink.Link.Text = text
	msgLink.Link.Title = title
	msgLink.Link.PicUrl = picUrl
	msgLink.Link.MessageUrl = messageUrl
	req, err := httplib.Post(API + accessToken).JSONBody(msgLink)
	log.Println(msgLink)
	if err != nil {
		log.Fatal(err)
	}

	str, err1 := req.String()
	log.Println(str)
	return str, err1
}

// 发送链接地址
/*
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
func SendActionCard(accessToken string, text string, title string, hideAvatar string, btnOrientation string, singleTitle string, singleURL string) (string, error) {
	var msg MsgActionCard
	msg.MsgType = "actionCard"
	msg.ActionCard.Text = text
	msg.ActionCard.Tile = title
	msg.ActionCard.HideAvatar = hideAvatar
	msg.ActionCard.BtnOrientation = btnOrientation
	msg.ActionCard.SingleTitle = singleTitle
	msg.ActionCard.SingleURL = singleURL

	req, err := httplib.Post(API + accessToken).JSONBody(msg)
	log.Println(msg)
	if err != nil {
		log.Fatal(err)
	}

	str, err1 := req.String()
	log.Println(str)
	return str, err1
}

/**
独立跳转ActionCard类型
{
    "actionCard": {
        "title": "乔布斯 20 年前想打造一间苹果咖啡厅，而它正是 Apple Store 的前身",
        "text": "![screenshot](@lADOpwk3K80C0M0FoA)
 ### 乔布斯 20 年前想打造的苹果咖啡厅
 Apple Store 的设计正从原来满满的科技感走向生活化，而其生活化的走向其实可以追溯到 20 年前苹果一个建立咖啡馆的计划",
        "hideAvatar": "0",
        "btnOrientation": "0",
        "btns": [
            {
                "title": "内容不错",
                "actionURL": "https://www.dingtalk.com/"
            },
            {
                "title": "不感兴趣",
                "actionURL": "https://www.dingtalk.com/"
            }
        ]
    },
    "msgtype": "actionCard"
}
 */
func SendActionCardBtns(accessToken string, text string, title string, hideAvatar string, btnOrientation string, btns []Btns) (string, error) {
	var msg MsgActionCardBtns
	msg.MsgType = "actionCard"
	msg.ActionCard.Text = text
	msg.ActionCard.Tile = title
	msg.ActionCard.HideAvatar = hideAvatar
	msg.ActionCard.BtnOrientation = btnOrientation
	msg.ActionCard.Btns = btns;
	req, err := httplib.Post(API + accessToken).JSONBody(msg)
	log.Println(msg)
	if err != nil {
		log.Fatal(err)
	}

	str, err1 := req.String()
	log.Println(str)
	return str, err1
}

// FeedCard类型
func SendFeedCard(accessToken string, links []Link) (string, error) {
	var msgLink MsgFeedCard
	msgLink.MsgType = "feedCard"
	msgLink.FeedCardLink.Links = links
	req, err := httplib.Post(API + accessToken).JSONBody(msgLink)
	log.Println(msgLink)
	if err != nil {
		log.Fatal(err)
	}

	str, err1 := req.String()
	log.Println(str)
	return str, err1
}
