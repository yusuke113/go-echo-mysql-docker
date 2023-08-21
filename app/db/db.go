package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Initialize() *sql.DB {
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

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbTask, dbPassword, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	fmt.Println("Successfully connected to database!")
	return db
}

func CloseDB(db *sql.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("Error closing database: %v", err)
	}
	fmt.Println("DB切断成功")
}
