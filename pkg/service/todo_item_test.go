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

func TestNewToDoItemService(t *testing.T) {
	type arguments struct {
		repo     repository.TodoItem
		listRepo repository.TodoList
	}

	type expectations struct {
		todoItem *TodoItemService
	}

	type test struct {
		arguments    arguments
		expectations expectations
	}

	repo := &mocks.TodoItem{}
	listRepo := &mocks.TodoList{}
	todoItem := &TodoItemService{repo: repo, listRepo: listRepo}

	testCases := []test{
		{
			arguments: arguments{
				repo:     repo,
				listRepo: listRepo,
			},
			expectations: expectations{
				todoItem: todoItem,
			},
		},
		{
			arguments: arguments{
				repo:     nil,
				listRepo: nil,
			},
		},
	}
	for _, tc := range testCases {
		defer func() {
			recover()
		}()

		todoItem := NewTodoItemService(tc.arguments.repo, tc.arguments.listRepo)
		assert.Equal(t, tc.expectations.todoItem, todoItem)

	}
}

func TestTodoItemService_GetById(t *testing.T) {
	type arguments struct {
		userId int
		itemId int
	}

	type expectations struct {
		todoItem todo.TodoItem
		err      error
	}

	type dependencies struct {
		repo     *mocks.TodoItem
		listRepo *mocks.TodoList
	}

	type test struct {
		arguments    arguments
		expectations expectations
		dependencies dependencies
		prepare      func(repo *mocks.TodoItem, listRepo *mocks.TodoList, todoItem todo.TodoItem, err error)
	}

	userId := todo.UserList{}.Id
	itemId := todo.TodoItem{}.Id
	todoItem := todo.TodoItem{}
	testCases := []test{
		{
			arguments: arguments{
				userId: userId,
				itemId: itemId,
			},
			expectations: expectations{
				todoItem: todoItem,
				err:      nil,
			},
			dependencies: dependencies{
				repo:     &mocks.TodoItem{},
				listRepo: &mocks.TodoList{},
			},
			prepare: func(repo *mocks.TodoItem, listRepo *mocks.TodoList, todoItem todo.TodoItem, err error) {
				repo.On("GetById", mock.Anything, mock.Anything).Return(todoItem, err)
			},
		},
		{
			arguments: arguments{
				userId: userId,
				itemId: itemId,
			},
			expectations: expectations{
				todoItem: todoItem,
				err:      errors.New("random error"),
			},
			dependencies: dependencies{
				repo: &mocks.TodoItem{},
			},
			prepare: func(repo *mocks.TodoItem, listRepo *mocks.TodoList, todoItem todo.TodoItem, err error) {
				repo.On("GetById", mock.Anything, mock.Anything).Return(todoItem, err)
			},
		},
	}
	for _, tc := range testCases {

		tc.prepare(tc.dependencies.repo, tc.dependencies.listRepo, tc.expectations.todoItem, tc.expectations.err)

		s := TodoItemService{
			repo:     tc.dependencies.repo,
			listRepo: tc.dependencies.listRepo,
		}

		todoItem, err := s.GetById(tc.arguments.userId, tc.arguments.itemId)

		assert.Equal(t, tc.expectations.todoItem, todoItem)
		assert.Equal(t, tc.expectations.err, err)

	}
}

func TestTodoItemService_Delete(t *testing.T) {
	type arguments struct {
		userId int
		itemId int
	}

	type expectations struct {
		err error
	}

	type dependencies struct {
		repo     *mocks.TodoItem
		listRepo *mocks.TodoList
	}

	type test struct {
		arguments    arguments
		expectations expectations
		dependencies dependencies
		prepare      func(repo *mocks.TodoItem, listRepo *mocks.TodoList, err error)
	}

	userId := todo.UserList{}.Id
	itemId := todo.TodoItem{}.Id

	testCases := []test{
		{
			arguments: arguments{
				userId: userId,
				itemId: itemId,
			},
			expectations: expectations{
				err: nil,
			},
			dependencies: dependencies{
				repo:     &mocks.TodoItem{},
				listRepo: &mocks.TodoList{},
			},
			prepare: func(repo *mocks.TodoItem, listRepo *mocks.TodoList, err error) {
				repo.On("Delete", mock.Anything, mock.Anything).Return(err)
			},
		},
		{
			arguments: arguments{
				userId: userId,
				itemId: itemId,
			},
			expectations: expectations{
				err: errors.New("random error"),
			},
			dependencies: dependencies{
				repo: &mocks.TodoItem{},
			},
			prepare: func(repo *mocks.TodoItem, listRepo *mocks.TodoList, err error) {
				repo.On("Delete", mock.Anything, mock.Anything).Return(err)
			},
		},
	}
	for _, tc := range testCases {

		tc.prepare(tc.dependencies.repo, tc.dependencies.listRepo, tc.expectations.err)

		s := TodoItemService{
			repo:     tc.dependencies.repo,
			listRepo: tc.dependencies.listRepo,
		}

		err := s.Delete(tc.arguments.userId, tc.arguments.itemId)

		assert.Equal(t, tc.expectations.err, err)

	}
}