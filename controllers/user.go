package controllers

import "github.com/astaxie/beego"

type UserController struct{
	beego.Controller
}

func(this *UserController)Get(){
	this.TplName = "login.html"
}

func(this *UserController)Post(){
	this.Redirect("/home",301)
}

func(this *UserController)Register(){
	this.TplName = "register.html"
}