package webhook

import (
	"testing"
	"log"
	"time"
	"strconv"
)

func init() {
	AccessToken = "11b53105e91cf109aaa16690a6a9177e5dfd148d86c6bfd8863f213952c9c985"
}

//// 测试发送文本内容结构体是否正常
//func TestTextStruct(t *testing.T) {
//
//}

// 测试封装接口
func TestSendText(t *testing.T) {
	i := time.Now().Unix()
	log.Println(i)
	SendText("单元测试 - 发送文本内容: 自动执行，忽略。" + strconv.FormatInt(i, 10))
}

func TestSendLink(t *testing.T) {
	i := time.Now().Unix()
	log.Println(i)
	SendLink("单元测试 - 自动化发送", "单元测试 - 发送文本内容: 自动执行，忽略。"+strconv.FormatInt(i, 10), "http://m.sh.189.cn/wx/dist/images/broad05-1b42d.png", "http://m.sh.189.cn")
}
