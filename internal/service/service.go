package service

import (
	"boilerplate/internal/repository"

	"github.com/gin-gonic/gin"
)

type Services struct {
	// list of service interfaces
}

func NewService(
	handler *gin.Engine,
	repo repository.Repository,
) Services {
	// init all services

	return Services{}
}
