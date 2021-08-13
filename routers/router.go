package routers

import (
	"github.com/astaxie/beego"
	"learnbeego/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
