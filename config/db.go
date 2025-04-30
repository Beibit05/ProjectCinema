package config

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

//var DB *sql.DB

var DB *gorm.DB

func InitDB() {
	///////////////////////////////GORM
	dns := "host=localhost user=postgres password=2005b dbname=cinema_db port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Filed to connect to database: ", err)
	}
	DB = db

	//err = db.AutoMigrate(&models.Film{}, &models.Genre{}, &models.Director{}, models.User{})
	//if err != nil {
	//	log.Fatal("Migration filed", err)
	//}

	fmt.Println("Database connected and migrated successfully ")
}
