package models

import "time"

type User struct {
	Id         int `json:"id"`
	Password   string
	CellNumber string    `json:"cellNumber"`
	Email      string    `json:"email"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
