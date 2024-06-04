package homepage

import (
	"errors"
	"example/internal/banner"
	"example/model"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HomepageHandler interface {
	Create(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
	FindOne(ctx *fiber.Ctx) error
	FindAll(ctx *fiber.Ctx) error
}

type homepageHandler struct {
	homePageRepo HomepageRepository
	bannerRepo   banner.BannerRepository
}

func NewHomepageHandler(homePageRepo HomepageRepository, bannerRepo banner.BannerRepository) homepageHandler {
	return homepageHandler{homePageRepo, bannerRepo}
}

func (h *homepageHandler) Create(ctx *fiber.Ctx) error {
	homepage := new(model.Homepage)
	if err := ctx.BodyParser(homepage); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(homepage); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.homePageRepo.Create(homepage)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)
}

func (h *homepageHandler) FindAll(ctx *fiber.Ctx) error {
	response, err := h.homePageRepo.FindAll()

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)

}

func (h *homepageHandler) FindOne(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.homePageRepo.FindOne(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(response)

}

func (h *homepageHandler) Update(ctx *fiber.Ctx) error {
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

	homepageModel := model.Homepage{
		Title:       values.Title,
		Description: values.Description,
		SubTitle:    values.SubTitle,
	}

	mainBanner, err := h.bannerRepo.Update(int(values.MainBanner.ID), values.MainBanner)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	subBanner, err := h.bannerRepo.Update(int(values.SubBanner.ID), values.SubBanner)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.homePageRepo.Update(id, &homepageModel)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response.MainBanner = mainBanner
	response.SubBanner = subBanner

	return ctx.JSON(response)
}

func (h *homepageHandler) Delete(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	response, err := h.homePageRepo.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	return ctx.JSON(fiber.Map{
		"status": response,
	})
}
