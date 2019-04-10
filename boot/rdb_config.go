package boot

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func DbInit() {
	var cfg = beego.AppConfig

	// database
	dbDriver := cfg.String("dbdriver")
	dbUser := cfg.String("mysqluser")
	dbPass := cfg.String("mysqlpass")
	dbHost := cfg.String("mysqlhost")
	dbPort := cfg.String("mysqlport")
	dbName := cfg.String("mysqldb")
	maxIdleConn, _ := cfg.Int("maxidle")
	maxOpenConn, _ := cfg.Int("maxopen")

	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPass, dbHost, dbPort, dbName)

	err0 := orm.RegisterDriver(dbDriver, orm.DRMySQL)
	if err0 != nil {
		fmt.Println("register driver failed", dbDriver, err0)
	}
	err1 := orm.RegisterDataBase("default", dbDriver, dbLink, maxIdleConn, maxOpenConn)
	if err1 != nil {
		fmt.Println("register database failed", dbName, err1)
	}
	if HasTest == true {
		orm.Debug = true
	}
}
