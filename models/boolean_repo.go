package models

import (
	"github.com/jinzhu/gorm"
)

type booleanRepo struct {
	db *gorm.DB
}

type booleanRepoInterface interface {
	CreateBoolean(boolean Boolean) error
	GetBooleanByID(id string) (Boolean, error)
	UpdateBoolean(boolean Boolean) error
	DeleteBooleanByID(id string) error
}

//BooleanRepo ...
var (
	BooleanRepo booleanRepoInterface
)

//SetBooleanRepo ...
func SetBooleanRepo(db *gorm.DB) {
	BooleanRepo = &booleanRepo{db: db}
}

//CreateBoolean ... insert a new boolean to database
func (br *booleanRepo) CreateBoolean(boolean Boolean) error {
	err := br.db.Create(&boolean).Error
	return err
}

//GetBooleanByID ...  select a boolean based on its id
func (br *booleanRepo) GetBooleanByID(id string) (Boolean, error) {
	boolean := Boolean{}
	err := br.db.Where(`id = ?`, id).Find(&boolean).Error
	return boolean, err
}

//UpdateBoolean ...  update a boolean
func (br *booleanRepo) UpdateBoolean(boolean Boolean) error {
	err := br.db.Save(&boolean).Error
	return err
}

//DeleteBooleanByID ... delete a boolean based on its id
func (br *booleanRepo) DeleteBooleanByID(id string) error {
	boolean := Boolean{}
	err := br.db.Where("id= ?", id).Delete(&boolean).Error
	return err
}
