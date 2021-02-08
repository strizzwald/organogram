package models

import "github.com/google/uuid"

type Title struct {
	Href string    `json:"href"`
	Id   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
