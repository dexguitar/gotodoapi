package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/dexguitar/gotodoapi/internal/errs"
	"github.com/dexguitar/gotodoapi/internal/model"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type TodoStorage struct {
	DB *sqlx.DB
}

func NewTodoStorage(db *sqlx.DB) *TodoStorage {
	return &TodoStorage{
		DB: db,
	}
}

func (s *TodoStorage) GetTodoByTitle(ctx context.Context, title string) (*model.Todo, error) {
	const op = "TodoStorage.GetTodoByTitle"

	var todo model.Todo
	q := "select * from todos where title = $1"
	err := s.DB.QueryRowContext(ctx,
		q,
		title).Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Done)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, errs.ErrTodoNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &todo, nil
}

func (s *TodoStorage) GetTodoById(ctx context.Context, id int64) (*model.Todo, error) {
	const op = "TodoStorage.GetTodoById"

	var todo model.Todo
	q := "select * from todos where id = $1"
	err := s.DB.QueryRowContext(ctx,
		q,
		id).Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Done)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, errs.ErrTodoNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &todo, nil
}

func (s *TodoStorage) Create(ctx context.Context, todo *model.Todo) (int64, error) {
	const op = "TodoStorage.Create"

	if todo == nil {
		return 0, fmt.Errorf("%s: todo cannot be nil", op)
	}

	var id int64
	q := "insert into todos (title, content, done) values ($1, $2, $3) returning id"
	err := s.DB.QueryRowContext(ctx,
		q,
		todo.Title, todo.Content, todo.Done).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}

func (s *TodoStorage) GetAllTodos(ctx context.Context) ([]*model.Todo, error) {
	const op = "TodoStorage.GetAllTodos"

	todos := make([]*model.Todo, 0)
	q := "select * from todos"
	rows, err := s.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo model.Todo
		err := rows.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Done)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		todos = append(todos, &todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return todos, nil
}

func (s *TodoStorage) DeleteTodoById(ctx context.Context, id int64) error {
	const op = "TodoStorage.DeleteTodoById"

	q := "delete from todos where id = $1"
	_, err := s.DB.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *TodoStorage) CompleteTodo(ctx context.Context, id int64) error {
	const op = "TodoStorage.CompleteTodo"

	q := "update todos set done = true where id = $1"
	_, err := s.DB.ExecContext(ctx, q, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
