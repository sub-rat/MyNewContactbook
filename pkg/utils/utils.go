package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(c *gin.Context) (int, int, error) {
	pageStr := c.Query("page")
	page := 0
	if pageStr != "" {
		p, err := strconv.Atoi(pageStr)
		if err != nil {
			return 0, 0, err
		}
		page = p - 1
		if page <= 0 {
			page = 0
		}
	}

	sizeString := c.Query("size")
	limit := 20
	if sizeString != "" {
		l, err := strconv.Atoi(sizeString)
		if err != nil {
			return 0, 0, err
		}
		limit = l
		if limit <= 0 {
			limit = 20
		}
	}
	return page, limit, nil
}
