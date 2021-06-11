package model

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) GetTodos() []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query(`SELECT id, name, completed, createdAt FROM todos`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt)
		todos = append(todos, &todo)
	}

	return todos
}

func (s *sqliteHandler) AddTodo(name string) *Todo {
	statement, err := s.db.Prepare(`INSERT INTO todos (name, completed, createdAt) VALUES (?,?,DATETIME('now'))`)
	if err != nil {
		panic(err)
	}

	rst, err := statement.Exec(name, false)
	if err != nil {
		panic(err)
	}

	// 추가한 데이터(id)를 알려줌
	id, _ := rst.LastInsertId()
	var todo Todo
	todo.ID = int(id)
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()

	return &todo
}
func (s *sqliteHandler) RemoveTodo(id int) bool {
	return false
}
func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	return false
}
func (s *sqliteHandler) Close() {
	s.db.Close()
}

func newSqliteHandler() DBHandler {
	database, err := sql.Open("sqlite3", "./test.db")

	if err != nil {
		panic(err)
	}

	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id					INTEGER PRIMARY KEY AUTOINCREMENT,
			name				TEXT,
			completed		BOOLEAN,
			createdAt		DATETIME
		)`)

	statement.Exec()

	return &sqliteHandler{db: database}
}
