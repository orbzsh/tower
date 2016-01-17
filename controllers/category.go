/*
分类的controller
*/
package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
)

type CategoryController struct {
	beego.Controller
}

func (self *CategoryController) Get() {
	op := self.Input().Get("op")
	switch op {
	case "add":
		name := self.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}
		self.Redirect("/category", 301)
		return
	case "del":
		id := self.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		self.Redirect("/category", 301)
		return
	}
	self.Data["IsCategory"] = true
	self.TplNames = "category.html"
	var err error
	self.Data["Categories"], err = models.GetAllCategory()
	if err != nil {
		beego.Error(err)
	}
}
