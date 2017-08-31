package controllers

import (
	"github.com/astaxie/beego"
	"time"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["json"]=map[string]interface{}{"code":0,"msg":time.Now(),"data":"this is msg"}
	this.ServeJSON()
}
