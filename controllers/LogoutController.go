package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}


func (c *LogoutController) URLMapping() {
c.Mapping("index", c.Index)
}

// @router /logout/:key [get]
func (this *LogoutController) Index(){

	this.Ctx.WriteString("===logout")

}
