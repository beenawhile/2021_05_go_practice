package model

import (
	// database들과 연결하는 interface
	"database/sql"

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

	return nil
}

func (s *sqliteHandler) AddTodo(name string) *Todo {
	return nil
}

func (s *sqliteHandler) RemoveTodo(id int) bool {

	return false
}
func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {

	return false
}

func newSqliteHandler() DBHandler {
	database, err := sql.Open("sqlite3", "./test.db")
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
