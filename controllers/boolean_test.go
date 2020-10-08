package controllers

import (
	"boolean-as-a-service/auth"
	"boolean-as-a-service/mocks"
	"boolean-as-a-service/models"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

type booleanRequest struct {
	Value *bool  `json:"value"`
	Key   string `json:"key"`
}

func TestCreateTokenInternalServerError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().CreateToken(gomock.Any()).Return(errors.New("some kind of internal error"))

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodPost, "", nil)
	assert.Nil(t, err)
	c.Request = req

	CreateToken(c)
	assert.Equal(t, response.Code, http.StatusInternalServerError)

}

func TestCreateTokenStatusOK(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokenRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokenRepoInterface
	mocktokenRepoInterface.EXPECT().CreateToken(gomock.Any()).Return(nil)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodPost, "", nil)
	assert.Nil(t, err)
	c.Request = req

	CreateToken(c)
	assert.Equal(t, response.Code, http.StatusOK)

}

func TestCreateBooleanStatusUnauthorized(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(false)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodPost, "", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req

	CreateBoolean(c)
	assert.Equal(t, response.Code, http.StatusUnauthorized)

}

func TestCreateBooleanBadRequest(t *testing.T) {
	testRequests := []booleanRequest{
		{
			Key: "This is test1",
		},
		{},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true).Times(len(testRequests))

	for _, testRequest := range testRequests {

		jsonRequest, err := json.Marshal(booleanRequest{
			Value: testRequest.Value,
			Key:   testRequest.Key,
		})
		assert.Nil(t, err)

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(jsonRequest))
		assert.Nil(t, err)
		req.Header.Add("Authorization", "Token "+uuid.New().String())
		c.Request = req

		CreateBoolean(c)
		assert.Equal(t, response.Code, http.StatusBadRequest)
	}
}

func TestCreateBooleanStatusInternalServerError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true)

	ctl = gomock.NewController(t)
	defer ctl.Finish()
	mockbooleanRepoInterface := mocks.NewMockbooleanRepoInterface(ctl)
	models.BooleanRepo = mockbooleanRepoInterface
	mockbooleanRepoInterface.EXPECT().CreateBoolean(gomock.Any()).Return(errors.New("some kind of internal error"))

	jsonRequest, err := json.Marshal(booleanRequest{
		Value: new(bool),
		Key:   "some key",
	})
	assert.Nil(t, err)
	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(jsonRequest))
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req

	CreateBoolean(c)
	assert.Equal(t, response.Code, http.StatusInternalServerError)
}

func TestCreateBooleanStatusOK(t *testing.T) {
	testRequests := []booleanRequest{
		{
			Value: new(bool),
			Key:   "This is test1",
		},
		{
			Value: new(bool),
		},
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true).Times(len(testRequests))

	ctl = gomock.NewController(t)
	defer ctl.Finish()
	mockbooleanRepoInterface := mocks.NewMockbooleanRepoInterface(ctl)
	models.BooleanRepo = mockbooleanRepoInterface
	mockbooleanRepoInterface.EXPECT().CreateBoolean(gomock.Any()).Return(nil).Times(len(testRequests))

	for _, testRequest := range testRequests {

		jsonRequest, err := json.Marshal(booleanRequest{
			Value: testRequest.Value,
			Key:   testRequest.Key,
		})
		assert.Nil(t, err)
		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		req, err := http.NewRequest(http.MethodPost, "", bytes.NewBuffer(jsonRequest))
		assert.Nil(t, err)
		req.Header.Add("Authorization", "Token "+uuid.New().String())
		c.Request = req

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

func TestGetBooleanStatusUnauthorized(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(false)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodGet, "", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req

	GetBoolean(c)
	assert.Equal(t, response.Code, http.StatusUnauthorized)

}

func TestGetBooleanStatusNotFound(t *testing.T) {
	id := uuid.New().String()

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true)

	ctl = gomock.NewController(t)
	defer ctl.Finish()
	mockbooleanRepoInterface := mocks.NewMockbooleanRepoInterface(ctl)
	models.BooleanRepo = mockbooleanRepoInterface
	mockbooleanRepoInterface.EXPECT().GetBooleanByID(id).Return(models.Boolean{}, errors.New("record not found"))

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodGet, "", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: id},
	}

	GetBoolean(c)
	assert.Equal(t, response.Code, http.StatusNotFound)

}

func TestGetBooleanStatusInternalServerError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true)

	ctl = gomock.NewController(t)
	defer ctl.Finish()
	mockbooleanRepoInterface := mocks.NewMockbooleanRepoInterface(ctl)
	models.BooleanRepo = mockbooleanRepoInterface
	mockbooleanRepoInterface.EXPECT().GetBooleanByID(gomock.Any()).Return(models.Boolean{}, errors.New("some kind of internal error"))

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodGet, "", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: uuid.New().String()},
	}

	GetBoolean(c)
	assert.Equal(t, response.Code, http.StatusInternalServerError)

}

