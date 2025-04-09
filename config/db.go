package config

import (
	"ProjectCinema/models"
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
	dns := "host=localhost user=postgres password=2005b dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Almaty"
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal("Error Filed to connect to database: ", err)
	}
	DB = db

	err = db.AutoMigrate(&models.Film{})
	if err != nil {
		log.Fatal("Migration filed", err)
	}

	fmt.Println("Database connected and migrated successfully ")
	/////////////////////////////// postgres!!!
	//connStr := "host=localhost port=5432 user=postgres password=2005b dbname=postgres sslmode=disable"
	//var err error
	//DB, err := sql.Open("postgres", connStr)
	//if err != nil {
	//	log.Fatalf("Error connecting to DB: %s", err)
	//}
	//
	//if err = DB.Ping(); err != nil {
	//	log.Fatalf("DB ping error: %s", err)
	//}
	//
	//fmt.Println("DB successfully connected to postgres")

}
