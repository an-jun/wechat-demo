package wxutils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/an-jun/wechat-demo/backend/models"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
)

type AccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
}

//type AccessTokenErrorResponse struct {
//  Errcode float64
//  Errmsg  string
//}
//获取wx_AccessToken 拼接get请求 解析返回json结果 返回 AccessToken和err
func FetchAccessToken(appID, appSecret, accessTokenFetchUrl string) (string, error) {

	requestLine := strings.Join([]string{accessTokenFetchUrl,
		"?grant_type=client_credential&appid=",
		appID,
		"&secret=",
		appSecret}, "")

	resp, err := http.Get(requestLine)
	if err != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("发送get请求获取 atoken 错误", err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("发送get请求获取 atoken 读取返回body错误", err)
		return "", err
	}

	if bytes.Contains(body, []byte("access_token")) {
		atr := AccessTokenResponse{}
		err = json.Unmarshal(body, &atr)
		if err != nil {
			fmt.Println("发送get请求获取 atoken 返回数据json解析错误", err)
			return "", err
		}
		return atr.AccessToken, err
	} else {
		fmt.Println("发送get请求获取 微信返回 err")
		ater := models.AccessTokenErrorResponse{}
		err = json.Unmarshal(body, &ater)
		fmt.Printf("发送get请求获取 微信返回 的错误信息 %+v\n", ater)
		if err != nil {
			return "", err
		}
		return "", fmt.Errorf("%s", ater.Errmsg)
	}
}

//type WxAccessToken struct {
//  Id          int `orm:"auto"`
//  AccessToken string
//}
//微信公众平台的参数
//type WxBase struct {
//  Id             int `orm:"auto"`
//  AppID          string
//  AppSecret      string
//  Token          string
//  EncodingAESKey string
//}
//数据库存取更新access_token
func GetAndUpdateDBWxAToken(o orm.Ormer) error {

	at := models.WxAccessToken{Id: 1}
	o.ReadOrCreate(&at, "id")

	wxBase := models.WxBase{Id: 1}
	err := orm.NewOrm().Read(&wxBase)
	if err != nil {
		fmt.Println("从数据库查询WxBase失败", err)
		return err
	}

	//向微信服务器发送获取accessToken的get请求
	accessToken, err := FetchAccessToken(wxBase.AppID, wxBase.AppSecret, "https://api.weixin.qq.com/cgi-bin/token")
	if err != nil {
		fmt.Println("向微信服务器发送获取accessToken的get请求失败", err)
		return err
	}
	fmt.Println(accessToken)
	o.Update(&at, "access_token")

	return nil
}
func startTimer(f func()) {
	go func() {
		for {
			f()
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
			t := time.NewTimer(next.Sub(now))
			<-t.C
		}
	}()
}

//从task_test.go文件里找使用方法
func shopTimeTask(o orm.Ormer) {

	timeStr2 := "0 */60 * * * *" // 获取wx token
	t2 := toolbox.NewTask("getAtoken", timeStr2, func() error {

		err := GetAndUpdateDBWxAToken(o)
		if err != nil {
			//todo 向微信请求access_token失败 结合业务逻辑处理
		}
		return nil
	})
	toolbox.AddTask("tk2", t2)
	toolbox.StartTask()

}
