package upload

import (
	azureblob "example/utils/azure-blob"
	"io"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UploadHandler interface {
	Create(ctx *fiber.Ctx) error
}

type uploadHandler struct {
}

func NewUploadHandler() uploadHandler {
	return uploadHandler{}
}

func (h uploadHandler) Create(ctx *fiber.Ctx) error {
	files, err := ctx.Context().Request.MultipartForm()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}
	fileUpload := files.File["files"]

	if len(fileUpload) < 1 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}
	content, _ := fileUpload[0].Open()
	contentBytes, _ := io.ReadAll(content)
	req := UploadRequest{
		FileType:      fileUpload[0].Filename,
		ContentUpload: contentBytes,
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "errors": err.Error()})
	}

	newUpload, err := azureblob.Upload(req.ContentUpload, req.FileType)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Couldn't Upload ", "errors": err.Error()})
	}

	return ctx.JSON(&UploadResponse{
		Url: *newUpload,
	})
}
