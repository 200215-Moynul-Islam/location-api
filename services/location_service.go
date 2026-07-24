package services

import (
	"context"

	"location-api/repositories"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type LocationService interface {
	GetBaseImage(ctx context.Context, id int) (*s3.GetObjectOutput, error)
}

type locationService struct {
	locationRepository repositories.LocationRepository
	storageRepository  repositories.StorageRepository
}

func NewLocationService(
	locationRepository repositories.LocationRepository,
	storageRepository repositories.StorageRepository,
) LocationService {
	return &locationService{
		locationRepository: locationRepository,
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
