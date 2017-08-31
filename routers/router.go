package routers

import (
	"beego_shorturl/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/short",&controllers.ShortController{})
}
