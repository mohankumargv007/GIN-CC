package Controllers

import (
	"net/http"

	"alexedwards.net/CC-GO/Models"
	"github.com/gin-gonic/gin"
)

type CategoryModelValidator struct {
	CatName string `json:"cat_name" validate:"required"`
	CatDesc string `json:"cat_desc" validate:"required"`
}

func CreateCategory(context *gin.Context) {
	var category Models.Catergory
	CategoryModelValidator := CategoryModelValidator{}
	if err := context.ShouldBindJSON(&CategoryModelValidator); err == nil {
		//	c.BindJSON(&category)
		err := Models.CreateACategory(&category)
		if err != nil {
			context.AbortWithStatus(http.StatusNotFound)
		} else {
			context.JSON(http.StatusOK, category)
		}
	} else {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func GetAllCategories(c *gin.Context) {
	var category []Models.Catergory
	err := Models.GetAllCategories(&category)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, category)
	}
}

func GetACategory(c *gin.Context) {
	id := c.Params.ByName("id")
	var category Models.Catergory
	err := Models.GetACategory(&category, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, category)
	}
}

func UpdateACategory(c *gin.Context) {
	var category Models.Catergory
	id := c.Params.ByName("id")
	err := Models.GetACategory(&category, id)
	if err != nil {
		c.JSON(http.StatusNotFound, category)
	}
	c.BindJSON(&category)
	err = Models.UpdateACategory(&category, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, category)
	}
}

func DeleteACategory(c *gin.Context) {
	var category Models.Catergory
	id := c.Params.ByName("id")
	err := Models.DeleteACategory(&category, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id:" + id: "deleted"})
	}
}
