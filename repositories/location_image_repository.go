package repositories

import (
	"location-api/models"

	"github.com/beego/beego/v2/client/orm"
)

type LocationImageRepository interface {
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
