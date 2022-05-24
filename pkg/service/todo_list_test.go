package service

import (
	"errors"
	"testing"

	todo "github.com/MukhammedAli/GoFinalProject"
	"github.com/MukhammedAli/GoFinalProject/pkg/repository"
	"github.com/MukhammedAli/GoFinalProject/pkg/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewTodoListService(t *testing.T) {
	type arguments struct {
		repo repository.TodoList
	}

	type expectations struct {
		todoList *TodoListService
	}

	type test struct {
		arguments    arguments
		expectations expectations
	}

	repo := &mocks.TodoList{}
	todoList := &TodoListService{repo: repo}

	testCases := []test{
		{
			arguments: arguments{
				repo: repo,
			},
			expectations: expectations{
				todoList: todoList,
			},
		},
		{
			arguments: arguments{
				repo: nil,
			},
		},
	}
	for _, tc := range testCases {
		defer func() {
			recover()
		}()

		todoList := NewTodoListService(tc.arguments.repo)
		assert.Equal(t, tc.expectations.todoList, todoList)

	}
}

func TestTodoListService_Create(t *testing.T) {
	type arguments struct {
		userId int
		list   todo.TodoList
	}

	type expectations struct {
		id  int
		err error
	}

	type dependencies struct {
		repo *mocks.TodoList
	}

	type test struct {
		arguments    arguments
		expectations expectations
		dependencies dependencies
		prepare      func(repo *mocks.TodoList, id int, err error)
	}

	userId := todo.UserList{}.UserId
	list := todo.TodoList{}
	id := todo.TodoList{}.Id

	testCases := []test{
		{
			arguments: arguments{
				userId: userId,
				list:   list,
			},
			expectations: expectations{
				id:  id,
				err: nil,
			},
			dependencies: dependencies{
				repo: &mocks.TodoList{},
			},
			prepare: func(repo *mocks.TodoList, id int, err error) {
				repo.On("Create", mock.Anything, mock.Anything).Return(id, err)
			},
		},
		{
			arguments: arguments{
				userId: userId,
				list:   list,
			},
			expectations: expectations{
				id:  0,
				err: errors.New("random error"),
			},
			dependencies: dependencies{
				repo: &mocks.TodoList{},
			},
			prepare: func(repo *mocks.TodoList, id int, err error) {
				repo.On("Create", mock.Anything, mock.Anything).Return(id, err)
			},
		},
	}
	for _, tc := range testCases {

		tc.prepare(tc.dependencies.repo, tc.expectations.id, tc.expectations.err)

		s := TodoListService{
			repo: tc.dependencies.repo,
		}

		id, err := s.Create(tc.arguments.userId, tc.arguments.list)

		assert.Equal(t, tc.expectations.id, id)
		assert.Equal(t, tc.expectations.err, err)

	}
}

func TestTodoListService_GetById(t *testing.T) {
	type arguments struct {
		userId int
		listId int
	}

	type expectations struct {
		todoList todo.TodoList
		err      error
	}

	type dependencies struct {
		repo *mocks.TodoList
	}

	type test struct {
		arguments    arguments
		expectations expectations
		dependencies dependencies
		prepare      func(repo *mocks.TodoList, todoList todo.TodoList, err error)
	}

	userId := todo.UserList{}.Id
	listId := todo.TodoList{}.Id
	todoList := todo.TodoList{}
	testCases := []test{
		{
			arguments: arguments{
				userId: userId,
				listId: listId,
			},
			expectations: expectations{
				todoList: todoList,
				err:      nil,
			},
			dependencies: dependencies{
				repo: &mocks.TodoList{},
			},
			prepare: func(repo *mocks.TodoList, todoList todo.TodoList, err error) {
				repo.On("GetById", mock.Anything, mock.Anything).Return(todoList, err)
			},
		},
		{
			arguments: arguments{
				userId: userId,
				listId: listId,
			},
			expectations: expectations{
				todoList: todoList,
				err:      errors.New("random error"),
			},
			dependencies: dependencies{
				repo: &mocks.TodoList{},
			},
			prepare: func(repo *mocks.TodoList, todoList todo.TodoList, err error) {
				repo.On("GetById", mock.Anything, mock.Anything).Return(todoList, err)
			},
		},
	}
	for _, tc := range testCases {

		tc.prepare(tc.dependencies.repo, tc.expectations.todoList, tc.expectations.err)

		s := TodoListService{
			repo: tc.dependencies.repo,
		}

		todoList, err := s.GetById(tc.arguments.userId, tc.arguments.listId)

		assert.Equal(t, tc.expectations.todoList, todoList)
		assert.Equal(t, tc.expectations.err, err)

	}
}

func TestTodoListService_Delete(t *testing.T) {
	type arguments struct {
		userId int
		listId int
	}

	type expectations struct {
		err error
	}

	type dependencies struct {
		repo *mocks.TodoList
	}

	type test struct {
		arguments    arguments
		expectations expectations
		dependencies dependencies
		prepare      func(repo *mocks.TodoList, err error)
	}

	userId := todo.UserList{}.Id
	listId := todo.TodoList{}.Id

	testCases := []test{
		{
			arguments: arguments{
				userId: userId,
				listId: listId,
			},
			expectations: expectations{
				err: nil,
			},
			dependencies: dependencies{
				repo: &mocks.TodoList{},
			},
			prepare: func(repo *mocks.TodoList, err error) {
				repo.On("Delete", mock.Anything, mock.Anything).Return(err)
			},
		},
		{
			arguments: arguments{
				userId: userId,
				listId: listId,
			},
			expectations: expectations{
				err: errors.New("random error"),
			},
			dependencies: dependencies{
				repo: &mocks.TodoList{},
			},
			prepare: func(repo *mocks.TodoList, err error) {
				repo.On("Delete", mock.Anything, mock.Anything).Return(err)
			},
		},
	}
	for _, tc := range testCases {

		tc.prepare(tc.dependencies.repo, tc.expectations.err)

		s := TodoListService{
			repo: tc.dependencies.repo,
		}

		err := s.Delete(tc.arguments.userId, tc.arguments.listId)

		assert.Equal(t, tc.expectations.err, err)

	}
}
