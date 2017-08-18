# dingtalk-open
钉钉开放平台 GO SDK

[![Build Status](https://travis-ci.org/infoepoch/dingtalk-open.svg?branch=master)](https://travis-ci.org/infoepoch/dingtalk-open)
[![Go Report Card](https://goreportcard.com/badge/github.com/infoepoch/dingtalk-open)](https://goreportcard.com/report/github.com/infoepoch/dingtalk-open)

## 安装:
```
go get github.com/infoepoch/dingtalk-open
```

## 更新:
```
go get -u github.com/infoepoch/dingtalk-open  
```

## 快速使用案例

发送到群机器人信息

```Go
package main

import (
	"fmt"
	"github.com/infoepoch/dingtalk-open/webhook"
	"log"
)

// 钉钉机器人 token
const dingTalkBotAccessToken = ""

// 发送钉钉纯文本内容
func SendDingTalkText(text string) {
	str, err := webhook.SendText(dingTalkBotAccessToken, text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
```