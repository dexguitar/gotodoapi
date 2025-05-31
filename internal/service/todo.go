package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/dexguitar/gotodoapi/internal/errs"
	"github.com/dexguitar/gotodoapi/internal/model"
)

type TodoService struct {
	TodoStorage
}

func NewTodoService(storage TodoStorage) *TodoService {
	return &TodoService{
		TodoStorage: storage,
	}
}

func (s *TodoService) Create(ctx context.Context, newTodo *model.Todo) (int64, error) {
	const op = "TodoService.Create"

	if newTodo == nil {
		return 0, fmt.Errorf("%s: todo cannot be nil", op)
	}

	existing, err := s.TodoStorage.GetTodoByTitle(ctx, newTodo.Title)
	if existing != nil {
		return 0, fmt.Errorf("%s: %w", op, errs.ErrTodoExists)
	}
	if err != nil && !errors.Is(err, errs.ErrTodoNotFound) {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	id, err := s.TodoStorage.Create(ctx, newTodo)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *TodoService) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	const op = "TodoService.GetAllTodos"

	todos, err := s.TodoStorage.GetAllTodos(ctx)
	if err != nil {
		return []*model.Todo{}, fmt.Errorf("%s: %w", op, err)
	}

	return todos, nil
}

func (s *TodoService) DeleteTodoById(ctx context.Context, id int64) error {
	const op = "TodoService.DeleteTodoById"

	err := s.TodoStorage.DeleteTodoById(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *TodoService) CompleteTodo(ctx context.Context, id int64) error {
	const op = "TodoService.CompleteTodo"

	todo, err := s.TodoStorage.GetTodoById(ctx, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = s.TodoStorage.CompleteTodo(ctx, todo.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
