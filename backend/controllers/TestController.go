package controllers

import (
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"

	"fmt"
	"encoding/json"
	"github.com/an-jun/wechat-demo/backend/models"
)

type TestController struct {
	beego.Controller
}

var bm   cache.Cache
func (c* TestController)Test()  {
	s2 :=`{
"access_token": "4_9yMP8ZOktkobSxVDgiAHTdiQWeCbvycH3YnjRQ87shCiUqPEMuKh6RpB7dmiLPnFVOkPazoDNLQUn2-TvFeZol3ettqIZ62IFmwXlekKUcNjbE66PJxes7IKtEiQhZ_nlsqeMOw0yO5Tk5WBEQUgABAUGK",
"expires_in": 7200
}`
//var aMap map[string]interface{}
var aMap models.AccessToken
	json.Unmarshal([]byte(s2),&aMap)
	ss:=fmt.Sprintf("%v",aMap)
	c.Ctx.WriteString(ss)
	//var cf map[string]string
	//s1 := `{"conn":"127.0.0.1:6379"}`
	//json.Unmarshal([]byte(s1), &cf)
	//fmt.Println("%v",cf)
	//
	//
	//fmt.Printf("bm:%v",bm)
	//bm.Put("astaxie", "afaf", 10*time.Second)
	////bm.IsExist("astaxie")
	//s :=	bm.Get("astaxie").([]byte)
	////bm.Delete("astaxie")
	//c.Ctx.WriteString("test"+string(s))
}
//func init()  {
//	var err error
//	bm, err = cache.NewCache("redis", `{"conn":"127.0.0.1:6379"}`)
//	if err !=nil{
//		fmt.Printf("bm err:%v",err)
//	}
//}