package entities

import "time"

type Users []User

/*
This struct must be same as /models/users.go type User
*/
type User struct {
	Id        int       `json:"id"`
	Username  string    `json:"username" `
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname" `
	Email     string    `json:"email"`
	Contact   string    `json:"contact"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// For Future
// Desc         string    `json:"desc"`
// Tags         []string  `json:"tags"`
// Progress     int       `json:"progress"`
// Organization bool      `json:"organization"`
