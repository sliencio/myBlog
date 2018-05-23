package routers

import (
	"myBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.MainController{})
    beego.Router("/editor",&controllers.EditorController{})
    beego.Router("/showArticle",&controllers.ArticleController{})
    beego.Router("/login",&controllers.UserController{})
	beego.Router("/register",&controllers.UserController{},"get:Register")
}

