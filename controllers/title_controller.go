package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/strizzwald/orgonogram/models"
)

type TitlesController struct{}

func (t *TitlesController) GetTitles(w http.ResponseWriter, r *http.Request) {

	mrHref := "http://localhost:10000/titles"
	mrGuid, _ := uuid.Parse("ecab2406-cd8c-4aaf-ad7b-c940bed939ef")
	mrHref = strings.Join([]string{mrHref, mrGuid.String()}, "/")

	mrTitle := models.Title{Href: mrHref, Id: mrGuid, Name: "Mr."}

	missHref := "http://localhost:10000/titles"
	missGuid, _ := uuid.Parse("139f3d4b-32dc-4928-b281-8231248dee97")
	missHref = strings.Join([]string{missHref, missGuid.String()}, "/")

	missTitle := models.Title{Href: missHref, Id: missGuid, Name: "Miss"}

	msHref := "http://localhost:10000/titles"
	msGuid, _ := uuid.Parse("437d0f67-1626-4b9e-bb8d-1a73a3932994")
	msHref = strings.Join([]string{msHref, msGuid.String()}, "/")

	msTitle := models.Title{Href: missHref, Id: missGuid, Name: "Ms"}

	mrsHref := "http://localhost:10000/titles"
	mrsGuid, _ := uuid.Parse("437d0f67-1626-4b9e-bb8d-1a73a3932994")
	mrsHref = strings.Join([]string{mrsHref, msGuid.String()}, "/")

	mrsTitle := models.Title{Href: mrsHref, Id: mrsGuid, Name: "Mrs"}

	mxHref := "http://localhost:10000/titles"
	mxGuid, _ := uuid.Parse("28d62566-8964-4841-a500-22f4b56eee09")
	mxHref = strings.Join([]string{mxHref, mxGuid.String()}, "/")

	mxTitle := models.Title{Href: mxHref, Id: mxGuid, Name: "Mx"}

	drHref := "http://localhost:10000/titles"
	drGuid, _ := uuid.Parse("0d3b4189-d647-435d-a3e2-3a13d7dad86e")
	drHref = strings.Join([]string{drHref, drGuid.String()}, "/")

	drTitle := models.Title{Href: mxHref, Id: mxGuid, Name: "Dr."}

	profHref := "http://localhost:10000/titles"
	profGuid, _ := uuid.Parse("13ea729a-ebbe-44a8-b0f2-23032eb8c0ef")
	profHref = strings.Join([]string{profHref, drGuid.String()}, "/")

	profTitle := models.Title{Href: profHref, Id: profGuid, Name: "Prof."}

	json.NewEncoder(w).Encode([]models.Title{mrTitle, missTitle, mrsTitle, msTitle, mxTitle, drTitle, profTitle})
}
