package model

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
}

type DBHandler interface {
	GetTodos() []*Todo
	AddTodo(string) *Todo
	CompleteTodo(int, bool) bool
	RemoveTodo(int) bool
	Close()
}

func NewDBHandler(filepath string) DBHandler {
	return newSqliteHandler(filepath)
}
