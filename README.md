# gin-slack

Gin framework middleware for reporting panic to Slack.

[![Travis branch](https://img.shields.io/travis/easonlin404/gin-slack/master.svg)](https://travis-ci.org/easonlin404/gin-slack)
[![Codecov branch](https://img.shields.io/codecov/c/github/easonlin404/gin-slack/master.svg)](https://codecov.io/gh/easonlin404/gin-slack)
[![Go Report Card](https://goreportcard.com/badge/github.com/easonlin404/gin-slack)](https://goreportcard.com/report/github.com/easonlin404/gin-slack)
![Slack Image](https://easonlin404.github.io/gin-slack/images/slack.png)

## Features
* todo

## Installation
```sh
$ go get -u github.com/easonlin404/gin-slack/
```
## Usage
```go
package main

import (
	"github.com/easonlin404/gin-slack"
	"github.com/easonlin404/go-slack"
	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.New()

	// You need get your slack WebhookURL from https://api.slack.com/incoming-webhooks
	s := slack.New().WebhookURL("https://hooks.slack.com/services/T00000000/B00000000/XXXXXXXXXXXXXXXXXXXXXXXX")

	// Use ginslack middleware
	r.Use(ginslack.Recovery(s))

	// Other your handler	
	// r.Get("get",....)
	
	r.Run()
}


```
