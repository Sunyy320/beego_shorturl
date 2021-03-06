// @APIVersion 1.0.0
// @Title ShortURL
// @Description short url
// @Contact 1062666905@qq.com
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
		beego.NSRouter("/long",&controllers.LongController{}),
	)
	beego.AddNamespace(ns)

}
