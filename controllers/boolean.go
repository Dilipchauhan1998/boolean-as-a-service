package controllers

import (
	"boolean-as-a-service/auth"
	"boolean-as-a-service/models"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

//CreateToken ...
func CreateToken(c *gin.Context) {
	token := auth.Token{
		ID: uuid.New().String(),
	}

	if err := auth.TokenRepo.CreateToken(token); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token.ID,
	})

}

//CreateBoolean  create a boolean
func CreateBoolean(c *gin.Context) {
	var boolean models.Boolean
	if value, _ := c.Request.Header["Authorization"]; len(value) == 0 || !auth.TokenRepo.ExistToken(strings.Split(value[0], " ")[1]) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err := c.BindJSON(&boolean); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	boolean.ID = uuid.New().String()
	if err := models.BooleanRepo.CreateBoolean(boolean); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, boolean)
}

//GetBoolean  get a boolean
func GetBoolean(c *gin.Context) {
	var (
		boolean models.Boolean
		err     error
	)
	if value, _ := c.Request.Header["Authorization"]; len(value) == 0 || !auth.TokenRepo.ExistToken(strings.Split(value[0], " ")[1]) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	id := c.Params.ByName("id")
	if boolean, err = models.BooleanRepo.GetBooleanByID(id); err != nil {
		if err.Error() == "record not found" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, boolean)

}

//UpdateBoolean  update a boolean
func UpdateBoolean(c *gin.Context) {
	var (
		boolean models.Boolean
		err     error
	)

	if value, _ := c.Request.Header["Authorization"]; len(value) == 0 || !auth.TokenRepo.ExistToken(strings.Split(value[0], " ")[1]) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	id := c.Params.ByName("id")
	if err := c.BindJSON(&boolean); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var booln models.Boolean
	if booln, err = models.BooleanRepo.GetBooleanByID(id); err != nil {
		if err.Error() == "record not found" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	boolean.ID = id
	if strings.Compare(boolean.Key, "") == 0 {
		boolean.Key = booln.Key
	}

	if err := models.BooleanRepo.UpdateBoolean(boolean); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, boolean)

}

//DeleteBoolean ... delete a boolean
func DeleteBoolean(c *gin.Context) {
	if value, _ := c.Request.Header["Authorization"]; len(value) == 0 || !auth.TokenRepo.ExistToken(strings.Split(value[0], " ")[1]) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	id := c.Params.ByName("id")
	if _, err := models.BooleanRepo.GetBooleanByID(id); err != nil {
		if err.Error() == "record not found" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if err := models.BooleanRepo.DeleteBooleanByID(id); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
