package contact

import (
	"fmt"
	"github.com/sub-rat/MyNewContactbook/internals/middleware"
	"github.com/sub-rat/MyNewContactbook/pkg/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type resource struct {
	service ServiceInterface
}

func RegisterRoutes(r *gin.Engine, service ServiceInterface) {
	resource := &resource{service}
	r.GET("users/:id/contacts", middleware.CheckToken, resource.Query)
	r.POST("users/:id/contacts", middleware.CheckToken, resource.Create)
	r.PUT("/contacts/:id", middleware.CheckToken, resource.Update)
	r.GET("/contacts/:id", middleware.CheckToken, resource.Get)
	r.DELETE("/contacts/:id", middleware.CheckToken, resource.Delete)
}

func (resource *resource) Query(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Params.ByName("id"))
	page, limit, err := utils.Pagination(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	fmt.Println(userId)
	contactList, err := resource.service.Query(page*limit, limit, c.Query("first_name"), userId)
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

func (resource *resource) Create(c *gin.Context) {
	contact := Contact{}
	userId, _ := strconv.Atoi(c.Params.ByName("id"))
	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	//validation for userId exists or not
	contact.UserID = uint(userId)
	contact, err := resource.service.Create(&contact)
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

func (resource *resource) Update(c *gin.Context) {
	contact := Contact{}
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	_, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	if err := c.BindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cont, err := resource.service.Update(uint(id), &contact)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update the contact",
		"data":    cont,
	})
}

func (resource *resource) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	_, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	err = resource.service.Delete(uint(id))
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

func (resource *resource) Get(c *gin.Context) {
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	contact, err := resource.service.Get(uint(id))
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
