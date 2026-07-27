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
		controller.sendBadRequestResponse("Invalid location id")
		return
	}

	ctx := context.Background()

	locationService, err := controller.newLocationService(ctx)
	if err != nil {
		controller.sendInternalServerErrorResponse()
		return
	}

	locationImages, err := locationService.GetLocationImages(ctx, locationID)
	if err != nil {
		if errors.Is(err, services.ErrLocationNotFound) {
			controller.sendNotFoundResponse("Location not found")
			return
		}

		controller.sendInternalServerErrorResponse()
		return
	}

	controller.sendSuccessResponse(
		http.StatusOK,
		"Location images retrieved successfully",
		locationImages,
	)
}

func (controller *LocationController) GetLocations() {
	ctx := context.Background()

	locationService, err := controller.newLocationService(ctx)
	if err != nil {
		controller.sendInternalServerErrorResponse()
		return
	}

	locations, err := locationService.GetLocations(ctx)
	if err != nil {
		controller.sendInternalServerErrorResponse()
		return
	}

	controller.sendSuccessResponse(
		http.StatusOK,
		"Locations retrieved successfully",
		locations,
	)
}

func (controller *LocationController) UpdateBaseImage() {
	locationID, err := strconv.Atoi(controller.Ctx.Input.Param(":id"))
	if err != nil {
		controller.sendBadRequestResponse("Invalid location id")
		return
	}

	var request dtos.UpdateBaseImageRequest

	if err := json.Unmarshal(controller.Ctx.Input.RequestBody, &request); err != nil {
		controller.sendBadRequestResponse("Invalid request body")
		return
	}

	ctx := context.Background()

	locationService, err := controller.newLocationService(ctx)
	if err != nil {
		controller.sendInternalServerErrorResponse()
		return
	}

	err = locationService.UpdateBaseImage(
		ctx,
		locationID,
		request.LocationImageID,
	)
	if err != nil {
		controller.sendInternalServerErrorResponse()
		return
	}

	controller.sendSuccessResponse(
		http.StatusOK,
		"Base image updated successfully",
		nil,
	)
}

func (controller *LocationController) newLocationService(ctx context.Context) (services.LocationService, error) {
	s3Client, err := utils.NewMinIOS3Client(ctx)
	if err != nil {
		return nil, err
	}

	locationRepository := repositories.NewLocationRepository()
	locationImageRepository := repositories.NewLocationImageRepository()
	storageRepository := repositories.NewStorageRepository(s3Client)

	return services.NewLocationService(
		locationRepository,
		locationImageRepository,
		storageRepository,
	), nil
}
