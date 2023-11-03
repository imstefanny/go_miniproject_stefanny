package database

import (
	"fmt"
	"miniproject/config"
	"miniproject/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitialMigration()
}

func InitDB() *gorm.DB{
	dbconfig := config.ReadEnv()
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbconfig.DB_USERNAME,
		dbconfig.DB_PASSWORD,
		dbconfig.DB_HOSTNAME,
		dbconfig.DB_PORT,
		dbconfig.DB_NAME,
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
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Account{})
	DB.AutoMigrate(&model.Studio{})
	DB.AutoMigrate(&model.Show{})
	DB.AutoMigrate(&model.Seat{})
	DB.AutoMigrate(&model.Ticket{})
	DB.AutoMigrate(&model.Transaction{})
}
