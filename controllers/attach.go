package controllers

import (
	"github.com/astaxie/beego"
	"io"
	"net/url"
	"os"
)

type AttachController struct {
	beego.Controller
}

func (self *AttachController) Get() {
	filePath, err := url.QueryUnescape(self.Ctx.Request.RequestURI[1:])
	if err != nil {
		self.Ctx.WriteString(err.Error())
		return
	}
	f, err := os.Open(filePath)
	if err != nil {
		self.Ctx.WriteString(err.Error())
		return
	}
	defer f.Close()
	_, err = io.Copy(self.Ctx.ResponseWriter, f)
	if err != nil {
		self.Ctx.WriteString(err.Error())
		return
	}
}
