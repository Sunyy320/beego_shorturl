package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"beego_shorturl/cache"
	"beego_shorturl/models"
)

type ShortController struct {
	beego.Controller
}

type ShortResult struct {
	UrlShort string
	UrlLong  string
}

func (this *ShortController) Get() {
	longurl := this.Input().Get("longurl")
	// 根据longurl生产key
	key:=models.GetMD5(longurl)
	mc:=cache.GetMemoryCache()

	var short_res ShortResult
	short_res.UrlLong=longurl

	if mc.IsExist(key){
		short_res.UrlShort=mc.Get(key).(string)
		this.Data["json"] = map[string]interface{}{"code": 0, "msg": time.Now(), "data": short_res}
	} else{
		shortUrl:=models.GenerateShortUrl(longurl)
		if err:=mc.Put(key,shortUrl,time.Minute * 2);err != nil{
			beego.Info(err)
			this.Data["json"] = map[string]interface{}{"code": 0, "msg": err, "data": ""}
		}
		short_res.UrlShort=shortUrl
		this.Data["json"] = map[string]interface{}{"code": 0, "msg": time.Now(), "data":short_res}
	}
	this.ServeJSON()
	return
}
