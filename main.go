package main

import (
	"boilerplate/internal/handler"
	"boilerplate/internal/repository"
	"boilerplate/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	db := repository.ConnectDB()
	defer db.Close()

	repo := repository.NewRepo(db)

	router := gin.Default()

	services := service.NewService(router, repo)
	handler.SetRoute(services)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}
