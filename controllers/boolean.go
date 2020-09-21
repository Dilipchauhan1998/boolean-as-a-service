package controllers

import (
	"boolean-as-a-service/conn"
	"boolean-as-a-service/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func init() {
	models.NewBooleanRepo(conn.DB)
}

//CreateBoolean  insert a boolean
func CreateBoolean(c *gin.Context) {
	var boolean models.Boolean
	booleanRepo := models.BooleanRepo

	err := c.BindJSON(&boolean)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {

		boolean.ID = uuid.New().String()

		err := booleanRepo.CreateBoolean(&boolean)
		if err != nil {
			fmt.Println(err.Error())
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, boolean)
		}
	}
}

//GetBoolean  Get a boolean
func GetBoolean(c *gin.Context) {
	var boolean models.Boolean
	booleanRepo := models.BooleanRepo
	id := c.Params.ByName("id")
	fmt.Println("id:", id)
	err := booleanRepo.GetBooleanByID(&boolean, id)
	if err != nil {
		fmt.Println("err:", err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		fmt.Println("not err:", err)
		c.JSON(http.StatusOK, boolean)
	}
}

//UpdateBoolean  update a boolean
func UpdateBoolean(c *gin.Context) {
	var boolean models.Boolean
	booleanRepo := models.BooleanRepo
	id := c.Params.ByName("id")

	err := c.BindJSON(&boolean)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	} else {
		var booln models.Boolean
		err := booleanRepo.GetBooleanByID(&booln, id)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		var nilBoolean models.Boolean
		if booln == nilBoolean {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		boolean.ID = id
		if strings.Compare(boolean.Key, "") == 0 {
			boolean.Key = booln.Key
		}

		err = booleanRepo.UpdateBoolean(&boolean)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
		} else {
			c.JSON(http.StatusOK, boolean)
		}
	}

}

//DeleteBoolean ... delete a boolean
func DeleteBoolean(c *gin.Context) {
	var boolean models.Boolean
	booleanRepo := models.BooleanRepo

	id := c.Params.ByName("id")

	//check if the boolean exists with given id
	if err := booleanRepo.GetBooleanByID(&boolean, id); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	boolean = models.Boolean{}
	err := booleanRepo.DeleteBoolean(&boolean, id)

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.AbortWithStatus(http.StatusNoContent)
	}
}
