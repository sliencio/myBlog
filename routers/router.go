package routers
import (
	"myBlog/controllers"
	"github.com/astaxie/beego"
)

func init() {


	beego.Router("/", &controllers.MainController{})
	beego.Router("/home", &controllers.MainController{})

	//文章主页和home一致
	beego.Router("/article",&controllers.ArticleController{})
	beego.Router("/article/edit",&controllers.ArticleController{},"get:Edit")
	beego.Router("/article/show/:id",&controllers.ArticleController{},"get:Show")
	beego.Router("/article/publish",&controllers.ArticleController{},"post:Publish")


    beego.Router("/login",&controllers.UserController{})
	beego.Router("/register",&controllers.UserController{},"get:Register")
}

