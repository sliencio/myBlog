package main

import (
	_ "myBlog/routers"
	"github.com/astaxie/beego"
	_"myBlog/DB"
)

func main() {
	beego.Run()
}

