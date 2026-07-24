package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Location struct {
	ID           int       `orm:"auto"`
	Country      string    `orm:"size(100)"`
	State        string    `orm:"size(100)"`
	City         string    `orm:"size(100)"`
	Slug         string    `orm:"size(255);unique"`
	BaseImageKey string    `orm:"size(255)"`
	CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt    time.Time `orm:"auto_now;type(datetime)"`
}

// Set the table name
func (t *Location) TableName() string {
	return "locations"
}

func init() {
	orm.RegisterModel(new(Location))
}
