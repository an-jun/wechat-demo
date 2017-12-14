package models

import (
	"fmt"
	"time"
	"crypto/sha1"
	"io"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"github.com/astaxie/beego/cache"
)


func (this *Jssdk) GetSignPackage() SignPackage{
	jsapiTicket := this.GetJsApiTicket()
	fmt.Printf("jsapiTicket:[%v]",jsapiTicket)
	t := time.Now()
	timestamp := t.Unix()
	nonceStr := this.createNonceStr();

	// 这里参数的顺序要按照 key 值 ASCII 码升序排序
	str := "jsapi_ticket="+jsapiTicket+"&noncestr="+nonceStr+"&timestamp="+string(timestamp)+"&url="+this.Url;

	s := sha1.New()
	io.WriteString(s, str)

	sign := fmt.Sprintf("%x", s.Sum(nil))

	signPackage :=SignPackage{
		AppId:this.AppId,
		NonceStr:nonceStr,
		Timestamp:timestamp,
		Url:this.Url,
		Signature:sign,
		RawString:str,
	}
	return signPackage
}
func (this *Jssdk)GetJsApiTicket() string {
	fmt.Println("CC:%v",CC)
	jsapiTicket,ok :=  CC.Get("jsapiTicket").(JsapiTicket)
	t := time.Now()
	fmt.Printf("jsapiTicket,ok:[%v,%v]",jsapiTicket,ok)
	if !ok && jsapiTicket.expires_in < t.Unix(){
		accessToken := this.getAccessToken();
		// 如果是企业号用以下 URL 获取 ticket
		// $url = "https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket?access_token=$accessToken";
		url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token="+accessToken;
		resp,err := http.Get(url)
		if err !=nil{
			fmt.Println("getJsApiTicket err:%v",err)
		}
		body,err := ioutil.ReadAll(resp.Body)
		json.Unmarshal(body,&jsapiTicket)
		jsapiTicket.expires_in = t.Unix()+7000
		err = CC.Put("accessToken",accessToken,7200*time.Second)
		if err != nil {
			fmt.Println("getJsapiTicket err:%v",err)
		}
	}

	return jsapiTicket.ticket;
}
func (this *Jssdk)createNonceStr()string  {
	s := rand.Intn(100)
	return string(s)
}
func (this *Jssdk)getAccessToken()string  {

	accessToken,ok :=  CC.Get("accessToken").(AccessToken)
	fmt.Printf("accessToken:[%v]",accessToken)
	t := time.Now()
	if !ok && accessToken.Expires_in < t.Unix(){
		// 如果是企业号用以下URL获取access_token
		// $url = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=$this->appId&corpsecret=$this->appSecret";
		url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid="+this.AppId+"&secret="+this.AppSecret;
		fmt.Println("url:"+url)
		resp,err := http.Get(url)
		defer resp.Body.Close()
		if err != nil {
			fmt.Println("getAccessToken err:%v",err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("getAccessToken err:%v",err)
		}
		json.Unmarshal(body, &accessToken)
		accessToken.Expires_in = t.Unix()+7000
		fmt.Printf("put accessToken:%v",accessToken)
		err = CC.Put("accessToken",accessToken,7200*time.Second)
		if err != nil {
			fmt.Println("getAccessToken err:%v",err)
		}
	}
	return accessToken.Access_token
}
func init()  {
	if CC == nil{
		CC, _ = cache.NewCache("redis", `{"conn":"127.0.0.1:6379"}`)
	}
}