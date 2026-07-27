package services

import (
	"context"
	"errors"

	"location-api/dtos"
	"location-api/repositories"

	"github.com/beego/beego/v2/client/orm"
)

var ErrLocationNotFound = errors.New("location not found")

type LocationService interface {
	GetLocationImages(ctx context.Context, locationID int) ([]dtos.LocationImageResponse, error)
	GetLocations(ctx context.Context) ([]dtos.LocationResponse, error)
	UpdateBaseImage(ctx context.Context, locationID int, locationImageID int) error
}

type locationService struct {
	locationRepository      repositories.LocationRepository
	locationImageRepository repositories.LocationImageRepository
	storageRepository       repositories.StorageRepository
}

func NewLocationService(
	locationRepository repositories.LocationRepository,
	locationImageRepository repositories.LocationImageRepository,
	storageRepository repositories.StorageRepository,
) LocationService {
	return &locationService{
		locationRepository:      locationRepository,
		locationImageRepository: locationImageRepository,
		storageRepository:       storageRepository,
	}
}

func (service *locationService) GetLocationImages(ctx context.Context, locationID int) ([]dtos.LocationImageResponse, error) {
	_, err := service.locationRepository.GetByID(locationID)
	if err != nil {
		if errors.Is(err, orm.ErrNoRows) {
			return nil, ErrLocationNotFound
		}

		return nil, err
	}

	locationImages, err := service.locationImageRepository.GetByLocationID(locationID)
	if err != nil {
		return nil, err
	}

	response := make([]dtos.LocationImageResponse, 0, len(locationImages))

	for _, locationImage := range locationImages {
		imageURL, err := service.storageRepository.GeneratePresignedGetObjectURL(ctx, locationImage.ImageKey)
		if err != nil {
			return nil, err
		}

		response = append(response, dtos.LocationImageResponse{
			ID:       locationImage.ID,
			ImageURL: imageURL,
		})
	}

	return response, nil
}

func (service *locationService) GetLocations(ctx context.Context) ([]dtos.LocationResponse, error) {
	locations, err := service.locationRepository.GetAll()
	if err != nil {
		return nil, err
	}

	responses := make([]dtos.LocationResponse, 0)

	for _, location := range locations {
		baseImageURL, err := service.storageRepository.GeneratePresignedGetObjectURL(ctx, location.BaseImageKey)
		if err != nil {
			return nil, err
		}

		responses = append(responses, dtos.LocationResponse{
			ID:           location.ID,
			Country:      location.Country,
			State:        location.State,
			City:         location.City,
			BaseImageURL: baseImageURL,
		})
	}

	return responses, nil
}

func (service *locationService) UpdateBaseImage(
	ctx context.Context,
	locationID int,
	locationImageID int,
) error {
	location, err := service.locationRepository.GetByID(locationID)
	if err != nil {
		return err
	}

	locationImage, err := service.locationImageRepository.GetByID(locationImageID)
	if err != nil {
		return err
	}

	if locationImage.Location.ID != location.ID {
		return errors.New("image does not belong to location")
	}

	return service.storageRepository.CopyObject(
		ctx,
		locationImage.ImageKey,
		location.BaseImageKey,
	)
}
