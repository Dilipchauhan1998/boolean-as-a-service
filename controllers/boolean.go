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

//CreateBoolean  insert a boolean
func CreateBoolean(c *gin.Context) {
	var boolean models.Boolean
	err := c.BindJSON(&boolean)

	var booleanRepo models.BooleanRepoInterface
	booleanRepo = models.NewBooleanRepo(conn.DB)

	//fmt.Println("c:", c)
	//fmt.Println("boolean:", boolean)

	if err != nil {
		c.JSON(http.StatusOK, "Failed to bind the json")
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
	id := c.Params.ByName("id")
	//booleanRepo := models.NewBooleanRepo(conn.DB)
	var booleanRepo models.BooleanRepoInterface
	booleanRepo = models.NewBooleanRepo(conn.DB)

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
	id := c.Params.ByName("id")
	//booleanRepo := models.NewBooleanRepo(conn.DB)
	var booleanRepo models.BooleanRepoInterface
	booleanRepo = models.NewBooleanRepo(conn.DB)

	err := c.BindJSON(&boolean)

	if err != nil {
		c.JSON(http.StatusOK, "Failed to bind the json")
	} else {
		var booln models.Boolean
		err := booleanRepo.GetBooleanByID(&booln, id)

		//fmt.Println("booln:", booln)
		//fmt.Println("err:", err)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		var nilBoolean models.Boolean
		if booln == nilBoolean {
			//fmt.Println("nilBool")
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		boolean.ID = id
		if strings.Compare(boolean.Key, "") == 0 {
			//fmt.Println("empty Key")
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
	//booleanRepo := models.NewBooleanRepo(conn.DB)
	var booleanRepo models.BooleanRepoInterface
	booleanRepo = models.NewBooleanRepo(conn.DB)

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

// //CreateBoolean  insert a boolean
// func CreateBoolean(c *gin.Context) {
// 	var boolean models.Boolean
// 	err := c.BindJSON(&boolean)

// 	fmt.Println("c:", c)
// 	fmt.Println("boolean:", boolean)

// 	if err != nil {
// 		c.JSON(http.StatusOK, "Failed to bind the json")
// 	} else {
// 		boolean.ID = uuid.New().String()

// 		err := models.CreateBoolean(&boolean)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			c.AbortWithStatus(http.StatusNotFound)
// 		} else {
// 			c.JSON(http.StatusOK, boolean)
// 		}
// 	}
// }

// //GetBoolean  Get a boolean
// func GetBoolean(c *gin.Context) {
// 	var boolean models.Boolean
// 	id := c.Params.ByName("id")

// 	err := models.GetBooleanByID(&boolean, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.JSON(http.StatusOK, boolean)
// 	}
// }

// //UpdateBoolean  update a boolean
// func UpdateBoolean(c *gin.Context) {
// 	var boolean models.Boolean
// 	id := c.Params.ByName("id")
// 	err := c.BindJSON(&boolean)

// 	if err != nil {
// 		c.JSON(http.StatusOK, "Failed to bind the json")
// 	} else {

// 		err := models.UpdateBoolean(&boolean, id)
// 		if err != nil {
// 			c.AbortWithStatus(http.StatusNotFound)
// 		} else {
// 			c.JSON(http.StatusOK, boolean)
// 		}
// 	}

// }

// //DeleteBoolean ... delete a boolean
// func DeleteBoolean(c *gin.Context) {
// 	var boolean models.Boolean
// 	id := c.Params.ByName("id")

// 	err := models.DeleteBoolean(&boolean, id)
// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	} else {
// 		c.AbortWithStatus(http.StatusNoContent)
// 	}
// }
