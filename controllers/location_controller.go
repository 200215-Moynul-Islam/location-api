package controllers

import (
	"context"
	"io"
	"net/http"
	"strconv"

	"location-api/repositories"
	"location-api/services"
	"location-api/utils"
)

type LocationController struct {
	BaseController
}

func (controller *LocationController) GetBaseImage() {
	locationID, err := strconv.Atoi(controller.Ctx.Input.Param(":id"))
	if err != nil {
		utils.SendJSONResponse(
			controller.Ctx,
			http.StatusBadRequest,
			false,
			"Invalid location id",
			nil,
		)
		return
	}

	ctx := context.Background()

	s3Client, err := utils.NewMinIOS3Client(ctx)
	if err != nil {
		utils.SendJSONResponse(
			controller.Ctx,
			http.StatusInternalServerError,
			false,
			"Failed to initialize MinIO client",
			nil,
		)
		return
	}

	locationRepository := repositories.NewLocationRepository()
	storageRepository := repositories.NewStorageRepository(s3Client)

	locationService := services.NewLocationService(
		locationRepository,
		storageRepository,
	)

	object, err := locationService.GetBaseImage(ctx, locationID)
	if err != nil {
		utils.SendJSONResponse(
			controller.Ctx,
			http.StatusNotFound,
			false,
			"Base image not found",
			nil,
		)
		return
	}
	defer object.Body.Close()

	contentType := "application/octet-stream"
	if object.ContentType != nil {
		contentType = *object.ContentType
	}

	controller.Ctx.Output.Header("Content-Type", contentType)
	controller.Ctx.Output.SetStatus(http.StatusOK)

	if _, err := io.Copy(controller.Ctx.ResponseWriter, object.Body); err != nil {
		utils.SendJSONResponse(
			controller.Ctx,
			http.StatusInternalServerError,
			false,
			"Failed to stream image",
			nil,
		)
		return
	}
}
