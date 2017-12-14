package routers

import (
	"github.com/an-jun/wechat-demo/backend/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wx_connect", &controllers.WxConnectController{})
	beego.Router("/wx/createmenu", &controllers.WxConnectController{}, "*:Createmenu")
	beego.Router("/wx/demo", &controllers.WxController{}, "*:Demo")
	beego.Router("/test/test",&controllers.TestController{},"*:Test")
}
