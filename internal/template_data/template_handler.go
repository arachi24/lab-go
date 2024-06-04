package template

import (
	"example/model"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
)

type TemplateHandler interface {
	Create(ctx *fiber.Ctx) error
	FindOne(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
}

type templateHandler struct {
	templateRepo TemplateHandlerRepository
}

func NewTemplateHandler(templateRepo TemplateHandlerRepository) templateHandler {
	return templateHandler{templateRepo}
}

func (h *templateHandler) Create(ctx *fiber.Ctx) error {
	template := new(model.Template)
	if err := ctx.BodyParser(template); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(template); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.templateRepo.Create(template)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)
}

func (h *templateHandler) FindOne(ctx *fiber.Ctx) error {
	filters := new(ThemeQuery)

	err := ctx.QueryParser(filters)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.templateRepo.FindOne(filters)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)

}

func (h *templateHandler) Update(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	values := new(model.Template)

	if err := ctx.BodyParser(values); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(values); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	values.ID = uint(id)

	response, err := h.templateRepo.Update(id, values)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)
}
