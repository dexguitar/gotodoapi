package handler

import (
	"net/http"

	"github.com/dexguitar/gotodoapi/internal/model"
	"github.com/gin-gonic/gin"

	validation "github.com/go-ozzo/ozzo-validation"
)

type TodoHandler struct {
	TodoService
}

func NewTodoHandler(service TodoService) *TodoHandler {
	return &TodoHandler{
		TodoService: service,
	}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var req CreateTodoReq

	c.ShouldBindJSON(&req)
	err := req.Validate()
	if err != nil {
		c.JSON(400, &CreateTodoRes{
			ID:  0,
			Err: err.Error(),
		})
		return
	}

	newToDo := &model.Todo{
		Title:   req.Title,
		Content: req.Content,
	}

	todoID, err := h.TodoService.Create(c, newToDo)
	if err != nil {
		c.JSON(500, &CreateTodoRes{
			ID:  todoID,
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, &CreateTodoRes{
		ID: todoID,
	})
}

func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos, err := h.TodoService.GetAllTodos(c)
	if err != nil {
		c.JSON(500, &GetAllTodosRes{
			Todos: todos,
			Err:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &GetAllTodosRes{
		Todos: todos,
	})
}

func (h *TodoHandler) DeleteTodoById(c *gin.Context) {
	var req DeleteTodoByIdReq

	c.ShouldBindJSON(&req)
	err := req.Validate()
	if err != nil {
		c.JSON(400, &DeleteTodoByIdRes{
			Err: err.Error(),
		})
		return
	}

	err = h.TodoService.DeleteTodoById(c, req.ID)
	if err != nil {
		c.JSON(500, &DeleteTodoByIdRes{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &DeleteTodoByIdRes{
		Err: "",
	})
}

func (h *TodoHandler) CompleteTodo(c *gin.Context) {
	var req CompleteTodoReq

	c.ShouldBindJSON(&req)
	err := req.Validate()
	if err != nil {
		c.JSON(400, &CompleteTodoRes{
			Err: err.Error(),
		})
		return
	}

	err = h.TodoService.CompleteTodo(c, req.ID)
	if err != nil {
		c.JSON(500, &CompleteTodoRes{
			Err: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &CompleteTodoRes{
		Err: "",
	})
}

type CreateTodoReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type CreateTodoRes struct {
	ID  int64  `json:"id"`
	Err string `json:"error"`
}

type GetAllTodosRes struct {
	Todos []*model.Todo `json:"todos"`
	Err   string        `json:"error"`
}

type DeleteTodoByIdReq struct {
	ID int64 `json:"id"`
}

type DeleteTodoByIdRes struct {
	Err string `json:"error"`
}

type CompleteTodoReq struct {
	ID int64 `json:"id"`
}

type CompleteTodoRes struct {
	Err string `json:"error"`
}

func (r CreateTodoReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.Title, validation.Required),
		validation.Field(&r.Content, validation.Required),
	)
}

func (r DeleteTodoByIdReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.ID, validation.Required),
	)
}

func (r CompleteTodoReq) Validate() error {
	return validation.ValidateStruct(
		&r,
		validation.Field(&r.ID, validation.Required),
	)
}
