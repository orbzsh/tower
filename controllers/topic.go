package controllers

import (
	"github.com/astaxie/beego"
	"goblog/models"
)

type TopicController struct {
	beego.Controller
}

func (self *TopicController) Get() {
	self.Data["IsTopic"] = true
	self.TplNames = "topic.html"
	topics, err := models.GetAllTopics("", false)
	if err != nil {
		beego.Error(err)
	} else {
		self.Data["Topics"] = topics
	}
}

func (self *TopicController) Post() {
	tid := self.Input().Get("tid")
	title := self.Input().Get("title")
	content := self.Input().Get("content")
	category := self.Input().Get("category")

	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, category, content)
	} else {
		err = models.ModifyTopic(tid, title, category, content)
	}
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/topic", 302)
}

func (self *TopicController) Add() {
	self.TplNames = "topic_add.html"
}

func (self *TopicController) View() {
	self.TplNames = "topic_view.html"
	topic, err := models.GetTopic(self.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		self.Redirect("/", 302)
		return
	}
	self.Data["Topic"] = topic
	self.Data["Tid"] = self.Ctx.Input.Param("0")
	replies, err := models.GetAllReply(self.Ctx.Input.Param("0"))
	if err != nil {
		beego.Error(err)
		return
	}
	self.Data["Replies"] = replies

}

func (self *TopicController) Modify() {
	self.TplNames = "topic_modify.html"
	tid := self.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		self.Redirect("/", 302)
		return
	}
	self.Data["Topic"] = topic
	self.Data["Tid"] = tid
}

func (self *TopicController) Delete() {
	err := models.DeleteTopic(self.Input().Get("tid"))
	if err != nil {
		beego.Error(err)
	}
	self.Redirect("/", 302)
}
