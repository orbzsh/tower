package routers

import (
	"github.com/astaxie/beego"
	"goblog/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
}
