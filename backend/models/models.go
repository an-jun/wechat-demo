package models

import (
	"github.com/astaxie/beego/cache"
)
var CC cache.Cache
type WxAccessToken struct {
	Id int
}
type WxBase struct {
	Id        int
	AppID     string
	AppSecret string
}
type AccessTokenErrorResponse struct {
	Errcode float64
	Errmsg  string
}
type AccessToken struct {
	Access_token string `json:"access_token"`
	Expires_in int64 `json:"expires_in"`
}
type JsapiTicket struct {
	ticket string
	expires_in int64
}

type Jssdk struct {
	AppId     string
	AppSecret string
	Url       string
}
type SignPackage struct {
	AppId string
	NonceStr string
	Timestamp int64
	Url string
	Signature string
	RawString string
}
//func init() {
//	orm.Debug = true
//	orm.RegisterModel(new(WxBase))
//}
