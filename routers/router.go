package routers

import (
	"gosex/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user", &controllers.UserController{})
	beego.Router("/api/:id[0-9]+", &controllers.ApiController{})
	beego.Router("/api/list", &controllers.ApiController{}, "*:ListFood")
}
