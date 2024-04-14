package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/remusa/devtube/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"inserted_at"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
	}
}
