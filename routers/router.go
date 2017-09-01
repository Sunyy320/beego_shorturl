// @APIVersion 1.0.0
// @Title ShortURL
// @Description short url
// @Contact sunyuanyuan@bilibili.com
package routers

import (
	"beego_shorturl/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	ns := beego.NewNamespace("/api",
		beego.NSCond(func(ctx *context.Context) bool {
			return true
		}),
		beego.NSRouter("/short",&controllers.ShortController{}),
	)
	beego.AddNamespace(ns)

}
