package routers

import (
	"restsnowflake/controllers"
	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("snowflake",
			beego.NSInclude(
				&controllers.AutoIDController{},
			),
	)
	beego.AddNamespace(ns)
}
