package auth

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"

	"github.com/DATA-DOG/go-sqlmock"
)

type TestToken struct {
	token    string
	sqlQuery string
	err      error
}

func TestCreateToken(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // mock sql DB
	assert.Nil(t, err, "error while mocking sql database")

	gdb, err := gorm.Open("mysql", db) //oprn gorm DB
	defer gdb.Close()
	assert.Nil(t, err)

	SetTokenRepo(gdb)
	tokenRepo := TokenRepo

	testTokens := []TestToken{
		{
			token:    uuid.New().String(),
			sqlQuery: "INSERT INTO `tokens` (`id`) VALUES (?)",
			err:      nil,
		},
		{
			token:    uuid.New().String(),
			sqlQuery: "INSERT INTO `tokens` VALUES(?)",
			err:      errors.New("some error"),
		},
	}

	for _, testToken := range testTokens {

		mock.ExpectBegin()
		mock.ExpectExec(testToken.sqlQuery).
			WithArgs(testToken.token).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		err := tokenRepo.CreateToken(Token{ID: testToken.token})

		if testToken.err != nil {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}

func TestExistToken(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) // mock sql DB
	assert.Nil(t, err)

	gdb, err := gorm.Open("mysql", db) // open gorm DB
	assert.Nil(t, err)

	SetTokenRepo(gdb)
	tokenRepo := TokenRepo

	testTokens := []TestToken{
		{
			token:    uuid.New().String(),
			sqlQuery: "SELECT * FROM `tokens` WHERE (id = ?)",
			err:      nil,
		},
		{
			token:    uuid.New().String(),
			sqlQuery: "SELECT * FROM `tokens` WHERE `id`= ?",
			err:      errors.New(" some error"),
		},
	}

	for _, testToken := range testTokens {
		rows := sqlmock.
			NewRows([]string{"id"}).
			AddRow(testToken.token)

		mock.ExpectQuery(testToken.sqlQuery).
			WithArgs(testToken.token).
			WillReturnRows(rows)

		doesExist := tokenRepo.ExistToken(testToken.token)
		fmt.Println(doesExist)

		if testToken.err != nil {
			assert.Equal(t, doesExist, false)
		} else {
			assert.Equal(t, doesExist, true)
		}

	}

}
