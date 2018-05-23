package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type ArticleController struct{
	beego.Controller
}

func(this* ArticleController)Get(){
	content :=this.GetString("content")
	fmt.Println("-------------",content)
	this.Data["showArticle"] = content
	this.TplName = "showArticle.html"
}

