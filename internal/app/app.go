package app

import (
	"fmt"

	"github.com/dexguitar/gotodoapi/config"
	"github.com/dexguitar/gotodoapi/internal/handler"
	"github.com/dexguitar/gotodoapi/internal/router"
	"github.com/dexguitar/gotodoapi/internal/service"
	"github.com/dexguitar/gotodoapi/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	Router *gin.Engine
}

func New(c *config.Config) (*App, error) {
	const op = "App.New"

	db, err := sqlx.Open("postgres", c.DBPrimary)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	storage := storage.NewTodoStorage(db)
	service := service.NewTodoService(storage)
	handler := handler.NewTodoHandler(service)
	router := router.Init(handler)

	return &App{
		Router: router,
	}, nil
}
