/*
首页：index的controller
*/
package controllers

import (
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (self *IndexController) Get() {
	self.Ctx.WriteString("appname:" + beego.AppConfig.String("appname") + "\n")
	self.Ctx.WriteString("\n\n" + beego.AppName)

	beego.SetLevel(beego.LevelInformational)
	beego.Trace("trace index")
	beego.Info("trace info")
}
