package model

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func newSqliteHandler(filepath string) DBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	stmt, err := database.Prepare(`CREATE TABLE IF NOT EXISTS todos (
		id					INTEGER PRIMARY KEY AUTOINCREMENT,
		name				TEXT,
		completed		BOOLEAN,
		createdAt		DATETIME
	)`)

	if err != nil {
		panic(err)
	}
	stmt.Exec()

	s := &sqliteHandler{
		db: database,
	}
	return s
}

func (s *sqliteHandler) Close() {
	s.db.Close()
}

func (s *sqliteHandler) GetTodos() []*Todo {
	rows, err := s.db.Query(`SELECT id, name, completed, createdAt FROM todos`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	todos := []*Todo{}

	for rows.Next() {
		var todo Todo
		rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt)
		todos = append(todos, &todo)
	}

	return todos
}

func (s *sqliteHandler) AddTodo(name string) *Todo {

	stmt, err := s.db.Prepare(`INSERT INTO todos (name, completed, createdAt) VALUES (?,?,datetime('now'))`)
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(name, false)

	if err != nil {
		panic(err)
	}
	var todo Todo
	id, _ := rst.LastInsertId()
	todo.ID = int(id)
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()

	return &todo
}

func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {

	stmt, err := s.db.Prepare(`UPDATE todos SET completed = ? WHERE id = ?`)
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(complete, id)
	if err != nil {
		panic(err)
	}

	cnt, _ := rst.RowsAffected()
	return cnt > 0
}

func (s *sqliteHandler) RemoveTodo(id int) bool {
	stmt, err := s.db.Prepare(`DELETE FROM todos WHERE id = ?`)
	if err != nil {
		panic(err)
	}
	rst, err := stmt.Exec(id)
	if err != nil {
		panic(err)
	}

	cnt, _ := rst.RowsAffected()

	return cnt > 0
}
