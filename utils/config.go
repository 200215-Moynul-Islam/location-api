package utils

import beego "github.com/beego/beego/v2/server/web"

func GetConfig(key string) string {
	value, err := beego.AppConfig.String(key)
	if err != nil {
		panic("missing configuration: " + key + ": " + err.Error())
	}

	return value
}
