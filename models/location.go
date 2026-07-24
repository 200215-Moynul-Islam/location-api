package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Location struct {
	ID int `orm:"column(id);auto"`

	Country string `orm:"column(country);size(100)"`
	State   string `orm:"column(state);size(100)"`
	City    string `orm:"column(city);size(100)"`

	Slug         string `orm:"column(slug);size(255);unique"`
	BaseImageKey string `orm:"column(base_image_key);size(255)"`

	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

// Set the table name
func (t *Location) TableName() string {
	return "locations"
}

func init() {
	orm.RegisterModel(new(Location))
}
