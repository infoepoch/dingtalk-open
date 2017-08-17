package utils

import (
	"testing"
)

type ErrCodeMsg struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type AccessToken struct {
	ErrCode     int    `json:"errcode"`
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
}

func TestTransStrToJSON1(t *testing.T) {
	str := `{"errcode": 0,"errmsg": "ok","access_token": "fw8ef8we8f76e6f7s8df8s"}`
	t.Log(str)
	var accessToken AccessToken
	err := TransStrToJSON(str, &accessToken)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(accessToken.ErrCode)
	t.Log(accessToken.ErrMsg)
	t.Log(accessToken.AccessToken)
}

func TestTransStrToJSON2(t *testing.T) {
	str := `{"errcode": 43003,"errmsg": "require https"}`
	t.Log(str)
	var errCode ErrCodeMsg
	err := TransStrToJSON(str, &errCode)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(errCode.ErrCode)
	t.Log(errCode.ErrMsg)

}

func TestTransStrToJSON3(t *testing.T) {
	str := `{"errcode": 43003,"errmsg": "require https"}`
	t.Log(str)
	var accessToken AccessToken
	err := TransStrToJSON(str, &accessToken)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(accessToken.ErrCode)
	t.Log(accessToken.ErrMsg)
	t.Log(accessToken.AccessToken)
}
