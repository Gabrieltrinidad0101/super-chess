package database

import (
	"backend/src/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	conf := utils.Enviroments()
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=5432 sslmode=disable TimeZone=America/Los_Angeles",
		conf.DbHost,
		conf.DbUser,
		conf.DbPassword,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Errorf("Error %s", err.Error())
		panic("ERROR IN THE CONNECTION")
	}
	db.Exec("CREATE DATABASE IF NOT EXISTS " + conf.DbName)

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()

	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=America/Los_Angeles",
		conf.DbHost,
		conf.DbUser,
		conf.DbPassword,
		conf.DbName,
		conf.DbPort,
	)

	gormDB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Errorf("Error %s", err.Error())
		panic("ERROR IN THE CONNECTION")
	}
	return gormDB
}
