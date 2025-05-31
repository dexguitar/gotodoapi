package handler

import (
	"context"

	"github.com/dexguitar/gotodoapi/internal/model"
)

type TodoService interface {
	Create(ctx context.Context, todo *model.Todo) (int64, error)
	GetAllTodos(ctx context.Context) ([]*model.Todo, error)
	DeleteTodoById(ctx context.Context, id int64) error
	CompleteTodo(ctx context.Context, id int64) error
}
