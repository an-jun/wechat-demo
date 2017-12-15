package utils

import (
	"crypto/sha1"
	"fmt"
	"time"
	"io"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"github.com/an-jun/wechat-demo/backend2/models"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"

	"strconv"
)
type AccessToken struct {
	Access_token string `json:"access_token"`
	Expires_in int64 `json:"expires_in"`
}
type JsapiTicket struct {
	Ticket string `json:"ticket"`
	Errcode int64 `json:"errcode"`
	Errmsg string `json:"errmsg"`
	Expires_in int64 `json:"expires_in"`
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
	Jsapi_ticket string
}
func (this *Jssdk) GetSignPackage() SignPackage{
	jsapiTicket := this.GetJsApiTicket()
	t := time.Now()
	timestamp := t.Unix()
	nonceStr := this.createNonceStr();

	// 这里参数的顺序要按照 key 值 ASCII 码升序排序
	str := "jsapi_ticket="+jsapiTicket+"&noncestr="+nonceStr+"&timestamp="+strconv.FormatInt(timestamp,10)+"&url="+this.Url;
	fmt.Printf("rawStr=======",str)
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
		Jsapi_ticket:jsapiTicket,
	}
	return signPackage
}
func (this *Jssdk)GetJsApiTicket() string {
	jsapiTicket_obj := models.CC.Get("jsapiTicket")

	var jsapiTicket JsapiTicket
	var err error
	if( jsapiTicket_obj !=nil){
		err = json.Unmarshal(jsapiTicket_obj.([]byte), &jsapiTicket)
		if err !=nil{
			fmt.Printf("%v",err)
		}
	}
	t := time.Now()
	fmt.Printf("jsapiTicket:[%v]",jsapiTicket_obj)
	if jsapiTicket_obj==nil || jsapiTicket.Expires_in < t.Unix(){
		accessToken := this.getAccessToken();
		// 如果是企业号用以下 URL 获取 ticket
		// $url = "https://qyapi.weixin.qq.com/cgi-bin/get_jsapi_ticket?access_token=$accessToken";
		url := "https://api.weixin.qq.com/cgi-bin/ticket/getticket?type=jsapi&access_token="+accessToken;
		resp,err := http.Get(url)
		if err !=nil{
			fmt.Println("getJsApiTicket err:%v",err)
		}
		body,err := ioutil.ReadAll(resp.Body)
		t2 :=string(body)
		fmt.Printf("getticketbody :%v",t2)

		json.Unmarshal(body,&jsapiTicket)
		jsapiTicket.Expires_in = t.Unix()+7000
		serialized, err := json.Marshal(jsapiTicket)
		err = models.CC.Put("jsapiTicket",serialized,7000*time.Second)
		if err != nil {
			fmt.Println("getJsapiTicket err:%v",err)
		}
	}

	return jsapiTicket.Ticket;
}
func (this *Jssdk)createNonceStr()string  {
	t:=time.Now()
	tt:=t.UnixNano()
	return strconv.FormatInt(tt,10)
}
func (this *Jssdk)getAccessToken()string  {
	accessToken_obj := models.CC.Get("accessToken")

	var accessToken AccessToken
	var err error
	if( accessToken_obj !=nil){
		err = json.Unmarshal(accessToken_obj.([]byte), &accessToken)
		if err !=nil{
			fmt.Printf("%v",err)
		}
	}

	fmt.Printf("accessToken:[%v]",accessToken)
	t := time.Now()
	if accessToken_obj==nil || accessToken.Expires_in < t.Unix(){
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
		serialized, err := json.Marshal(accessToken)
		err = models.CC.Put("accessToken",serialized,7000*time.Second)
		if err != nil {
			fmt.Println("getAccessToken err:%v",err)
		}
	}
	return accessToken.Access_token
}
func init()  {
	if models.CC ==nil{
		models.CC ,_= cache.NewCache("redis", `{"conn":"127.0.0.1:6379"}`)
	}
}