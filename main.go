/*
Auth:orbzsh
version1：参考视频学习beego基本,做一个blog
*/
package main

import (
	"github.com/astaxie/beego"
	_ "goblog/routers"
)

func main() {
	beego.Run()
}
