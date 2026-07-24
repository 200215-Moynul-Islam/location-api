package database

import (
	"fmt"

	"location-api/utils"

	"github.com/beego/beego/v2/client/orm"

	_ "github.com/lib/pq"
)

func InitDatabase() {
	registerDatabase()
	verifyConnection()
}

func registerDatabase() {
	if err := orm.RegisterDriver("postgres", orm.DRPostgres); err != nil {
		panic(fmt.Errorf("failed to register postgres driver: %w", err))
	}

	dataSource := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		utils.GetConfig("POSTGRES_USER"),
		utils.GetConfig("POSTGRES_PASSWORD"),
		utils.GetConfig("POSTGRES_DB"),
		utils.GetConfig("POSTGRES_HOST"),
		utils.GetConfig("POSTGRES_PORT"),
		utils.GetConfig("POSTGRES_SSLMODE"),
	)

	if err := orm.RegisterDataBase("default", "postgres", dataSource); err != nil {
		panic(fmt.Errorf("failed to register database: %w", err))
	}
}

func verifyConnection() {
	ormInstance := orm.NewOrm()

	var result int
	if err := ormInstance.Raw("SELECT 1").QueryRow(&result); err != nil {
		panic(fmt.Errorf("failed to connect to database: %w", err))
	}
}
