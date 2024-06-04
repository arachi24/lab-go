package query

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func QueryPaginate(ctx *fiber.Ctx) (page int, pageSize int, sort string, err error) {

	queryPage := ctx.Query("page", "1")
	queryPageSize := ctx.Query("page_size", "10")
	sort = ctx.Query("sort", "desc")
	page, _ = strconv.Atoi(queryPage)
	pageSize, _ = strconv.Atoi(queryPageSize)

	if page < 1 {
		return 0, 0, "", errors.New("page must be more than zero")
	}

	return
}
