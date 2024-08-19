package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(c *gin.Context) (limit int, page int, err error) {

	limitQr := c.Query("limit")

	if limitQr == "" {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitQr)
		if err != nil {
			return 0, 0, err
		}
	}

	pageQr := c.Query("page")

	if pageQr == "" {
		page = 1
	} else {
		page, err = strconv.Atoi(pageQr)
		if err != nil {
			return 0, 0, err
		}
	}

	return limit, page, nil
}

func PaginationNull(c *gin.Context) (limit int, page int, err error) {

	limitQr := c.Query("limit")

	if limitQr == "" {
		limit = 0
	} else {
		limit, err = strconv.Atoi(limitQr)
		if err != nil {
			return 0, 0, err
		}
	}

	pageQr := c.Query("page")

	if pageQr == "" {
		page = 0
	} else {
		page, err = strconv.Atoi(pageQr)
		if err != nil {
			return 0, 0, err
		}
	}

	return limit, page, nil
}
