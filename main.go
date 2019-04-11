package main

import (
	"gitee.com/ipanocloud/bootgo/bootconfig"
	_ "gitee.com/ipanocloud/bootgo/routers"
	"github.com/astaxie/beego"
)

func main() {
	bootconfig.InitEnv()
	beego.Run()
}
