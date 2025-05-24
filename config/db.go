package config

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

// var DB *sql.DB
var DB *gorm.DB

//	func InitDB() {
//		///////////////////////////////GORM
//		dns := "host=db user=postgres password=2005b dbname=cinema_db port=5432 sslmode=disable TimeZone=Asia/Almaty"
//		db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
//		if err != nil {
//			log.Fatal("Error Filed to connect to database: ", err)
//		}
//		DB = db
//
//		//err = db.AutoMigrate(&models.Film{}, &models.Genre{}, &models.Director{}, models.User{})
//		//if err != nil {
//		//	log.Fatal("Migration filed", err)
//		//}
//
//		fmt.Println("Database connected and migrated successfully ")
//	}
func InitDB() {
	dns := "host=db user=postgres password=2005b dbname=cinema_db port=5432 sslmode=disable TimeZone=Asia/Almaty"

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err == nil {
			DB = db
			fmt.Println("Database connected successfully")
			return
		}

		log.Printf("Database connection failed: %v. Retrying...", err)
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Error: Failed to connect to database after retries: ", err)
}
