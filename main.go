/*
Auth:orbzsh
version1：参考视频学习beego基本,做一个blog
*/
package main

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
	// "goblog/models"
	_ "goblog/routers"
)

// func init() {
// 	models.RegisterDB()
// }

func main() {
	// orm.Debug = true
	// orm.RunSyncdb("default", false, true)
	beego.Run()
}
