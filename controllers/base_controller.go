package controllers

import (
	"location-api/utils"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
}

func (controller *BaseController) sendSuccessResponse(statusCode int, message string, data any) {
	utils.SendJSONResponse(
		controller.Ctx,
		statusCode,
		true,
		message,
		data,
	)
}

func (controller *BaseController) sendErrorResponse(statusCode int, message string) {
	utils.SendJSONResponse(
		controller.Ctx,
		statusCode,
		false,
		message,
		nil,
	)
}

func (controller *BaseController) sendBadRequestResponse(message string) {
	controller.sendErrorResponse(http.StatusBadRequest, message)
}

func (controller *BaseController) sendNotFoundResponse(message string) {
	controller.sendErrorResponse(http.StatusNotFound, message)
}

func (controller *BaseController) sendInternalServerErrorResponse() {
	controller.sendErrorResponse(
		http.StatusInternalServerError,
		"Internal server error",
	)
}
