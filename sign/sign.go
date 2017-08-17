package sign

import (
	"crypto/sha1"
	"encoding/hex"
)

// js 授权加密
type JsSign struct {
	Ticket    string
	NonceStr  string
	TimeStamp string
	Url       string
}

// noncestr=Zn4zmLFKD0wzilzM
// jsapi_ticket=mS5k98fdkdgDKxkXGEs8LORVREiweeWETE40P37wkidkfksDSKDJFD5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcKIDU8l
// timestamp=1414588745
// url=//open.dingtalk.com

func GetJsapiSign(sign JsSign) string {
	plain := "jsapi_ticket=" + sign.Ticket + "&noncestr=" + sign.NonceStr +
		"&timestamp=" + sign.TimeStamp + "&url=" + sign.Url
	r := sha1.Sum([]byte(plain))
	return hex.EncodeToString(r[:])
}
