package model

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Config struct {
	User   string
	Pass   string
	Host   string
	Port   string
	DBName string
}

func ConnectDB() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error read .env file")
	}

	config := &Config{
		User:   os.Getenv("USER"),
		Pass:   os.Getenv("PASS"),
		Host:   os.Getenv("HOST"),
		Port:   os.Getenv("PORT"),
		DBName: os.Getenv("DB_NAME"),
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Pass, config.Host, config.Port, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("err")
	}

	db.AutoMigrate(&Product{}, &User{})

	DB = db
}
