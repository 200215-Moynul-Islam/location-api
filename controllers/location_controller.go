package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"location-api/dtos"
	"location-api/repositories"
	"location-api/services"
	"location-api/utils"
)

type LocationController struct {
	BaseController
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

	locationImages, err := locationService.GetLocationImages(ctx, locationID)
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

	locations, err := locationService.GetLocations(ctx)
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

func (controller *LocationController) UpdateBaseImage() {
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

	var request dtos.UpdateBaseImageRequest

	if err := json.Unmarshal(controller.Ctx.Input.RequestBody, &request); err != nil {
		utils.SendJSONResponse(
			controller.Ctx,
			http.StatusBadRequest,
			false,
			"Invalid request body",
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

	err = locationService.UpdateBaseImage(
		ctx,
		locationID,
		request.LocationImageID,
	)
	if err != nil {
		utils.SendJSONResponse(
			controller.Ctx,
			http.StatusInternalServerError,
			false,
			"Failed to update base image",
			nil,
		)
		return
	}

	utils.SendJSONResponse(
		controller.Ctx,
		http.StatusOK,
		true,
		"Base image updated successfully",
		nil,
	)
}
