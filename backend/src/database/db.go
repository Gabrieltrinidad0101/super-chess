package database

import (
	"backend/src/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection() *gorm.DB {
	conf := utils.Enviroments()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DbUser,
		conf.DbPassword,
		conf.DbHost,
		conf.DbPort,
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

	dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.DbUser,
		conf.DbPassword,
		conf.DbHost,
		conf.DbPort,
		conf.DbName,
	)

	gormDB, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Errorf("Error %s", err.Error())
		panic("ERROR IN THE CONNECTION")
	}
	return gormDB
}
