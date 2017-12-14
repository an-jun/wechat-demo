package main

import (
	"github.com/astaxie/beego/context"

	_ "github.com/astaxie/beego/cache/redis"
	_ "github.com/an-jun/wechat-demo/backend/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	"github.com/an-jun/wechat-demo/backend/models"
	"fmt"
)

func main() {
	var err error
	models.CC, err = cache.NewCache("redis", `{"conn":"127.0.0.1:6379"}`)
	if err !=nil{
		fmt.Println("%v",err)
	}
	beego.Any("/test", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bob"))
	})
	beego.Run()
}
