package service

import (
	// "errors"
	"errors"
	"testing"

	todo "github.com/MukhammedAli/GoFinalProject"
	"github.com/MukhammedAli/GoFinalProject/pkg/repository"
	"github.com/MukhammedAli/GoFinalProject/pkg/repository/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewAuthService(t *testing.T) {
	type arguments struct {
		repo repository.Authorization
	}

	type expectations struct {
		auth *AuthService
	}

	type test struct {
		arguments    arguments
		expectations expectations
	}
	repo := &mocks.Authorization{}
	auth := &AuthService{repo: repo}

	testCases := []test{
		{
			arguments: arguments{
				repo: repo,
			},
			expectations: expectations{
				auth: auth,
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

		auth := NewAuthService(tc.arguments.repo)
		assert.Equal(t, tc.expectations.auth, auth)

	}
}

func TestAuthService_CreateUser(t *testing.T) {
	type arguments struct {
		user todo.User
	}

	type expectations struct {
		id  int
		err error
	}

	type dependencies struct {
		repo *mocks.Authorization
	}

	type test struct {
		arguments    arguments
		expectations expectations
		dependencies dependencies
		prepare      func(repo *mocks.Authorization, id int, err error)
	}

	user := todo.User{}

	testCases := []test{
		{
			arguments: arguments{
				user: user,
			},
			expectations: expectations{
				id:  user.Id,
				err: nil,
			},
			dependencies: dependencies{
				repo: &mocks.Authorization{},
			},
			prepare: func(repo *mocks.Authorization, id int, err error) {
				repo.On("CreateUser", mock.Anything, mock.Anything).Return(id, err)
			},
		},
		{
			arguments: arguments{
				user: user,
			},
			expectations: expectations{
				id:  -1,
				err: errors.New("random error"),
			},
			dependencies: dependencies{
				repo: &mocks.Authorization{},
			},
			prepare: func(repo *mocks.Authorization, id int, err error) {
				repo.On("CreateUser", mock.Anything, mock.Anything).Return(id, err)
			},
		},
	}
	for _, tc := range testCases {

		tc.prepare(tc.dependencies.repo, tc.expectations.id, tc.expectations.err)

		s := AuthService{
			repo: tc.dependencies.repo,
		}

		id, err := s.CreateUser(tc.arguments.user)

		assert.Equal(t, tc.expectations.id, id)
		assert.Equal(t, tc.expectations.err, err)

	}
}
