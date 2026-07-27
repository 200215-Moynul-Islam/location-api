package services

import (
	"context"
	"errors"

	"location-api/dtos"
	"location-api/repositories"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/beego/beego/v2/client/orm"
)

var ErrLocationNotFound = errors.New("location not found")

type LocationService interface {
	GetBaseImage(ctx context.Context, id int) (*s3.GetObjectOutput, error)
	GetLocationImages(locationID int) ([]dtos.LocationImageResponse, error)
}

type locationService struct {
	locationRepository repositories.LocationRepository
	locationImageRepository repositories.LocationImageRepository
	storageRepository  repositories.StorageRepository
}

func NewLocationService(
	locationRepository repositories.LocationRepository,
	locationImageRepository repositories.LocationImageRepository,
	storageRepository repositories.StorageRepository,
) LocationService {
	return &locationService{
		locationRepository: locationRepository,
		locationImageRepository: locationImageRepository,
		storageRepository:  storageRepository,
	}
}

func (service *locationService) GetBaseImage(
	ctx context.Context,
	id int,
) (*s3.GetObjectOutput, error) {

	location, err := service.locationRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return service.storageRepository.GetObject(
		ctx,
		location.BaseImageKey,
	)
}

func (service *locationService) GetLocationImages(locationID int) ([]dtos.LocationImageResponse, error) {
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
		response = append(response, dtos.LocationImageResponse{
			ID:       locationImage.ID,
			ImageURL: locationImage.ImageKey,
		})
	}

	return response, nil
}
