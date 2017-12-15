package controllers

import (
	"github.com/kataras/iris/mvc"
	"fmt"
	"github.com/an-jun/wechat-demo/backend2/utils"
)

type WxController struct{
	mvc.C
}
func (c *WxController) Get() string {

	return "This is my default action..."
}

//
// GET: /helloworld/{name:string}

func (c *WxController) GetBy(name string) string {
	return "Hello " + name
}

//
// GET: /helloworld/welcome

func (c *WxController) GetDemo() {
	url :="http://"+c.Ctx.Request().Host+c.Ctx.Request().RequestURI
	fmt.Println(url)
	jssdk := &utils.Jssdk{AppId:"wx4c10dcbc112ceaeb",AppSecret:"6f802e7cef82c74596b9760e3b31f4ff",Url:url}
	signPackage :=	jssdk.GetSignPackage()
	c.Ctx.ViewData("signPackage",signPackage)
	c.Ctx.View("index.html")
}

//
// GET: /helloworld/welcome/{name:string}/{numTimes:int}

func (c *WxController) GetWelcomeBy(name string, numTimes int) {
	// Access to the low-level Context,
	// output arguments are optional of course so we don't have to use them here.
	c.Ctx.Writef("Hello %s, NumTimes is: %d", name, numTimes)
}
