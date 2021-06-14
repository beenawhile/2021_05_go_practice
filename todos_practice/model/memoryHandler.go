package model

import "time"

type memoryHandler struct {
	todoMap map[int]*Todo
}

func (m *memoryHandler) GetTodos() []*Todo {
	list := []*Todo{}
	for _, todo := range m.todoMap {
		list = append(list, todo)
	}
	return list
}

func (m *memoryHandler) AddTodo(name string) *Todo {
	id := len(m.todoMap)
	var todo Todo
	todo.ID = id
	todo.Name = name
	todo.Completed = false
	todo.CreatedAt = time.Now()
	m.todoMap[id] = &todo
	return &todo
}

func (m *memoryHandler) CompleteTodo(id int, complete bool) bool {

	if todo, ok := m.todoMap[id]; ok {
		todo.Completed = complete
		return true
	}
	return false

}

func (m *memoryHandler) RemoveTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id)
		return true
	}
	return false
}

func (m *memoryHandler) Close() {}

func newMemoryHandler() DBHandler {
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	return m
}
