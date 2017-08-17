package sign

import (
	"log"
	"testing"
	"time"
)

func TestGetJsapiSign(t *testing.T) {
	i := time.Now().Unix()
	log.Println(i)
	str := GetJsapiSign(JsSign{Ticket: "123123", TimeStamp: "1502939123387", NonceStr: "abcdefg", Url: "http://www.baidu.com"})
	log.Println("TestGetJsapiSign 加密后结果：" + str)
	if str == "754432c2626ef2688503797d1799a97fd5493d23" {
		t.Log("TestGetJsapiSign 测试成功")
	} else {
		t.Error("TestGetJsapiSign 测试失败")
	}
}
