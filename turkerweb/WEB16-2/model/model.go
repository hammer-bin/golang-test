package model

import (
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

type DBHandler interface {
	GetTodos() []*Todo
	AddTodo(name string) *Todo
	RemoveTodo(id int) bool
	CompleteTodo(id int, complete bool) bool
	Close()
}

func NewDBHandler(filpath string) DBHandler {
	//2. dbHandler를 반환하는 newMemoryHandler() function을 만든다.
	//handler = newMemoryHandler()

	return newSqliteHandler(filpath)
}
