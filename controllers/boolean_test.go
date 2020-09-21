package controllers

import (
	"boolean-as-a-service/models"
	"testing"

	"github.com/google/uuid"
)

type BooleanRepoTest struct{}

func (br *BooleanRepoTest) CreateBoolean(boolean *models.Boolean) error {
	val := true
	boolean.ID = uuid.New().String()
	boolean.Value = &val
	boolean.Key = "my name is Dilip"
	return nil

}
func (br *BooleanRepoTest) GetBooleanByID(boolean *models.Boolean, id string) error {
	return nil
}

func (br *BooleanRepoTest) UpdateBoolean(boolean *models.Boolean) error {
	return nil
}

func (br *BooleanRepoTest) DeleteBoolean(boolean *models.Boolean, id string) error {
	return nil

}

func TestCreateBoolean(t *testing.T) {

}
