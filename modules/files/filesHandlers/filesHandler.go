package filesHandlers

import (
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/muhammadfarhankt/nft-marketplace/config"
	"github.com/muhammadfarhankt/nft-marketplace/pkg/utils"

	"github.com/muhammadfarhankt/nft-marketplace/modules/entities"
	"github.com/muhammadfarhankt/nft-marketplace/modules/files"
	filesUsecases "github.com/muhammadfarhankt/nft-marketplace/modules/files/fileUsecases"
)

type filesHandlersErrCode string

const (
	uploadToGCPErr   filesHandlersErrCode = "files-001"
	deleteFromGCPErr filesHandlersErrCode = "files-002"
)

type IFilesHandler interface {
	UploadToGCP(c *fiber.Ctx) error
	DeleteFromGCP(c *fiber.Ctx) error
}

type filesHandler struct {
	cfg          config.IConfig
	filesUsecase filesUsecases.IFilesUsecase
}

func FilesHandler(cfg config.IConfig, filesUsecase filesUsecases.IFilesUsecase) IFilesHandler {
	return &filesHandler{
		cfg:          cfg,
		filesUsecase: filesUsecase,
	}
}

func (f *filesHandler) UploadToGCP(c *fiber.Ctx) error {
	req := make([]*files.FileReq, 0)
	form, err := c.MultipartForm()
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(uploadToGCPErr),
			err.Error(),
		).Res()
	}

	filesReq := form.File["files"]
	destination := c.FormValue("destination")

	// files  extension validaton
	imageExtensions := map[string]bool{
		"jpg":  true,
		"jpeg": true,
		"png":  true,
	}
	for _, file := range filesReq {
		extension := strings.TrimPrefix(filepath.Ext(file.Filename), ".")
		if _, ok := imageExtensions[extension]; !ok {
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(uploadToGCPErr),
				"invalid file extension",
			).Res()
		}

		if file.Size > int64(f.cfg.App().FileLimit()) {
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(uploadToGCPErr),
				"file size too large",
			).Res()
		}

		filename := utils.RandFileName(extension)
		req = append(req, &files.FileReq{
			File:        file,
			Destination: destination + "/" + filename,
			FileName:    filename,
			Extension:   extension,
		},
		)
	}

	res, err := f.filesUsecase.UploadToGCP(req)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(uploadToGCPErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(
		fiber.StatusCreated,
		res,
	).Res()
}

func (f *filesHandler) DeleteFromGCP(c *fiber.Ctx) error {

	req := make([]*files.DeleteFileReq, 0)

	if err := c.BodyParser(&req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(deleteFromGCPErr),
			err.Error(),
		).Res()
	}

	if err := f.filesUsecase.DeleteFileFromGCP(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(deleteFromGCPErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		"File Deleted Successfully",
	).Res()
}
