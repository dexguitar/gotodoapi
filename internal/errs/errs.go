package errs

import "errors"

var ErrTodoNotFound = errors.New("todo not found")
var ErrInvalidCreds = errors.New("invalid username or password")
var ErrInternal = errors.New("internal server error")
var ErrTodoExists = errors.New("todo already exists")
