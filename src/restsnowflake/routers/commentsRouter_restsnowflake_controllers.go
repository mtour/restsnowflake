package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["restsnowflake/controllers:AutoIDController"] = append(beego.GlobalControllerRouter["restsnowflake/controllers:AutoIDController"],
		beego.ControllerComments{
			"GetID",
			`/getid/:workerid`,
			[]string{"GET"},
			nil})

}
