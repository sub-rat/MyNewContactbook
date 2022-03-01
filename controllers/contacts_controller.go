package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sub-rat/MyNewContactbook/models"
	"github.com/sub-rat/MyNewContactbook/utils"
)

func GetAllContacts(c *gin.Context) {
	first_name := c.Query("first_name")
	last_name := c.Query("last_name")

	fmt.Println(first_name, last_name)
	var contactList []models.Contact
	page, limit, err := utils.Pagination(c)
	fmt.Println(page, limit)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = models.DB.
		Debug().
		Model(&models.Contact{}).
		Where("first_name like ? ", "%"+first_name+"%").
		Limit(limit).
		Offset(limit * page).
		Find(&contactList).
		Error
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "getting all contacts",
		"data":    contactList,
	})
}

func CreateContact(c *gin.Context) {
	contact := models.Contact{}
	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := models.DB.
		Debug().
		Model(&models.Contact{}).
		Create(&contact).
		Error

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Create contact Successfully",
		"data":    contact,
	})
}

func UpdateContact(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Update contact comming soon",
	})
}

func DeleteContactsById(c *gin.Context) {
	id := c.Params.ByName("id")
	// delete recored with id
	err := models.DB.
		Debug().
		Delete(&models.Contact{}, id).
		Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{
		"message": "Successfully Deleted",
	})
}

func GetContactById(c *gin.Context) {
	contactId := c.Params.ByName("id")
	contact := models.Contact{}
	err := models.DB.Debug().Model(&models.Contact{}).First(&contact, contactId).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Contact By Id ",
		"data":    contact,
	})
}
