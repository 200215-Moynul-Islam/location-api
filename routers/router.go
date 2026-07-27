package routers

import (
	"location-api/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	api := beego.NewNamespace("/api/v1",
		beego.NSRouter(
			"/health",
			&controllers.HealthController{},
		),
		beego.NSRouter(
			"/locations/:id/base-image",
			&controllers.LocationController{},
			"get:GetBaseImage",
		),
		beego.NSRouter(
			"/locations/:id/images",
			&controllers.LocationController{},
			"get:GetLocationImages",
		),
	)

	beego.AddNamespace(api)
}
