package main

import (
	"fmt"
	_ "gitee.com/ipanocloud/bootgo/routers"
	"github.com/astaxie/beego"
	"os"
)

func main() {
	var appConfig string
	env := os.Getenv("ENV_CLUSTER")
	if env == "prod" {
		appConfig = "./conf/prod.conf"
	} else if env == "stg" {
		appConfig = "./conf/stg.conf"
	} else if env == "qa" {
		appConfig = "./conf/qa.conf"
	} else if env == "dev" {
		appConfig = "./conf/dev.conf"
	}
	err := beego.LoadAppConfig("ini", appConfig)
	if err != nil {
		fmt.Println("load conf failed,", err)
	}

	if env == "dev" || env == "qa" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
