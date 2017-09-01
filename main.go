package main

import (
	_ "beego_shorturl/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.DirectoryIndex=true
	beego.BConfig.WebConfig.StaticDir["/swagger"]="swagger"

	beego.Run()
}

