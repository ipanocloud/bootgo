package bootconfig

import (
	"fmt"
	"github.com/astaxie/beego"
	"os"
)

var HasTest bool

//初始化启动环境
func EnvInit() {

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
		os.Exit(1)
	}

	//开发环境则开启swagger
	if env == "dev" || env == "qa" {
		HasTest = true
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
}
