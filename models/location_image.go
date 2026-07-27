package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type LocationImage struct {
	ID int `orm:"column(id);auto"`

	Location *Location `orm:"column(location_id);rel(fk)"`

	ImageKey string `orm:"column(image_key);size(255)"`

	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)"`
}

func (t *LocationImage) TableName() string {
	return "location_images"
}

func init() {
	orm.RegisterModel(new(LocationImage))
}
