package webhook

import (
	"log"
	"strconv"
	"testing"
	"time"
)

var accessToken = "11b53105e91cf109aaa16690a6a9177e5dfd148d86c6bfd8863f213952c9c985"

//// 测试发送文本内容结构体是否正常
//func TestTextStruct(t *testing.T) {
//
//}

// 测试封装接口
func TestSendText(t *testing.T) {
	i := time.Now().Unix()
	log.Println(i)
	SendText(accessToken, "单元测试 - 发送文本内容: 自动执行，忽略。"+strconv.FormatInt(i, 10))
}

func TestSendLink(t *testing.T) {
	i := time.Now().Unix()
	log.Println(i)
	SendLink(accessToken, "单元测试 - 自动化发送", "单元测试 - 发送文本内容: 自动执行，忽略。"+strconv.FormatInt(i, 10), "http://m.sh.189.cn/wx/dist/images/broad05-1b42d.png", "http://m.sh.189.cn")
}

func TestSendActionCard(t *testing.T) {
	i := time.Now().Unix()
	SendActionCard(accessToken, "TestSendActionCard-"+strconv.FormatInt(i, 10), "TestSendActionCardTitle", "0", "0", "点击阅读图文", "http://m.sh.189.cn")
}

func TestSendActionCardBtns(t *testing.T) {
	i := time.Now().Unix()
	//TODO 本案例中 数组和切片的问题需要深思下
	var btns [2]Btns

	btns[0].Title = "内容不错"
	btns[0].ActionURL = "http://m.sh.189.cn"
	btns[1].Title = "内容不好"
	btns[1].ActionURL = "http://m.sh.189.cn"

	btnss := btns[:]
	SendActionCardBtns(accessToken, "TestSendActionCardBtns-"+strconv.FormatInt(i, 10), "TestSendActionCardTitle", "0", "0", btnss)
}

func TestSendFeedCard(t *testing.T) {
	i := time.Now().Unix()
	//TODO 本案例中 数组和切片的问题需要深思下
	var links [2]Link

	links[0].Title = "时代的火车向前开1 - " + strconv.FormatInt(i, 10)
	links[0].MessageUrl = "http://m.sh.189.cn"
	links[0].PicUrl = "http://m.sh.189.cn/wx/dist/images/broad05-1b42d.png"

	links[1].Title = "时代的火车向前开2 - " + strconv.FormatInt(i, 10)
	links[1].MessageUrl = "http://m.sh.189.cn"
	links[1].PicUrl = "http://m.sh.189.cn/wx/dist/images/broad05-1b42d.png"

	linkss := links[:]
	SendFeedCard(accessToken, linkss)
}
