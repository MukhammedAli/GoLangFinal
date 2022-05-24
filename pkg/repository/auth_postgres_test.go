package repository

import (
	"errors"
	"testing"

	todo "github.com/MukhammedAli/GoFinalProject"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthService(t *testing.T) {
	type arguments struct {
		db *sqlx.DB
	}

	type expectations struct {
		auth *AuthPostgres
	}

	type test struct {
		arguments    arguments
		expectations expectations
	}
	db := &sqlx.DB{}
	auth := &AuthPostgres{db: db}

	testCases := []test{
		{
			arguments: arguments{
				db: db,
			},
			expectations: expectations{
				auth: auth,
			},
		},
		{
			arguments: arguments{
				db: nil,
			},
		},
	}
	for _, tc := range testCases {
		defer func() {
			recover()
		}()

		auth := NewAuthPostgres(tc.arguments.db)
		assert.Equal(t, tc.expectations.auth, auth)

	}
}

func TestAuthPostgres_CreateUser(t *testing.T) {
	type arguments struct {
		user todo.User
	}

	type expectations struct {
		id  int
		err error
	}

	type dependencies struct {
		db *sqlx.DB
	}

	type test struct {
		arguments    arguments
		expectations expectations
		dependencies dependencies
		prepare      func(db *sqlx.DB, id int, err error)
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
				db: &sqlx.DB{},
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
				db: &sqlx.DB{},
			},
			// prepare: func(repo *sqlx.DB, id int, err error) {
			// 	repo.On("CreateUser", mock.Anything, mock.Anything).Return(id, err)
			// },
		},
	}
	for _, tc := range testCases {

		// tc.prepare(tc.dependencies.repo, tc.expectations.id, tc.expectations.err)

		s := AuthPostgres{
			db: tc.dependencies.db,
		}

		id, err := s.CreateUser(tc.arguments.user)

		assert.Equal(t, tc.expectations.id, id)
		assert.Equal(t, tc.expectations.err, err)

	}
}

