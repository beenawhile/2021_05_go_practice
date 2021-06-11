package model

import (
	// database들과 연결하는 interface
	"database/sql"
	"time"

	// _ : 명시적이 아니라 암시적으로 사용하겠음
	_ "github.com/mattn/go-sqlite3"
)

type sqliteHandler struct {
	db *sql.DB
}

func (s *sqliteHandler) Close() {
	s.db.Close()
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
func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
	stmt, err := s.db.Prepare(`UPDATE todos SET completed = ? WHERE ID = ?`)
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

func newSqliteHandler(filepath string) DBHandler {
	database, err := sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}

	// data 저장할 테이블 만듬
	statement, _ := database.Prepare(
		`CREATE TABLE IF NOT EXISTS todos (
			id				INTEGER PRIMARY KEY AUTOINCREMENT,
			name 			TEXT,
			completed BOOLEAN,
			createdAt	DATETIME
		)`)

	// query문 실행
	statement.Exec()

	return &sqliteHandler{db: database}
}
