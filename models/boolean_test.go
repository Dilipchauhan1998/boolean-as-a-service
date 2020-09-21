package models

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestCreateBoolean(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) //mock sql.DB
	assert.Nil(t, err, "error while mocking database")

	gdb, err := gorm.Open("mysql", db) // open gorm db
	assert.Nil(t, err)

	booleanRepo := &BooleanRepo{db: gdb}
	booleans := []Boolean{
		{
			ID:    uuid.New().String(),
			Value: new(bool),
			Key:   "name",
		},
		{
			ID:    uuid.New().String(),
			Value: new(bool),
		},
	}

	const sqlInsert = "INSERT INTO `booleans` (`id`,`value`,`key`) VALUES (?,?,?)"

	for _, boolean := range booleans {
		mock.ExpectBegin() // start transaction
		mock.ExpectExec(sqlInsert).
			WithArgs(boolean.ID, boolean.Value, boolean.Key).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit() // commit transaction

		err = booleanRepo.CreateBoolean(&boolean)
		assert.Nil(t, err)

	}

}

func TestGetBoolean(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) //mock sql.DB
	assert.Nil(t, err, "error while mocking database")

	gdb, err := gorm.Open("mysql", db) // open gorm db
	assert.Nil(t, err)

	booleanRepo := &BooleanRepo{db: gdb}

	booleans := []Boolean{
		{
			ID:    uuid.New().String(),
			Value: new(bool),
			Key:   "name",
		},
		{
			ID:    uuid.New().String(),
			Value: new(bool),
		},
	}

	for _, boolean := range booleans {
		rows := sqlmock.
			NewRows([]string{"id", "value", "key"}).
			AddRow(boolean.ID, boolean.Value, boolean.Key)

		mock.ExpectQuery("SELECT * FROM `booleans` WHERE (id = ?)").
			WithArgs(boolean.ID).
			WillReturnRows(rows)

		var newBoolean Boolean
		err = booleanRepo.GetBooleanByID(&newBoolean, boolean.ID)
		assert.Nil(t, err)
		assert.Equal(t, newBoolean.ID, boolean.ID)
		assert.Equal(t, newBoolean.Key, boolean.Key)
		assert.Equal(t, newBoolean.Value, boolean.Value)
	}
}

func TestUpdateBoolean(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) //mock sql.DB
	assert.Nil(t, err, "error while mocking database")

	gdb, err := gorm.Open("mysql", db) // open gorm db
	assert.Nil(t, err)

	booleanRepo := &BooleanRepo{db: gdb}
	booleans := []Boolean{
		{
			ID:    uuid.New().String(),
			Value: new(bool),
			Key:   "name",
		},
		{
			ID:    uuid.New().String(),
			Value: new(bool),
		},
	}

	const sqlUpdate = "UPDATE `booleans` SET `value` = ?, `key` = ? WHERE `booleans`.`id` = ?"
	for _, boolean := range booleans {
		mock.ExpectBegin()
		mock.ExpectExec(sqlUpdate).
			WithArgs(boolean.Value, boolean.Key, boolean.ID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err = booleanRepo.UpdateBoolean(&boolean)
		assert.Nil(t, err)
	}
}

func TestDeleteBoolean(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual)) //sql.DB
	assert.Nil(t, err, "error while mocking database")

	gdb, err := gorm.Open("mysql", db) // open gorm db
	assert.Nil(t, err)

	booleanRepo := &BooleanRepo{db: gdb}
	id := uuid.New().String()

	const sqlDelete = "DELETE FROM `booleans`  WHERE (id= ?)"

	mock.ExpectBegin()
	mock.ExpectExec(sqlDelete).
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectCommit()

	var boolean Boolean
	err = booleanRepo.DeleteBoolean(&boolean, id)
	assert.Nil(t, err)
	assert.Equal(t, boolean.ID, "")
	assert.Equal(t, boolean.Value, (*bool)(nil))
	assert.Equal(t, boolean.Key, "")
}
