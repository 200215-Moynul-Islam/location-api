package controllers

import (
	"context"
	"errors"
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
	locationImageRepository := repositories.NewLocationImageRepository()
	storageRepository := repositories.NewStorageRepository(s3Client)

	locationService := services.NewLocationService(
		locationRepository,
		locationImageRepository,
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

func (controller *LocationController) GetLocationImages() {
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
	locationImageRepository := repositories.NewLocationImageRepository()
	storageRepository := repositories.NewStorageRepository(s3Client)

	locationService := services.NewLocationService(
		locationRepository,
		locationImageRepository,
		storageRepository,
	)

	locationImages, err := locationService.GetLocationImages(locationID)
	if err != nil {
		if errors.Is(err, services.ErrLocationNotFound) {
			utils.SendJSONResponse(
				controller.Ctx,
				http.StatusNotFound,
				false,
				"Location not found",
				nil,
			)
			return
		}

		utils.SendJSONResponse(
			controller.Ctx,
			http.StatusInternalServerError,
			false,
			"Internal server error",
			nil,
		)
		return
	}

	utils.SendJSONResponse(
		controller.Ctx,
		http.StatusOK,
		true,
		"Location images retrieved successfully",
		locationImages,
	)
}

func (controller *LocationController) GetLocations() {
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
	locationImageRepository := repositories.NewLocationImageRepository()
	storageRepository := repositories.NewStorageRepository(s3Client)

	locationService := services.NewLocationService(
		locationRepository,
		locationImageRepository,
		storageRepository,
	)

	locations, err := locationService.GetLocations()
	if err != nil {
		utils.SendJSONResponse(
			controller.Ctx,
			http.StatusInternalServerError,
			false,
			"Internal server error",
			nil,
		)
		return
	}

	utils.SendJSONResponse(
		controller.Ctx,
		http.StatusOK,
		true,
		"Locations retrieved successfully",
		locations,
	)
}

func (controller *LocationController) GetLocationImage() {
	imageID, err := strconv.Atoi(controller.Ctx.Input.Param(":id"))
	if err != nil {
		utils.SendJSONResponse(
			controller.Ctx,
			http.StatusBadRequest,
			false,
			"Invalid image id",
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
	locationImageRepository := repositories.NewLocationImageRepository()
	storageRepository := repositories.NewStorageRepository(s3Client)

	locationService := services.NewLocationService(
		locationRepository,
		locationImageRepository,
		storageRepository,
	)

	object, err := locationService.GetLocationImage(ctx, imageID)
	if err != nil {
		if errors.Is(err, services.ErrLocationNotFound) {
			utils.SendJSONResponse(
				controller.Ctx,
				http.StatusNotFound,
				false,
				"Image not found",
				nil,
			)
			return
		}

		utils.SendJSONResponse(
			controller.Ctx,
			http.StatusInternalServerError,
			false,
			"Internal server error",
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
