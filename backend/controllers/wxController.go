package controllers

import (
	"github.com/astaxie/beego"
	"github.com/an-jun/wechat-demo/backend/models"
	"fmt"
)

type WxController struct{
	beego.Controller
}
func(c * WxController)Demo(){

	url :="http://"+c.Ctx.Request.Host+c.Ctx.Request.RequestURI
	fmt.Println(url)
	jssdk := &models.Jssdk{AppId:beego.AppConfig.String("appId"),AppSecret:beego.AppConfig.String("appSecret"),Url:url}
	signPackage :=	jssdk.GetSignPackage()
	c.Data["signPackage"] = signPackage

}