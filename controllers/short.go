package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"beego_shorturl/cache"
	"beego_shorturl/models"
)

type ShortController struct {
	baseController
}

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

// @Title 生成长URL
// @Description 生成长URL或者读取缓存
// @Success 200 {object} models.response
// @router /api/short/:longurl [get]
func (this *ShortController) Get() {
	longurl := this.Input().Get("longurl")
	// 根据longurl生产key
	key:=models.GetMD5(longurl)
	// longurl-key : shorturl(生成的短地址)
	mc:=cache.GetMemoryCache()

	var short_res ShortResult
	short_res.UrlLong=longurl

	if mc.IsExist(key){
		short_res.UrlShort=mc.Get(key).(string)
		this.Data["json"]=this.baseController.BuildJSONResponse(0,short_res)
	} else{
		shortUrl:=models.GenerateShortUrl(longurl)
		if err:=mc.Put(key,shortUrl,time.Minute * 2);err != nil{
			beego.Info(err)
			this.Data["json"]=this.baseController.BuildJSONResponse(2001,nil)
		}
		if err:=mc.Put(shortUrl,longurl,time.Minute * 2);err != nil{
			beego.Info(err)
			this.Data["json"]=this.baseController.BuildJSONResponse(2001,nil)
		}
		short_res.UrlShort=shortUrl
		this.Data["json"]=this.baseController.BuildJSONResponse(0,short_res)
	}
	this.ServeJSON()
	return
}
