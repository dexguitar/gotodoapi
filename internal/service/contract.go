package service

import (
	"context"

	"github.com/dexguitar/gotodoapi/internal/model"
)

type TodoStorage interface {
	Create(ctx context.Context, todo *model.Todo) (int64, error)
	GetTodoByTitle(ctx context.Context, title string) (*model.Todo, error)
	GetTodoById(ctx context.Context, id int64) (*model.Todo, error)
	GetAllTodos(ctx context.Context) ([]*model.Todo, error)
	DeleteTodoById(ctx context.Context, id int64) error
	CompleteTodo(ctx context.Context, id int64) error
}
