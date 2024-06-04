package product

import (
	"errors"
	"example/model"
	"example/utils/query"
	"math"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindOne(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}

type productHandler struct {
	productRepo ProductRepository
}

func NewProductHandler(productRepo ProductRepository) productHandler {
	return productHandler{productRepo}
}

func (h *productHandler) Create(ctx *fiber.Ctx) error {
	values := new(model.Product)
	if err := ctx.BodyParser(values); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(values); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.productRepo.Create(values)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)
}

func (h *productHandler) FindAll(ctx *fiber.Ctx) error {
	page, pageSize, sort, err := query.QueryPaginate(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	offset := (page - 1) * pageSize

	response, totalPage, err := h.productRepo.FindAll(offset, pageSize, sort)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	var result ProductResp

	pageTotals := math.Ceil(float64(totalPage) / float64(pageSize))

	result = ProductResp{
		Records:   response,
		Page:      page,
		PageSize:  pageSize,
		TotalPage: int(pageTotals),
	}

	return ctx.JSON(result)

}

func (h *productHandler) FindOne(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.productRepo.FindOne(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)

}

func (h *productHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.productRepo.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"status": response,
	})
}

func (h *productHandler) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	values := new(model.Product)

	if err := ctx.BodyParser(values); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(values); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}
	values.ID = uint(id)

	response, err := h.productRepo.Update(id, values)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)
}
