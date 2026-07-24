package repositories

import (
	"location-api/models"

	"github.com/beego/beego/v2/client/orm"
)

type LocationRepository interface {
	GetAll() ([]models.Location, error)
	GetByID(id int) (*models.Location, error)
}

type locationRepository struct {
	ormInstance orm.Ormer
}

func NewLocationRepository() LocationRepository {
	return &locationRepository{
		ormInstance: orm.NewOrm(),
	}
}

func (repository *locationRepository) GetAll() ([]models.Location, error) {
	var locations []models.Location

	_, err := repository.ormInstance.QueryTable(new(models.Location)).All(&locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}

func (repository *locationRepository) GetByID(id int) (*models.Location, error) {
	location := &models.Location{
		ID: id,
	}

	err := repository.ormInstance.Read(location)
	if err != nil {
		return nil, err
	}

	return location, nil
}
