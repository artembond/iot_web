package main

import (
	"artem_test_net_http/iot"
	"artem_test_net_http/web"
)

func main() {
	iot.Init()
	web.Init()
}
