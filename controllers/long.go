package controllers

import (
	"beego_shorturl/cache"
	"github.com/astaxie/beego"
)

type LongController struct {
	baseController
}

// 在short.go中已经声明，不需要再次声明
//type ShortResult struct {
//	UrlShort string
//	UrlLong  string
//}

func (this *LongController) Get () {
	shortUrl := this.Input().Get("shorturl")

	var short_res ShortResult
	short_res.UrlShort = shortUrl

	// 从缓存中读取相应的数据
	mc := cache.GetMemoryCache()
	if mc.IsExist(shortUrl) {
		beego.Info(shortUrl)
		short_res.UrlLong = mc.Get(shortUrl).(string)
	}
	this.Data["json"] = this.baseController.BuildJSONResponse(0, short_res)
	this.ServeJSON()
	return
}
