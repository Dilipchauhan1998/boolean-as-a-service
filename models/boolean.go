package models

import (
	"github.com/jinzhu/gorm"
)

//BooleanRepoInterface ...
type BooleanRepoInterface interface {
	CreateBoolean(boolean *Boolean) error
	GetBooleanByID(boolean *Boolean, id string) error
	UpdateBoolean(boolean *Boolean) error
	DeleteBoolean(boolean *Boolean, id string) (err error)
}

//BooleanRepo ...
type BooleanRepo struct {
	db *gorm.DB
}

//NewBooleanRepo ...
func NewBooleanRepo(db *gorm.DB) *BooleanRepo {
	return &BooleanRepo{db: db}
}

//CreateBoolean ... insert new Boolean to database
func (br *BooleanRepo) CreateBoolean(boolean *Boolean) error {
	if err := br.db.Create(boolean).Error; err != nil {
		return err
	}
	return nil
}

//GetBooleanByID ...  select a boolean based on its id
func (br *BooleanRepo) GetBooleanByID(boolean *Boolean, id string) error {
	if err := br.db.Where(`id = ?`, id).Find(boolean).Error; err != nil {
		return err
	}
	return nil
}

//UpdateBoolean ...  update a boolean based on its id
func (br *BooleanRepo) UpdateBoolean(boolean *Boolean) error {
	if err := br.db.Save(&boolean).Error; err != nil {
		return err
	}
	return nil
}

//DeleteBoolean ... delete a boolean based on its id
func (br *BooleanRepo) DeleteBoolean(boolean *Boolean, id string) error {
	if err := br.db.Where("id= ?", id).Delete(boolean).Error; err != nil {
		return err
	}
	return nil
}
