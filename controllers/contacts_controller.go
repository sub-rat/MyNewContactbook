package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sub-rat/MyNewContactbook/models"
)

func GetAllContacts(c *gin.Context) {
	first_name := c.Query("first_name")
	last_name := c.Query("last_name")
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	sizeString := c.Query("size")
	limit, err := strconv.Atoi(sizeString)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	// TODO fetch contact list from database where first_name = "ram"
	// select * from contacts where first_name = {first_name} and last_name = {last_name}
	fmt.Println(first_name, last_name)
	var contactList []models.Contact

	filterString := "%" + first_name + "%"
	fmt.Println(filterString)
	query := models.DB.Debug().Model(&models.Contact{})
	if filterString != "" {
		query.Where("first_name like ? ", filterString)
	}
	query.Limit(limit).Offset(limit * page)
	err = query.Find(&contactList).Error
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

	err := models.DB.Model(&models.Contact{}).Create(&contact).Error

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

func DeleteContacts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Delete contact comming soon",
	})
}

func GetContactById(c *gin.Context) {
	id := c.Params.ByName("id")
	// Atoi
	// Query in database to fetch contact by id
	// select * from contacts where id = {id}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get Contact By Id ",
		"data":    id,
	})
}
