package controllers

import (
	"location-api/utils"
	"net/http"
)

type HealthController struct {
	BaseController
}

func (c *HealthController) Get() {
	utils.SendJSONResponse(c.Ctx, http.StatusOK, true, "Server is running.", nil)
}
