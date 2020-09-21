package controllers

import (
	"boolean-as-a-service/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/google/uuid"
)

var (
	createBooleanFunc func(boolean *models.Boolean) error
	getBooleanFunc    func(boolean *models.Boolean, id string) error
	updateBooleanFunc func(boolean *models.Boolean) error
	deleteBooleanFunc func(boolean *models.Boolean, id string) error
)

type booleanRepoMock struct{}

func (br *booleanRepoMock) CreateBoolean(boolean *models.Boolean) error {
	return createBooleanFunc(boolean)
}
func (br *booleanRepoMock) GetBooleanByID(boolean *models.Boolean, id string) error {
	return getBooleanFunc(boolean, id)
}

func (br *booleanRepoMock) UpdateBoolean(boolean *models.Boolean) error {
	return updateBooleanFunc(boolean)
}

func (br *booleanRepoMock) DeleteBoolean(boolean *models.Boolean, id string) error {
	return deleteBooleanFunc(boolean, id)
}

type booleanRequest struct {
	Value *bool  `json:"value"`
	Key   string `json:"key"`
}

func TestCreateBooleanBadRequest(t *testing.T) {
	testRequests := []booleanRequest{
		{
			Key: "This is test1",
		},
		{},
	}

	models.BooleanRepo = &booleanRepoMock{}
	for _, testRequest := range testRequests {
		createBooleanFunc = func(boolean *models.Boolean) error {
			boolean.ID = uuid.New().String()
			boolean.Value = testRequest.Value
			boolean.Key = testRequest.Key
			return nil
		}

		jsonRequest, err := json.Marshal(booleanRequest{
			Value: testRequest.Value,
			Key:   testRequest.Key,
		})
		assert.Nil(t, err)

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(http.MethodGet, "", bytes.NewBuffer(jsonRequest))

		CreateBoolean(c)
		assert.Equal(t, response.Code, http.StatusBadRequest)

	}
}

func TestCreateBooleanStatusOk(t *testing.T) {
	values := []bool{true, false}
	testRequests := []booleanRequest{
		{
			Value: &values[0],
			Key:   "This is test1",
		},
		{
			Value: &values[1],
		},
	}

	models.BooleanRepo = &booleanRepoMock{}
	for _, testRequest := range testRequests {
		createBooleanFunc = func(boolean *models.Boolean) error {
			boolean.ID = uuid.New().String()
			boolean.Value = testRequest.Value
			boolean.Key = testRequest.Key
			return nil
		}

		jsonRequest, err := json.Marshal(booleanRequest{
			Value: testRequest.Value,
			Key:   testRequest.Key,
		})
		assert.Nil(t, err)
		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(http.MethodGet, "", bytes.NewBuffer(jsonRequest))

		CreateBoolean(c)
		assert.Equal(t, response.Code, http.StatusOK)

		var boolean models.Boolean
		err = json.Unmarshal(response.Body.Bytes(), &boolean)
		assert.Nil(t, err)
		assert.NotEqual(t, boolean.ID, "")
		assert.Equal(t, boolean.Value, testRequest.Value)
		assert.Equal(t, boolean.Key, testRequest.Key)

	}
}

func TestGetBooleanStatusNotFound(t *testing.T) {
	testRequests := []string{"abcfsh-aghahasj-anjjns-bnadnakd"}

	models.BooleanRepo = &booleanRepoMock{}
	for _, testRequest := range testRequests {
		getBooleanFunc = func(boolean *models.Boolean, id string) error {
			return fmt.Errorf("id doesn't exists")
		}

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
		c.Params = gin.Params{
			{Key: "id", Value: testRequest},
		}

		GetBoolean(c)
		assert.Equal(t, response.Code, http.StatusNotFound)
	}
}

func TestGetBooleanStatusOk(t *testing.T) {
	values := []bool{true, false}
	testRequests := []models.Boolean{
		{
			ID:    uuid.New().String(),
			Value: &values[0],
			Key:   "test 1",
		},
		{
			ID:    uuid.New().String(),
			Value: &values[1],
		},
	}

	models.BooleanRepo = &booleanRepoMock{}
	for _, testRequest := range testRequests {
		getBooleanFunc = func(boolean *models.Boolean, id string) error {
			boolean.ID = testRequest.ID
			boolean.Value = testRequest.Value
			boolean.Key = testRequest.Key
			return nil
		}

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
		c.Params = gin.Params{
			{Key: "id", Value: testRequest.ID},
		}

		GetBoolean(c)
		assert.Equal(t, response.Code, http.StatusOK)

		var boolean models.Boolean
		err := json.Unmarshal(response.Body.Bytes(), &boolean)
		assert.Nil(t, err)
		assert.Equal(t, boolean.ID, testRequest.ID)
		assert.Equal(t, boolean.Value, testRequest.Value)
		assert.Equal(t, boolean.Key, testRequest.Key)

	}
}

