package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sub-rat/MyNewContactbook/models"
)

func GetAllContacts(c *gin.Context) {
	first_name := c.Query("first_name")
	last_name := c.Query("last_name")
	// TODO fetch contact list from database where first_name = "ram"
	// select * from contacts where first_name = {first_name} and last_name = {last_name}
	fmt.Println(first_name, last_name)

	c.JSON(http.StatusOK, gin.H{
		"message": "getting all contacts",
		"data":    []models.Contact{
			// {
			// 	ID:        1,
			// 	FirstName: "Ram",
			// 	LastName:  "Sharma",
			// 	// Phone:     []models.Phone{},
			// },
			// {
			// 	ID:        2,
			// 	FirstName: "Hari",
			// 	LastName:  "Sharma",
			// Phone: []models.Phone{
			// 	{
			// 		PhoneType:   "Mobile",
			// 		PhoneNumber: "9090909090",
			// 	},
			// },
			// },
		},
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
	// TODO save contact to database
	contact.ID = 5

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
