package user

import (
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
	r.POST("/login", resource.GetLogin)
	r.GET("/users", middleware.CheckToken, resource.Query)
	r.POST("/users", resource.Create)
	r.PUT("/users/:id", middleware.CheckToken, resource.Update)
	r.GET("/users/:id", middleware.CheckToken, resource.Get)
	r.DELETE("/users/:id", middleware.CheckToken, resource.Delete)
}

func (resource *resource) Query(c *gin.Context) {
	page, limit, err := utils.Pagination(c)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	userList, err := resource.service.Query(page*limit, limit, c.Query("q"), "")
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "getting all users",
		"data":    userList,
	})
}

func (resource *resource) Create(c *gin.Context) {
	user := User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := resource.service.Create(&user)
	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Create user Successfully",
		"data":    user,
	})
}

func (resource *resource) Update(c *gin.Context) {
	user := User{}
	id, _ := strconv.Atoi(c.Params.ByName("id"))
	_, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Record not found",
		})
		return
	}

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	cont, err := resource.service.Update(uint(id), &user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Update the user",
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
	user, err := resource.service.Get(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Get User By Id ",
		"data":    user,
	})
}

func (resource *resource) GetLogin(c *gin.Context) {
	user := User{}
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	q := user.Email
	qType := "email"
	if user.UserName != "" {
		q = user.UserName
		qType = "user_name"
	}
	userList, err := resource.service.Query(0, 1, q, qType)
	if len(userList) <= 0 || err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User Not Found",
		})
		return
	}
	dbUser := userList[0]
	if !utils.CheckPassword(user.Password, dbUser.Password) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password Mismatch",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": utils.GenerateJwtToken(dbUser.ID),
		"user":  dbUser,
	})

}