func TestGetBooleanStatusOK(t *testing.T) {
	boolean := models.Boolean{
		ID:    uuid.New().String(),
		Value: new(bool),
		Key:   "test 1",
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true)

	ctl = gomock.NewController(t)
	defer ctl.Finish()
	mockbooleanRepoInterface := mocks.NewMockbooleanRepoInterface(ctl)
	models.BooleanRepo = mockbooleanRepoInterface
	mockbooleanRepoInterface.EXPECT().GetBooleanByID(boolean.ID).Return(boolean, nil)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodGet, "", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: boolean.ID},
	}

	GetBoolean(c)
	assert.Equal(t, response.Code, http.StatusOK)

	var responseBoolean models.Boolean
	err = json.Unmarshal(response.Body.Bytes(), &responseBoolean)
	assert.Nil(t, err)
	assert.Equal(t, responseBoolean.ID, boolean.ID)
	assert.Equal(t, responseBoolean.Value, boolean.Value)
	assert.Equal(t, responseBoolean.Key, boolean.Key)

}

func TestUpdateBooleanStatusUnauthorized(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(false)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodPatch, "", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req

	UpdateBoolean(c)
	assert.Equal(t, response.Code, http.StatusUnauthorized)

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

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true).Times(len(testRequests))

	for _, testRequest := range testRequests {

		jsonRequest, err := json.Marshal(booleanRequest{
			Value: testRequest.Value,
			Key:   testRequest.Key,
		})
		assert.Nil(t, err)

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		req, err := http.NewRequest(http.MethodPatch, "", bytes.NewBuffer(jsonRequest))
		assert.Nil(t, err)
		req.Header.Add("Authorization", "Token "+uuid.New().String())
		c.Request = req
		c.Params = gin.Params{
			{Key: "id", Value: testRequest.ID},
		}

		UpdateBoolean(c)
		assert.Equal(t, response.Code, http.StatusBadRequest)

	}
}

func TestUpdateBooleanStatusNotFound(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true)

	ctl = gomock.NewController(t)
	defer ctl.Finish()
	mockbooleanRepoInterface := mocks.NewMockbooleanRepoInterface(ctl)
	models.BooleanRepo = mockbooleanRepoInterface
	mockbooleanRepoInterface.EXPECT().GetBooleanByID(gomock.Any()).Return(models.Boolean{}, errors.New("record not found"))

	jsonRequest, err := json.Marshal(booleanRequest{
		Value: new(bool),
		Key:   "some key",
	})
	assert.Nil(t, err)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodPatch, "", bytes.NewBuffer(jsonRequest))
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: uuid.New().String()},
	}

	UpdateBoolean(c)
	assert.Equal(t, response.Code, http.StatusNotFound)

}

func TestUpdateBooleanStatusInternalServerError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true).Times(2)

	ctl = gomock.NewController(t)
	defer ctl.Finish()
	mockbooleanRepoInterface := mocks.NewMockbooleanRepoInterface(ctl)
	models.BooleanRepo = mockbooleanRepoInterface
	models.BooleanRepo = mockbooleanRepoInterface
	gomock.InOrder(
		mockbooleanRepoInterface.EXPECT().GetBooleanByID(gomock.Any()).Return(models.Boolean{}, errors.New("some kind of internal error")),
		mockbooleanRepoInterface.EXPECT().GetBooleanByID(gomock.Any()).Return(models.Boolean{}, nil),
		mockbooleanRepoInterface.EXPECT().UpdateBoolean(gomock.Any()).Return(errors.New("some kind of internal error")),
	)

	jsonRequest, err := json.Marshal(booleanRequest{
		Value: new(bool),
		Key:   "some key",
	})
	assert.Nil(t, err)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodPatch, "", bytes.NewBuffer(jsonRequest))
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: uuid.New().String()},
	}
	UpdateBoolean(c)
	assert.Equal(t, response.Code, http.StatusInternalServerError)

	response = httptest.NewRecorder()
	c, _ = gin.CreateTestContext(response)
	req, err = http.NewRequest(http.MethodPatch, "", bytes.NewBuffer(jsonRequest))
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: uuid.New().String()},
	}
	UpdateBoolean(c)
	assert.Equal(t, response.Code, http.StatusInternalServerError)

}

