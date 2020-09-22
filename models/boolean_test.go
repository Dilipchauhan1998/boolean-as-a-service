package models

import (
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type testBoolean struct {
	boolean  Boolean
	sqlQuery string
	err      error
}

func TestCreateBoolean(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) //mock sql.DB
	assert.Nil(t, err, "error while mocking database")

	gdb, err := gorm.Open("mysql", db) // open gorm db
	assert.Nil(t, err)

	NewBooleanRepo(gdb)
	booleanRepo := BooleanRepo

	testBooleans := []testBoolean{
		{
			boolean: Boolean{
				ID:    uuid.New().String(),
				Value: new(bool),
				Key:   "name",
			},
			sqlQuery: "INSERT INTO `booleans` (`id`,`value`,`key`) VALUES (?,?,?)",
			err:      nil,
		},
		{
			boolean: Boolean{
				ID:    uuid.New().String(),
				Value: new(bool),
			},
			sqlQuery: "INSERT INTO `booleans` (`id`,`value`,`key`) VALUES (?,?,?)",
			err:      nil,
		},
		{
			boolean: Boolean{
				ID:    uuid.New().String(),
				Value: new(bool),
				Key:   "name",
			},
			sqlQuery: "INSERT INTO `booleans` (`id`,`value`) VALUES (?,?)",
			err:      errors.New("syntax error"),
		},
	}

	for _, testboolean := range testBooleans {
		mock.ExpectBegin() // start transaction
		mock.ExpectExec(testboolean.sqlQuery).
			WithArgs(testboolean.boolean.ID, testboolean.boolean.Value, testboolean.boolean.Key).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit() // commit transaction

		err = booleanRepo.CreateBoolean(&testboolean.boolean)

		if testboolean.err != nil {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}

}

func TestGetBoolean(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) //mock sql.DB
	assert.Nil(t, err, "error while mocking database")

	gdb, err := gorm.Open("mysql", db) // open gorm db
	assert.Nil(t, err)

	NewBooleanRepo(gdb)
	booleanRepo := BooleanRepo

	testBooleans := []testBoolean{
		{
			boolean: Boolean{
				ID:    uuid.New().String(),
				Value: new(bool),
				Key:   "name",
			},
			sqlQuery: "SELECT * FROM `booleans` WHERE (id = ?)",
			err:      nil,
		},
		{
			boolean: Boolean{
				ID:    uuid.New().String(),
				Value: new(bool),
			},
			sqlQuery: "SELECT * FROM `booleans` WHERE (id = ?)",
			err:      nil,
		},
		{
			boolean: Boolean{
				ID:    uuid.New().String(),
				Value: new(bool),
				Key:   "name",
			},
			sqlQuery: "SELECT `id` FROM `booleans` WHERE (id = ?)",
			err:      errors.New("syntax error"),
		},
	}

	for _, testboolean := range testBooleans {
		rows := sqlmock.
			NewRows([]string{"id", "value", "key"}).
			AddRow(testboolean.boolean.ID, testboolean.boolean.Value, testboolean.boolean.Key)

		mock.ExpectQuery(testboolean.sqlQuery).
			WithArgs(testboolean.boolean.ID).
			WillReturnRows(rows)

		var newBoolean Boolean
		err = booleanRepo.GetBooleanByID(&newBoolean, testboolean.boolean.ID)

		if testboolean.err != nil {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
			assert.Equal(t, newBoolean.ID, testboolean.boolean.ID)
			assert.Equal(t, newBoolean.Key, testboolean.boolean.Key)
			assert.Equal(t, newBoolean.Value, testboolean.boolean.Value)
		}
	}
}

func TestUpdateBoolean(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) //mock sql.DB
	assert.Nil(t, err, "error while mocking database")

	gdb, err := gorm.Open("mysql", db) // open gorm db
	assert.Nil(t, err)

	NewBooleanRepo(gdb)
	booleanRepo := BooleanRepo

	testBooleans := []testBoolean{
		{
			boolean: Boolean{
				ID:    uuid.New().String(),
				Value: new(bool),
				Key:   "name",
			},
			sqlQuery: "UPDATE `booleans` SET `value` = ?, `key` = ? WHERE `booleans`.`id` = ?",
			err:      nil,
		},
		{
			boolean: Boolean{
				ID:    uuid.New().String(),
				Value: new(bool),
			},
			sqlQuery: "UPDATE `booleans` SET `value` = ?, `key` = ? WHERE `booleans`.`id` = ?",
			err:      nil,
		},
		{
			boolean: Boolean{
				ID:    uuid.New().String(),
				Value: new(bool),
				Key:   "name",
			},
			sqlQuery: "UPDATE `booleans` SET `value` = ? WHERE `booleans`.`id` = ?",
			err:      errors.New("syntax error"),
		},
	}

	for _, testboolean := range testBooleans {
		mock.ExpectBegin()
		mock.ExpectExec(testboolean.sqlQuery).
			WithArgs(testboolean.boolean.Value, testboolean.boolean.Key, testboolean.boolean.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = booleanRepo.UpdateBoolean(&testboolean.boolean)

		if testboolean.err != nil {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestDeleteBoolean(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) //sql.DB
	assert.Nil(t, err, "error while mocking database")

	gdb, err := gorm.Open("mysql", db) // open gorm db
	assert.Nil(t, err)

	NewBooleanRepo(gdb)
	booleanRepo := BooleanRepo

	testBooleans := []testBoolean{
		{
			boolean: Boolean{
				ID: uuid.New().String(),
			},
			sqlQuery: "DELETE FROM `booleans`  WHERE (id= ?)",
			err:      nil,
		},
		{
			boolean: Boolean{
				ID: uuid.New().String(),
			},
			sqlQuery: "DELETE FROM `booleans` ",
			err:      errors.New("some error"),
		},
	}

	for _, testboolean := range testBooleans {

		mock.ExpectBegin()
		mock.ExpectExec(testboolean.sqlQuery).
			WithArgs(testboolean.boolean.ID).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		var boolean Boolean
		err = booleanRepo.DeleteBooleanByID(&boolean, testboolean.boolean.ID)

		if testboolean.err != nil {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

	}
}
