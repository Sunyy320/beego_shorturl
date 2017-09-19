package controllers

import (
	"github.com/astaxie/beego"
	"reflect"
)

type baseController struct {
	beego.Controller
}
// 返回数据格式化
type response struct {
	Code int `json:"code"`
	Msg string `json:"message"`
	Data interface{} `json:"data"`
}

func (this *baseController) BuildJSONResponse(code int,data interface{}) response{
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


