package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/strizzwald/orgonogram/domain/entities"
	"github.com/strizzwald/orgonogram/models"
	"gorm.io/gorm"
)

type TitleController struct {
	db *gorm.DB
}

func NewTitlesController(db *gorm.DB) *TitleController {
	return &TitleController{db}
}

func (t *TitleController) GetTitles(w http.ResponseWriter, r *http.Request) {
	var entities []entities.Title

	t.db.Find(&entities)

	titles := []models.Title{}

	for _, e := range entities {
		title := models.Title{
			Id:   e.Guid,
			Href: fmt.Sprintf("http://localhost:8081/titles/%s", e.Guid.String()),
			Name: e.Name,
		}

		titles = append(titles, title)
	}

	json.NewEncoder(w).Encode(titles)
}
