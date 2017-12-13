package models

import (
	"github.com/astaxie/beego/orm"
)

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

func init() {
	orm.Debug = true
	orm.RegisterModel(new(WxBase))
}
