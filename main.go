package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/strizzwald/orgonogram/controllers"
	"github.com/strizzwald/orgonogram/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Articles []Article

var titlesController *controllers.TitlesController

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page!")
	fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)

	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/articles", returnAllArticles)

	myRouter.HandleFunc("/articles/{id}", returnSingleArticle)

	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")

	myRouter.HandleFunc("/titles", titlesController.GetTitles)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	dsn := "root:Passw0rd@tcp(mysql:3306)/organogram"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	db.AutoMigrate(&models.Title{})

	uuid, _ := uuid.Parse("139f3d4b-32dc-4928-b281-8231248dee97")
	db.Create(&models.Title{Href: "http://localhost:10000/titles", Id: uuid, Name: "Mr."})

	fmt.Printf("Title: %v", db.Take(&models.Title{}))

	titlesController = new(controllers.TitlesController)

	handleRequests()
}

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
