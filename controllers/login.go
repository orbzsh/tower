/*
登录：login的controller
*/
package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (self *LoginController) Get() {
	self.TplNames = "login.html"
}

func (self *LoginController) Post() {
	self.Ctx.WriteString(fmt.Sprint(self.Input()))
	return
}
