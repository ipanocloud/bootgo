// @APIVersion 1.0.0
// @Title bootgo Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact zhoubin296@gmail.com
// @TermsOfServiceUrl https://gitee.com/ipanocloud/bootgo
// @License MIT
// @LicenseUrl https://gitee.com/ipanocloud/bootgo/blob/master/LICENSE
package routers

import (
	"gitee.com/ipanocloud/bootgo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
