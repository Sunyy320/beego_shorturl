package controllers

import (
	"github.com/astaxie/beego"
	"time"
	"strings"
)

type ShortController struct {
	beego.Controller
}


func (this *ShortController) Get() {
	longurl := this.Input().Get("longurl")
	beego.Info("longurl=" + longurl)
	if len(strings.TrimSpace(longurl)) != 0 {
		this.Data["json"] = map[string]interface{}{"code": 0, "msg": time.Now(), "data": longurl}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "msg": time.Now(), "data": "longurl is empty"}
	}
	this.ServeJSON()
}
