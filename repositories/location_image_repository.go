package repositories

import (
	"location-api/models"

	"github.com/beego/beego/v2/client/orm"
)

type LocationImageRepository interface {
	GetByID(id int) (*models.LocationImage, error)
	GetByLocationID(locationID int) ([]models.LocationImage, error)
}

type locationImageRepository struct {
	ormInstance orm.Ormer
}

func NewLocationImageRepository() LocationImageRepository {
	return &locationImageRepository{
		ormInstance: orm.NewOrm(),
	}
}

func (repository *locationImageRepository) GetByLocationID(locationID int) ([]models.LocationImage, error) {
	var images []models.LocationImage

	_, err := repository.ormInstance.
		QueryTable(new(models.LocationImage)).
		Filter("location_id", locationID).
		OrderBy("id").
		All(&images)

	if err != nil {
		return nil, err
	}

	return images, nil
}

func (repository *locationImageRepository) GetByID(id int) (*models.LocationImage, error) {
	locationImage := &models.LocationImage{
		ID: id,
	}

	err := repository.ormInstance.Read(locationImage)
	if err != nil {
		return nil, err
	}

	return locationImage, nil
}
