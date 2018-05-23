package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type EditorController struct{
	beego.Controller
}

func (this* EditorController) Get(){
	this.TplName = "editor.html"
}
func (this* EditorController) Post(){
	content:=this.GetString("content")
	fmt.Println(content)
	this.Ctx.WriteString("ok")
}

