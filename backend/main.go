package main

import (
	"github.com/kataras/iris"
	"github.com/an-jun/wechat-demo/backend2/controllers"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"fmt"
	"github.com/an-jun/wechat-demo/backend2/models"
)

func main() {
	var err error
	models.CC, err = cache.NewCache("redis", `{"conn":"127.0.0.1:6379"}`)
	if err != nil {
		fmt.Println("%v", err)
	}
	app := iris.New()
	app.StaticWeb("/static", "./static")
	app.RegisterView(iris.HTML("./views", ".html").Reload(true))
	app.Controller("/helloworld", new(controllers.HelloWorldController))
	app.Controller("/wx", new(controllers.WxController))
	app.Run(iris.Addr("0.0.0.0:8080"), iris.WithConfiguration(iris.YAML("./configs/iris.yml")))

}