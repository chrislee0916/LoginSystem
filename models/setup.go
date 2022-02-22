package models

import (
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Setup() {

	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("gorm.Open error: ", err.Error())
	}

	//建立三個table
	DB.AutoMigrate(&User{}, &Category{}, &Product{})

	c1 := Category{CName: "food"}
	c2 := Category{CName: "clothing"}
	c3 := Category{CName: "furniture"}
	categories := []Category{c1, c2, c3}

	//在categories table裡增加分類的資料
	DB.Create(&categories)

}
