package model

import (
	// database들과 연결하는 interface
	"database/sql"
	"time"

	// _ : 명시적이 아니라 암시적으로 사용하겠음
	_ "github.com/lib/pq"
)

type PQHandler struct {
	db *sql.DB
}

func (s *PQHandler) Close() {

	s.db.Close()
}

func (s *PQHandler) GetTodos(seessionId string) []*Todo {
	todos := []*Todo{}
	rows, err := s.db.Query(`SELECT id, name, completed, createdAt FROM todos WHERE sessionId = $1`, seessionId)
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

func (s *PQHandler) AddTodo(name string, sessionId string) *Todo {
	stmt, err := s.db.Prepare(`INSERT INTO todos (sessionId, name, completed, createdAt) VALUES ($1,$2,$3,NOW()) RETURNING id`)
	if err != nil {
		panic(err)
	}

	var id int
	err = stmt.QueryRow(sessionId, name, false).Scan(&id)
	if err != nil {
		panic(err)
	}

	if err != nil {
		panic(err)
	}

	// 추가한 데이터(id)를 알려줌
	var todo Todo
	todo.ID = id
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()

	return &todo
}

func (s *PQHandler) RemoveTodo(id int) bool {
	stmt, err := s.db.Prepare(`DELETE FROM todos WHERE id = $1`)
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
func (s *PQHandler) CompleteTodo(id int, complete bool) bool {
	stmt, err := s.db.Prepare(`UPDATE todos SET completed = $1 WHERE ID = $2`)
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

func newPQHandler(dbConn string) DBHandler {
	database, err := sql.Open("postgres", dbConn)
	if err != nil {
		panic(err)
	}

	// data 저장할 테이블 만듬
	statement, err := database.Prepare(
		`
		CREATE TABLE IF NOT EXISTS todos (
			id				SERIAL PRIMARY KEY,
			sessionId VARCHAR(256),
			name 			TEXT,
			completed BOOLEAN,
			createdAt	TIMESTAMP
		);`)
	if err != nil {
		panic(err)
	}

	_, err = database.Prepare(`CREATE INDEX IF NOT EXISTS sessionIdIndexOnTodos ON todos (
			sessionId ASC
		);
		`)
	if err != nil {
		panic(err)
	}

	// query문 실행
	_, err = statement.Exec()
	if err != nil {
		panic(err)
	}

	return &PQHandler{db: database}
}
