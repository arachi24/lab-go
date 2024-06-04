package business

import (
	"errors"
	"example/model"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type BusinessHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindOne(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}

type businessHandler struct {
	businessRepo BusinessRepository
}

func NewBusinessHandler(businessRepo BusinessRepository) businessHandler {
	return businessHandler{businessRepo}
}

func (h *businessHandler) Create(ctx *fiber.Ctx) error {
	bu := new(model.Business)
	if err := ctx.BodyParser(bu); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(bu); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.businessRepo.Create(bu)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)
}

func (h *businessHandler) FindAll(ctx *fiber.Ctx) error {
	response, err := h.businessRepo.FindAll()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)

}

func (h *businessHandler) FindOne(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.businessRepo.FindOne(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)

}

func (h *businessHandler) Update(ctx *fiber.Ctx) error {
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

	buModel := model.Business{
		Logo:        values.Logo,
		Image:       values.Image,
		Link:        values.Link,
		Title:       values.Title,
		Description: values.Description,
	}

	response, err := h.businessRepo.Update(id, &buModel)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)
}

func (h *businessHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.businessRepo.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"status": response,
	})
}