func TestUpdateBooleanStatusOK(t *testing.T) {
	testRequests := []models.Boolean{
		{
			ID:    uuid.New().String(),
			Value: new(bool),
			Key:   "This is test1",
		},
		{
			ID:    uuid.New().String(),
			Value: new(bool),
		},
	}

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true).Times(len(testRequests))

	ctl = gomock.NewController(t)
	defer ctl.Finish()
	mockbooleanRepoInterface := mocks.NewMockbooleanRepoInterface(ctl)
	models.BooleanRepo = mockbooleanRepoInterface
	gomock.InOrder(
		mockbooleanRepoInterface.EXPECT().GetBooleanByID(testRequests[0].ID).Return(models.Boolean{ID: testRequests[0].ID, Value: testRequests[0].Value}, nil),
		mockbooleanRepoInterface.EXPECT().UpdateBoolean(testRequests[0]).Return(nil),

		mockbooleanRepoInterface.EXPECT().GetBooleanByID(testRequests[1].ID).Return(models.Boolean{ID: testRequests[1].ID, Value: testRequests[1].Value, Key: "some key"}, nil),
		mockbooleanRepoInterface.EXPECT().UpdateBoolean(models.Boolean{ID: testRequests[1].ID, Value: testRequests[1].Value, Key: "some key"}).Return(nil),
	)

	for i, testRequest := range testRequests {

		jsonRequest, err := json.Marshal(booleanRequest{
			Value: testRequest.Value,
			Key:   testRequest.Key,
		})
		assert.Nil(t, err)

		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)
		req, err := http.NewRequest(http.MethodPatch, "", bytes.NewBuffer(jsonRequest))
		assert.Nil(t, err)
		req.Header.Add("Authorization", "Token "+uuid.New().String())
		c.Request = req
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

		if i == 0 {
			assert.Equal(t, boolean.Key, "This is test1")
		} else if i == 1 {
			assert.Equal(t, boolean.Key, "some key")
		}

	}
}

func TestDeleteBooleanStatusUnauthorized(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(false)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodDelete, "", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req

	DeleteBoolean(c)
	assert.Equal(t, response.Code, http.StatusUnauthorized)

}

func TestDeleteBooleanStatusNotFound(t *testing.T) {
	id := uuid.New().String()

	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true)

	ctl = gomock.NewController(t)
	defer ctl.Finish()
	mockbooleanRepoInterface := mocks.NewMockbooleanRepoInterface(ctl)
	models.BooleanRepo = mockbooleanRepoInterface
	mockbooleanRepoInterface.EXPECT().GetBooleanByID(id).Return(models.Boolean{}, errors.New("record not found"))

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodDelete, "", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: id},
	}

	DeleteBoolean(c)
	assert.Equal(t, response.Code, http.StatusNotFound)

}

func TestDeleteBooleanStatusInternalServerError(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true).Times(2)

	ctl = gomock.NewController(t)
	defer ctl.Finish()
	mockbooleanRepoInterface := mocks.NewMockbooleanRepoInterface(ctl)
	models.BooleanRepo = mockbooleanRepoInterface
	gomock.InOrder(
		mockbooleanRepoInterface.EXPECT().GetBooleanByID(gomock.Any()).Return(models.Boolean{}, errors.New("some kind of internal error")),
		mockbooleanRepoInterface.EXPECT().GetBooleanByID(gomock.Any()).Return(models.Boolean{}, nil),
		mockbooleanRepoInterface.EXPECT().DeleteBooleanByID(gomock.Any()).Return(errors.New("some kind of internal error")),
	)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodDelete, "", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: uuid.New().String()},
	}

	DeleteBoolean(c)
	assert.Equal(t, response.Code, http.StatusInternalServerError)

	DeleteBoolean(c)
	assert.Equal(t, response.Code, http.StatusInternalServerError)
}

func TestDeleteBooleanStatusNoContent(t *testing.T) {
	boolean := models.Boolean{
		ID:    uuid.New().String(),
		Value: new(bool),
		Key:   "test 1",
	}
	ctl := gomock.NewController(t)
	defer ctl.Finish()
	mocktokeRepoInterface := mocks.NewMocktokenRepoInterface(ctl)
	auth.TokenRepo = mocktokeRepoInterface
	mocktokeRepoInterface.EXPECT().ExistToken(gomock.Any()).Return(true)

	ctl = gomock.NewController(t)
	defer ctl.Finish()
	mockbooleanRepoInterface := mocks.NewMockbooleanRepoInterface(ctl)
	models.BooleanRepo = mockbooleanRepoInterface
	gomock.InOrder(
		mockbooleanRepoInterface.EXPECT().GetBooleanByID(boolean.ID).Return(boolean, nil),
		mockbooleanRepoInterface.EXPECT().DeleteBooleanByID(boolean.ID).Return(nil),
	)

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	req, err := http.NewRequest(http.MethodDelete, "", nil)
	assert.Nil(t, err)
	req.Header.Add("Authorization", "Token "+uuid.New().String())
	c.Request = req
	c.Params = gin.Params{
		{Key: "id", Value: boolean.ID},
	}

	DeleteBoolean(c)
	assert.Equal(t, response.Code, http.StatusNoContent)
}
