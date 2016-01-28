/*
定义路由
*/
package routers

import (
	"github.com/astaxie/beego"
	"goblog/controllers"
)

func init() {
	// beego.Router("/", &controllers.IndexController{})
	// beego.Router("/login", &controllers.LoginController{})
	// beego.Router("/category", &controllers.CategoryController{})
	// beego.Router("/reply", &controllers.ReplyController{})
	// beego.Router("/reply/add", &controllers.ReplyController{}, "post:Add")
	// beego.Router("/reply/del", &controllers.ReplyController{}, "get:Del")
	// beego.Router("/topic", &controllers.TopicController{})
	// beego.AutoRouter(&controllers.TopicController{})
	// beego.SetStaticPath("/data", "data")
	// // beego.Router("/data/:all", &controllers.AttachController{})
	beego.Router("/", &controllers.IndexChairController{})
}
