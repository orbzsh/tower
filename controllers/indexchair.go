package controllers

import (
	"github.com/astaxie/beego"
)

type IndexChairController struct {
	beego.Controller
}

func (self *IndexChairController) Get() {
	self.Data["IsIndex"] = true
	self.TplNames = "indexchair.html"
}
