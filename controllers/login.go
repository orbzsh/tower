/*
登录：login的controller
*/
package controllers

import (
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (self *LoginController) Get() {
	self.TplNames = "login.html"
}

func (self *LoginController) Post() {
	uname := self.Input().Get("uname")
	pwd := self.Input().Get("pwd")
	autologin := self.Input().Get("autologin") == "on"

	if beego.AppConfig.String("uname") == uname && beego.AppConfig.String("pwd") == pwd {
		maxAge := 0
		if autologin {
			maxAge = 1<<31 - 1
		}
		self.Ctx.SetCookie("uname", uname, maxAge, "/")
		self.Ctx.SetCookie("pwd", pwd, maxAge, "/")
	}
	self.Redirect("/", 301)
	return
}
