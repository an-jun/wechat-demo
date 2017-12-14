package controllers

import (
	"fmt"

	"github.com/an-jun/wechat-demo/backend/wxutils"
	"github.com/astaxie/beego"
)

type WxToolController struct {
	beego.Controller
}

func (c *WxConnectController) Createmenu() {
	// appId := "wx4c10dcbc112ceaeb"
	// appSecret := "6f802e7cef82c74596b9760e3b31f4ff"
	// accessToken, err := wxutils.FetchAccessToken(appId, appSecret, "https://api.weixin.qq.com/cgi-bin/token")
	// if err != nil {
	// 	fmt.Println("向微信服务器发送获取accessToken的get请求失败", err)
	// }
	accessToken := "4_wg2x9Pq8iR7pr0aOvV4vNJHZohDIxLoNN_E_Lz5GSGNsMs-Z2qcZrWs9L_sUamY7MHbF05EBz9RyEwO6wWmFnRwjdhbi8LeRaoXVlBRwiFEHKYso-svQkoRcK-nuEK6zM3h6HmPS8_6rVDnoQQQbAHAUUU"
	menuStr := `{
		"button": [
		{
			"name": "进入商城",
			"type": "view",
			"url": "http://wxaj.shdev.cpchina.cn/wx/1"
		},
		{

			"name":"管理中心",
			 "sub_button":[
					{
					"name": "用户中心",
					"type": "click",
					"key": "molan_user_center"
					},
					{
					"name": "公告",
					"type": "click",
					"key": "molan_institution"
					}]
		}
		]
	}`
	fmt.Println(menuStr)
	// wxutils.PushWxMenuCreate(accessToken, []byte(menuStr))
	wxutils.PushWxMenuDelete(accessToken)
	c.Ctx.WriteString(accessToken)
}
func (c *WxConnectController) demo() {

}
