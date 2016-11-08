package controllers

import (
	"github.com/astaxie/beego"
	"strconv"
	"restsnowflake/models"
	"restsnowflake/utils"
)

func (c *AutoIDController) URLMapping() {

}

type AutoIDController struct {
	beego.Controller
}

type ReturnData struct {
	ID uint64  `json:"id"`
}

// @router /getid/:workerid [GET]
func (c *AutoIDController) GetID() {
	idStr := c.Ctx.Input.Param(":workerid")

	worker_id,err:=strconv.Atoi(idStr)

	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, "failed"+err.Error(), nil)
		c.ServeJSON()
		return
	}

	if worker_id<0 || worker_id>1023 {
		c.Data["json"] = models.GetReturnData(-1, "wrong workerid", nil)
		c.ServeJSON()
		return
	}

	sf,_ := utils.GetSnowFlakeObj(uint32(worker_id))

	rd := ReturnData{}
	id,_ :=sf.Next()

	rd.ID = id
	c.Data["json"] = models.GetReturnData(0, "success !", rd)
	c.ServeJSON()

}
