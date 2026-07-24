package routers

import (
	"location-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/health", &controllers.HealthController{})

	beego.Router(
		"/locations/:id/base-image",
		&controllers.LocationController{},
		"get:GetBaseImage",
	)
}
