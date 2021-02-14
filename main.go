package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/strizzwald/orgonogram/controllers"
	"github.com/strizzwald/orgonogram/domain/entities"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Articles []Article

var titlesController *controllers.TitleController

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/titles", titlesController.GetTitles)

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	var db *gorm.DB
	var err error

	dsn := "organogramapi:P@ssw0rd@tcp(mysql:3306)/organogram?charset=utf8mb4&parseTime=True&loc=Local"
	maxConnectionRetryCount := 5

	for i := 0; i < maxConnectionRetryCount; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil && i == maxConnectionRetryCount {
			panic(fmt.Sprintf("failed to connect to database: %v", err))
		} else if err != nil && i != maxConnectionRetryCount {
			sleepTime := math.Pow(2, float64(i))
			time.Sleep(time.Duration(sleepTime) * time.Second)
			continue
		} else {
			break
		}
	}

	migrateDatabase(db)
	seedDatabase(db)
	registerEndpoints(db)
}

func registerEndpoints(db *gorm.DB) {
	titlesController = controllers.NewTitlesController(db)
	handleRequests()
}

func migrateDatabase(db *gorm.DB) {
	var title entities.Title

	if !db.Migrator().HasTable(&title) {
		db.Migrator().CreateTable(&title)
	}
}

func seedDatabase(db *gorm.DB) {
	seedTitles(db)
}

func seedTitles(db *gorm.DB) {

	var titles []entities.Title
	db.Find(&titles)

	if len(titles) == 0 {
		id, _ := uuid.Parse("ecab2406-cd8c-4aaf-ad7b-c940bed939ef")
		db.Create(&entities.Title{Guid: id, Name: "Mr."})

		id, _ = uuid.Parse("139f3d4b-32dc-4928-b281-8231248dee97")
		db.Create(&entities.Title{Guid: id, Name: "Miss"})

		id, _ = uuid.Parse("437d0f67-1626-4b9e-bb8d-1a73a3932994")
		db.Create(&entities.Title{Guid: id, Name: "Ms"})

		id, _ = uuid.Parse("50ba0dea-6c56-11eb-8a56-0242ac130002")
		db.Create(&entities.Title{Guid: id, Name: "Mrs"})

		id, _ = uuid.Parse("28d62566-8964-4841-a500-22f4b56eee09")
		db.Create(&entities.Title{Guid: id, Name: "Mx"})

		id, _ = uuid.Parse("0d3b4189-d647-435d-a3e2-3a13d7dad86e")
		db.Create(&entities.Title{Guid: id, Name: "Dr."})
	}

}

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
