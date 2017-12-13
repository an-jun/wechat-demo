package routers

import (
	"github.com/an-jun/wechat-demo/backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
