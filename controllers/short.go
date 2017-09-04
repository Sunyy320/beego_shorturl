package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"beego_shorturl/cache"
	"beego_shorturl/models"
	"reflect"
)

type ShortController struct {
	beego.Controller
}

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

// 返回数据格式化
type response struct {
	Code int `json:"code"`
	Msg string `json:"message"`
	Data interface{} `json:"data"`
}

// @Title 生成长URL
// @Description 生成长URL或者读取缓存
// @Success 200 {object} models.response
// @router /api/short/:longurl [get]
func (this *ShortController) Get() {
	longurl := this.Input().Get("longurl")
	// 根据longurl生产key
	key:=models.GetMD5(longurl)
	mc:=cache.GetMemoryCache()

	var short_res ShortResult
	short_res.UrlLong=longurl

	if mc.IsExist(key){
		short_res.UrlShort=mc.Get(key).(string)
		this.Data["json"]=this.BuildJSONResponse(0,short_res)
	} else{
		shortUrl:=models.GenerateShortUrl(longurl)
		if err:=mc.Put(key,shortUrl,time.Minute * 2);err != nil{
			beego.Info(err)
			this.Data["json"]=this.BuildJSONResponse(2001,nil)
		}
		short_res.UrlShort=shortUrl
		this.Data["json"]=this.BuildJSONResponse(0,short_res)
	}
	this.ServeJSON()
	return
}


func (this *ShortController) BuildJSONResponse(code int,data interface{}) response{
	resp:=response{
		Code:code,
		Msg:"ok",
	}
	resp.Msg=getErrorCodeMsgMapping(code)

	t:=reflect.ValueOf(data)

	switch t.Kind() {
	case reflect.Array,reflect.Slice,reflect.Map:
		if t.Len() >0{
			resp.Data=data
		}
	case reflect.Struct:
		fallthrough
	case reflect.Ptr:
		resp.Data=data
	case reflect.Invalid:
	default:
		beego.Debug(this.Ctx.Request.URL,t.Kind())
	}
	if resp.Data==nil{
		resp.Data=map[string]string{}
	}
	return  resp
}

// 得到code对应的编码
func getErrorCodeMsgMapping(code int) string{
	mapping:=make(map[int] string)
	mapping[0] = "ok"
	mapping[1001] = "invalid params"
	mapping[1002] = "invalid identity"
	mapping[1003] = "method not allowed"
	mapping[1004] = "params db validation failed"
	mapping[1005] = "permission denied"

	mapping[2001] = "data create failed"
	mapping[2002] = "data not found"
	mapping[2003] = "data update failed"
	mapping[2004] = "data already exists"

	mapping[3001] = "no permission"
	mapping[3002] = "abnormal permission,try again later"

	// 查找键值是否存在
	if msg,ok:=mapping[code];ok{
		return  msg
	}
	return "Oops......."
}
