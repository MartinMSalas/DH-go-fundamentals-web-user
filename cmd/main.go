package main

import (
	"DH-go-fundamentals-web-user/internal/domain"
	"DH-go-fundamentals-web-user/internal/user"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	server := http.NewServeMux()

	db := user.DB{
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
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)
	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoints(ctx, service))

	fmt.Println("Server running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
