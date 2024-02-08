package bootstrap

import (
	"DH-go-fundamentals-web-user/internal/domain"
	"DH-go-fundamentals-web-user/internal/user"

	"log"

	"os"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func NewDB() user.DB {

	return user.DB{
		Users: []domain.User{{
			ID:        1,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "jhon@doe.com",
		}, {
			ID:        2,
			FirstName: "Jane",
			LastName:  "Doe",
			Email:     "jane@doe.com",
		}, {
			ID:        3,
			FirstName: "John",
			LastName:  "Smith",
			Email:     "jhon@smith.com",
		}},
		MaxUserD: 3,
	}
}
