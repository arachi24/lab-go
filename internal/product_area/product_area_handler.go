package productarea

import (
	"errors"
	"example/model"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductareaHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindOne(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}

type productareaHandler struct {
	productareaRepo ProductareaRepository
}

func NewProductareaHandler(productareaRepo ProductareaRepository) productareaHandler {
	return productareaHandler{productareaRepo}
}

func (h *productareaHandler) Create(ctx *fiber.Ctx) error {
	homepage := new(model.ProductArea)
	if err := ctx.BodyParser(homepage); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(homepage); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.productareaRepo.Create(homepage)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)
}

func (h *productareaHandler) FindAll(ctx *fiber.Ctx) error {
	response, err := h.productareaRepo.FindAll()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)

}

func (h *productareaHandler) FindOne(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.productareaRepo.FindOne(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)

}

func (h *productareaHandler) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	values := new(UpdateReq)

	if err := ctx.BodyParser(values); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(values); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	homepageModel := model.ProductArea{
		Banner:      values.Banner,
		Logo:        values.Logo,
		Image:       values.Image,
		Link:        values.Link,
		Title:       values.Title,
		Description: values.Description,
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.productareaRepo.Update(id, &homepageModel)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)
}

func (h *productareaHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.productareaRepo.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"status": response,
	})
}
