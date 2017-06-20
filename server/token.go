package server

import (
	"github.com/infoepoch/dingtalk-open/config"
	"github.com/infoepoch/dingtalk-open/httplib"
)



// 获取AccessToken
// https://open-doc.dingtalk.com/docs/doc.htm?spm=a219a.7629140.0.0.0tLAGF&treeId=172&articleId=104980&docType=1#s2
func GetAccessToken() (string, error) {
	req := httplib.Get(config.BaseUrl + "/gettoken?corpid=" + config.CORP_ID + "&corpsecret=" + config.SSO_SECRECT)
	str, err := req.String()
	return str, err
}
