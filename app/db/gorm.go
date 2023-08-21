package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	// 開発環境の場合は環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatalln(err)
	}

	// DB接続
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbTask := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbTask, dbPassword, dbHost, dbPort, dbName)

	// gormで接続
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	fmt.Println("Successfully connected to database!")
	return db
}

func GormCloseDB(db *gorm.DB) {
	dbConn, _ := db.DB()
	err := dbConn.Close()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("DB切断成功")
}
