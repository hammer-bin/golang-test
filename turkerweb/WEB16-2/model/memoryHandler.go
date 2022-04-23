package model

import "time"

// 1. memoryHander는 dbHandler를 구현하고 있다.
//    그러면 dbHandler만 들고 있으면 된다. 전역변수로 dbHandler를 선언(handler)
type memoryHandler struct {
	todoMap map[int]*Todo
}

func (m *memoryHandler) Close() {

}

func (m *memoryHandler) GetTodos() []*Todo {
	list := []*Todo{}
	for _, v := range m.todoMap {
		list = append(list, v)
	}
	return list
}

func (m *memoryHandler) AddTodo(name string) *Todo {
	id := len(m.todoMap) + 1
	todo := &Todo{id, name, false, time.Now()}
	m.todoMap[id] = todo
	return todo
}

func (m *memoryHandler) RemoveTodo(id int) bool {
	if _, ok := m.todoMap[id]; ok {
		delete(m.todoMap, id)
		return true
	}
	return false
}

func (m *memoryHandler) CompleteTodo(id int, complete bool) bool {
	if todo, ok := m.todoMap[id]; ok {
		todo.Completed = complete
		return true
	}
	return false
}

func newMemoryHandler() DBHandler {
	//3. memoryHandler를 init해서 반환
	//4. memoryHandler가 dbHandler를 구현하고 있기 때문에 dbHandler로 반환할수 있다.
	m := &memoryHandler{}
	m.todoMap = make(map[int]*Todo)
	return m
}
