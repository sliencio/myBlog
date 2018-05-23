package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	//if this.Data["isLogin"] != true{
	//	this.Redirect("/login",301)
	//}
	this.TplName = "index.html"
}
