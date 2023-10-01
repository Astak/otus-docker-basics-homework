package handler

import (
	"fmt"

	"github.com/Astak/otus-docker-basics-homework/web-service-gin/config"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/data"
)

type Handler struct {
	UserRepo data.UserRepository
}

func NewHandler(ur data.UserRepository) *Handler {
	return &Handler{UserRepo: ur}
}

func LoadHandler(configPath *string) *Handler {
	cfg, _ := config.LoadConfig(configPath)
	return LoadHandlerFromConfig(cfg)
}

func LoadHandlerFromConfig(cfg config.Config) *Handler {
	fmt.Printf("*** DB URL %s", cfg.DbUrl)
	database := data.NewDatabase(cfg)
	userRepo, _ := data.NewUserRepository(database)
	handler := NewHandler(userRepo)
	return handler
}
