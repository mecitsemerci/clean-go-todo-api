package faker

import (
	"github.com/brianvoe/gofakeit/v5"
	"github.com/mecitsemerci/clean-go-todo-api/app/core/domain/todo"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/datetime"
	"github.com/mecitsemerci/clean-go-todo-api/app/infra/idgenerator"
)

func FakeTodo() *todo.Todo {
	t := todo.Todo{
		Id:          idgenerator.NewID(),
		Title:       gofakeit.Sentence(3),
		Description: gofakeit.Sentence(10),
		Completed:   gofakeit.Bool(),
		CreatedAt:   gofakeit.Date(),
	}
	t.UpdatedAt = gofakeit.DateRange(t.CreatedAt, datetime.Now())
	return &t
}

func FakeTodos(limit int) []*todo.Todo {
	var todos = make([]*todo.Todo, 0)
	for i := 0; i < limit; i++ {
		todos = append(todos, FakeTodo())
	}
	return todos
}