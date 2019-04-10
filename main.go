package main

import (
	"gitee.com/ipanocloud/bootgo/boot"
	_ "gitee.com/ipanocloud/bootgo/routers"
	"github.com/astaxie/beego"
)

func main() {
	boot.InitEnv()
	beego.Run()
}
