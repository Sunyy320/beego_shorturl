package cache

import (
	"github.com/astaxie/beego/cache"
)

var adapters = make(map[string]cache.Cache)

func GetMemoryCache() cache.Cache {
	bm , ok := adapters["memory"]
	if !ok {
		bm , _ := cache.NewCache("memory", `{"interval":60}`)
		adapters["memory"] = bm
		return adapters["memory"]
	}
	return bm
}
