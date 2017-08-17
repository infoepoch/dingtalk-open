package config

import (
	"fmt"
	"runtime"
)

// 钉钉开发平台地址
const BaseUrl string = "https://oapi.dingtalk.com/"

var version = "1.0.0"

const (
	ctypeAppName = ""
)

func SetAppName(userApp string) error {
	fmt.Sprintf(
		"DingTalkGo/%s (%s; %s; %s) %s", version, runtime.GOOS, runtime.GOARCH, userApp, runtime.Version())
	return nil
}

func init() {
	SetAppName("")
}
