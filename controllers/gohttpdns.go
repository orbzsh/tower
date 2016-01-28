package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"goblog/models"
)

type HttpdnsController struct {
	beego.Controller
}

func (self *HttpdnsController) Get() {
	self.Data["IsHdns"] = true
	self.TplNames = "gohttpdns.html"
}

func (self *HttpdnsController) Post() {
	domain := self.Input().Get("domain")
	dnss, err := models.ResolveFromRedis(domain)
	if err != nil {
		self.Ctx.WriteString(fmt.Sprint(err))
	}
	self.Ctx.WriteString(fmt.Sprint(dnss))
}
