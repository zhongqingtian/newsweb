package routers

import (
	"news/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/register", &controllers.UserController{},"get:ShowReg;post:HandleRge")

}
