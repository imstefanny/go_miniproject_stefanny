package config

import (
	"fmt"
	"miniproject/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
}

type Config struct {
	DB_Username string
	DB_Password string
	DB_Port     string
	DB_Host     string
	DB_Name     string
}

func InitDB() *gorm.DB{
	config := Config{
		DB_Username: "root",
		DB_Password: "Sw221103",
		DB_Port:     "3306",
		DB_Host:     "localhost",
		DB_Name:     "miniproject",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return DB
}

func InitialMigration() {
	DB.AutoMigrate(&model.Cinema{})
	DB.AutoMigrate(&model.Movie{})
}

