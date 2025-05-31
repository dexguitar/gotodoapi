package router

import (
	"github.com/dexguitar/gotodoapi/internal/handler"
	"github.com/gin-gonic/gin"
)

func Init(h *handler.TodoHandler) *gin.Engine {
	r := gin.Default()

	r.POST("/todos", h.CreateTodo)
	r.GET("/todos", h.GetAllTodos)
	r.DELETE("/todos", h.DeleteTodoById)
	r.PUT("/done", h.CompleteTodo)

	return r
}
