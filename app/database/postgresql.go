package database

import (
	"fakhry/cleanarch/app/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBPostgres(cfg *config.AppConfig) *gorm.DB {
	// declare struct config & variable connectionString
	// usernama:password@tcp(hostdb:portdb)/db_name
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DB_HOSTNAME, cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT)
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	// db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