func TestUpdateBooleanStatusBadRequest(t *testing.T) {
	testRequests := []models.Boolean{
		{
			ID:  uuid.New().String(),
			Key: "This is test1",
		},
		{
			ID: uuid.New().String(),
		},
	}

	models.BooleanRepo = &booleanRepoMock{}
	for _, testRequest := range testRequests {

		getBooleanFunc = func(boolean *models.Boolean, id string) error {
			return nil
		}

		updateBooleanFunc = func(boolean *models.Boolean) error {
			return nil
		}

		jsonRequest, err := json.Marshal(booleanRequest{
			Value: testRequest.Value,
			Key:   testRequest.Key,
		})
		assert.Nil(t, err)

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(http.MethodGet, "", bytes.NewBuffer(jsonRequest))
		c.Params = gin.Params{
			{Key: "id", Value: testRequest.ID},
		}

		UpdateBoolean(c)
		assert.Equal(t, response.Code, http.StatusBadRequest)

	}
}

func TestUpdateBooleanStatusNotFound(t *testing.T) {
	values := []bool{true, false}
	testRequests := []models.Boolean{
		{
			ID:    uuid.New().String(),
			Value: &values[0],
			Key:   "This is test1",
		},
		{
			ID:    uuid.New().String(),
			Value: &values[1],
		},
	}

	models.BooleanRepo = &booleanRepoMock{}
	for _, testRequest := range testRequests {

		getBooleanFunc = func(boolean *models.Boolean, id string) error {
			return nil
		}

		updateBooleanFunc = func(boolean *models.Boolean) error {
			return nil
		}

		jsonRequest, err := json.Marshal(booleanRequest{
			Value: testRequest.Value,
			Key:   testRequest.Key,
		})
		assert.Nil(t, err)

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(http.MethodGet, "", bytes.NewBuffer(jsonRequest))
		c.Params = gin.Params{
			{Key: "id", Value: testRequest.ID},
		}

		UpdateBoolean(c)
		assert.Equal(t, response.Code, http.StatusNotFound)

	}
}

func TestUpdateBooleanStatusOK(t *testing.T) {
	values := []bool{true, false}
	testRequests := []models.Boolean{
		{
			ID:    uuid.New().String(),
			Value: &values[0],
			Key:   "This is test1",
		},
		{
			ID:    uuid.New().String(),
			Value: &values[1],
		},
	}

	models.BooleanRepo = &booleanRepoMock{}
	for _, testRequest := range testRequests {

		getBooleanFunc = func(boolean *models.Boolean, id string) error {
			val := true
			boolean.ID = testRequest.ID
			boolean.Value = &val
			boolean.Key = "This is key"
			return nil
		}

		updateBooleanFunc = func(boolean *models.Boolean) error {
			boolean.ID = testRequest.ID
			boolean.Value = testRequest.Value
			if strings.Compare(testRequest.Key, "") == 0 {
				boolean.Key = "This is key"
			} else {
				boolean.Key = testRequest.Key
			}
			return nil
		}

		jsonRequest, err := json.Marshal(booleanRequest{
			Value: testRequest.Value,
			Key:   testRequest.Key,
		})
		assert.Nil(t, err)

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(http.MethodGet, "", bytes.NewBuffer(jsonRequest))
		c.Params = gin.Params{
			{Key: "id", Value: testRequest.ID},
		}

		UpdateBoolean(c)
		assert.Equal(t, response.Code, http.StatusOK)

		var boolean models.Boolean
		err = json.Unmarshal(response.Body.Bytes(), &boolean)
		assert.Nil(t, err)
		assert.Equal(t, boolean.ID, testRequest.ID)
		assert.Equal(t, boolean.Value, testRequest.Value)
		if strings.Compare(testRequest.Key, "") == 0 {
			assert.Equal(t, boolean.Key, "This is key")
		} else {
			assert.Equal(t, boolean.Key, testRequest.Key)
		}

	}
}

func TestDeleteBooleanStatusNotFound(t *testing.T) {
	testRequests := []string{"abcfsh-aghahasj-anjjns-bnadnakd"}

	models.BooleanRepo = &booleanRepoMock{}
	for _, testRequest := range testRequests {

		getBooleanFunc = func(boolean *models.Boolean, id string) error {
			return fmt.Errorf("id doesn't exists")
		}

		deleteBooleanFunc = func(boolean *models.Boolean, id string) error {
			return nil
		}

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
		c.Params = gin.Params{
			{Key: "id", Value: testRequest},
		}

		DeleteBoolean(c)
		assert.Equal(t, response.Code, http.StatusNotFound)
	}
}

func TestDeleteBooleanStatusNoContent(t *testing.T) {
	values := []bool{true, false}
	testRequests := []models.Boolean{
		{
			ID:    uuid.New().String(),
			Value: &values[0],
			Key:   "test 1",
		},
		{
			ID:    uuid.New().String(),
			Value: &values[1],
		},
	}

	models.BooleanRepo = &booleanRepoMock{}
	for _, testRequest := range testRequests {
		getBooleanFunc = func(boolean *models.Boolean, id string) error {
			boolean.ID = testRequest.ID
			boolean.Value = testRequest.Value
			boolean.Key = testRequest.Key
			return nil
		}

		deleteBooleanFunc = func(boolean *models.Boolean, id string) error {
			return nil
		}

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		c.Request, _ = http.NewRequest(http.MethodGet, "", nil)
		c.Params = gin.Params{
			{Key: "id", Value: testRequest.ID},
		}

		DeleteBoolean(c)
		assert.Equal(t, response.Code, http.StatusNoContent)
	}
}
