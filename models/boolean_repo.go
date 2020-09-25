package models

import (
	"github.com/jinzhu/gorm"
)

type booleanRepo struct {
	db *gorm.DB
}

type booleanRepoInterface interface {
	CreateBoolean(boolean *Boolean) error
	GetBooleanByID(boolean *Boolean, id string) error
	UpdateBoolean(boolean *Boolean) error
	DeleteBooleanByID(boolean *Boolean, id string) (err error)
}

//BooleanRepo ...
var (
	BooleanRepo booleanRepoInterface
)

//NewBooleanRepo ...
func NewBooleanRepo(db *gorm.DB) {
	BooleanRepo = &booleanRepo{db: db}
}

//CreateBoolean ... insert a new boolean to database
func (br *booleanRepo) CreateBoolean(boolean *Boolean) error {
	if err := br.db.Create(boolean).Error; err != nil {
		return err
	}
	return nil
}

//GetBooleanByID ...  select a boolean based on its id
func (br *booleanRepo) GetBooleanByID(boolean *Boolean, id string) error {
	if err := br.db.Where(`id = ?`, id).Find(boolean).Error; err != nil {
		return err
	}
	return nil
}

//UpdateBoolean ...  update a boolean
func (br *booleanRepo) UpdateBoolean(boolean *Boolean) error {
	if err := br.db.Save(&boolean).Error; err != nil {
		return err
	}
	return nil
}

//DeleteBooleanByID ... delete a boolean based on its id
func (br *booleanRepo) DeleteBooleanByID(boolean *Boolean, id string) error {
	if err := br.db.Where("id= ?", id).Delete(boolean).Error; err != nil {
		return err
	}
	return nil
}
