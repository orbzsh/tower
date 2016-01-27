package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
)

type ReplyController struct {
	beego.Controller
}

func (self *ReplyController) Add() {
	tid := self.Input().Get("tid")
	err := models.AddReply(
		tid,
		self.Input().Get("nickname"),
		self.Input().Get("content"),
	)
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/topic/view/"+tid, 302)
}

func (self *ReplyController) Del() {
	tid := self.Input().Get("tid")
	err := models.DelReply(self.Input().Get("rid"))
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/topic/view/"+tid, 302)
}
