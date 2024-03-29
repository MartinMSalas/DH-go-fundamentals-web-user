package main

import (
	"DH-go-fundamentals-web-user/pkg/handler"
	"DH-go-fundamentals-web-user/internal/user"
	"DH-go-fundamentals-web-user/pkg/bootstrap"
	"context"
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	db := bootstrap.NewDB()
	logger := bootstrap.NewLogger()
	repo := user.NewRepo(db, logger)
	service := user.NewService(logger, repo)
	ctx := context.Background()

	handler.NewUserHTTPServer(ctx, server, user.MakeEndpoints(ctx, service))
	

	fmt.Println("Server running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
