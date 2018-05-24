package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
	"myBlog/DB"
)

type ArticleController struct{
	beego.Controller
}

func(this* ArticleController)Get(){
	this.TplName = "index.html"
}
//编辑
func(this* ArticleController)Edit(){
	this.TplName = "editor.html"
	fmt.Println("-----Edit----")
}
//展示
func(this* ArticleController)Show(){
	objId := this.GetString("_id")
	fmt.Println("-----Show----"+objId)
	this.Data["showArticle"] = "ddddd"
	this.TplName = "showArticle.html"
}
//发布
func(this* ArticleController)Publish(){
	objId:=bson.NewObjectIdWithTime(time.Now())
	DB.Insert("articles",bson.M{
		"_id":objId,
		"uid":objId,
		"title":"test",
		"public_time":time.Now(),
		"article_classify":"go",
		"article_content":"content",
		"view_count":0,
	})
	this.Redirect("/",302)
}