package controllers

import (
	"boolean-as-a-service/models"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

//CreateBoolean  create a boolean
func CreateBoolean(c *gin.Context) {
	var boolean models.Boolean
	booleanRepo := models.BooleanRepo

	if err := c.BindJSON(&boolean); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	boolean.ID = uuid.New().String()
	if err := booleanRepo.CreateBoolean(&boolean); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, boolean)
}

//GetBoolean  get a boolean
func GetBoolean(c *gin.Context) {
	var boolean models.Boolean
	booleanRepo := models.BooleanRepo

	id := c.Params.ByName("id")
	if err := booleanRepo.GetBooleanByID(&boolean, id); err != nil {
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
	var boolean models.Boolean
	booleanRepo := models.BooleanRepo

	id := c.Params.ByName("id")
	if err := c.BindJSON(&boolean); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var booln models.Boolean
	if err := booleanRepo.GetBooleanByID(&booln, id); err != nil {
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

	if err := booleanRepo.UpdateBoolean(&boolean); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, boolean)

}

//DeleteBoolean ... delete a boolean
func DeleteBoolean(c *gin.Context) {
	var boolean models.Boolean
	booleanRepo := models.BooleanRepo

	id := c.Params.ByName("id")
	if err := booleanRepo.GetBooleanByID(&boolean, id); err != nil {
		if err.Error() == "record not found" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	boolean = models.Boolean{}
	if err := booleanRepo.DeleteBooleanByID(&boolean, id); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(http.StatusNoContent)
}
