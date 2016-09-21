package main

import (
	"github.com/artembond/iot_web/iot"
	"github.com/artembond/iot_web/web"
)

func main() {
	iot.Init()
	web.Init()
}
