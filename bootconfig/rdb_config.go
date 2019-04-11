package bootconfig

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

//初始化db配置
func DbInit() {
	var cfg = beego.AppConfig

	// database
	dbDriver := cfg.String("orm.database.driver")
	dbUser := cfg.String("orm.mysql.user")
	dbPass := cfg.String("orm.mysql.pass")
	dbHost := cfg.String("orm.mysql.host")
	dbPort := cfg.String("orm.mysql.port")
	dbName := cfg.String("orm.mysql.db")
	maxIdleConn, _ := cfg.Int("orm.maxIdle")
	maxOpenConn, _ := cfg.Int("orm.maxOpen")

	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)

	err := orm.RegisterDriver(dbDriver, orm.DRMySQL)
	if err != nil {
		fmt.Println("register driver failed", dbDriver, err)
		os.Exit(1)
	}
	err = orm.RegisterDataBase("default", dbDriver, dbLink, maxIdleConn, maxOpenConn)
	if err != nil {
		fmt.Println("register database failed", dbName, err)
		os.Exit(1)
	}
	if HasTest == true {
		orm.Debug = true
	}
}
