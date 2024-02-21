package appinfoHandlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"

	"github.com/muhammadfarhankt/nft-marketplace/config"

	"github.com/muhammadfarhankt/nft-marketplace/modules/appinfo"
	"github.com/muhammadfarhankt/nft-marketplace/modules/appinfo/appinfoUsecases"
	"github.com/muhammadfarhankt/nft-marketplace/modules/entities"

	"github.com/muhammadfarhankt/nft-marketplace/pkg/nftauth"
)

type appinfoHandlersErr string

const (
	generateApiKeyErr appinfoHandlersErr = "appinfo-001"
	findCategoryErr   appinfoHandlersErr = "appinfo-002"
	addCategoryErr    appinfoHandlersErr = "appinfo-003"
	deleteCategoryErr appinfoHandlersErr = "appinfo-004"
)

type IAppinfoHandler interface {
	GenerateApiKey(c *fiber.Ctx) error
	FindCategory(c *fiber.Ctx) error
	InsertCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}

type appinfoHandler struct {
	cfg            config.IConfig
	appinfoUsecase appinfoUsecases.IAppinfoUsecase
}

func AppinfoHandler(cfg config.IConfig, appinfoUsecase appinfoUsecases.IAppinfoUsecase) IAppinfoHandler {
	return &appinfoHandler{
		cfg:            cfg,
		appinfoUsecase: appinfoUsecase,
	}
}

func (h *appinfoHandler) GenerateApiKey(c *fiber.Ctx) error {
	apiKey, err := nftauth.NewAuth(
		nftauth.ApiKey,
		h.cfg.Jwt(),
		nil,
	)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.StatusInternalServerError,
			string(generateApiKeyErr),
			err.Error(),
		).Res()
	}
	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		&struct {
			Key string `json:"key"`
		}{
			Key: apiKey.SignToken(),
		},
	).Res()
}

func (h *appinfoHandler) FindCategory(c *fiber.Ctx) error {
	req := new(appinfo.CategoryFilter)
	if err := c.QueryParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.StatusBadRequest,
			string(findCategoryErr),
			err.Error(),
		).Res()
	}
	category, err := h.appinfoUsecase.FindCategory(req)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.StatusInternalServerError,
			string(findCategoryErr),
			err.Error(),
		).Res()
	}
	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		category,
	).Res()
}

func (h *appinfoHandler) InsertCategory(c *fiber.Ctx) error {
	category := make([]*appinfo.Category, 0)
	if err := c.BodyParser(&category); err != nil {
		return entities.NewResponse(c).Error(
			fiber.StatusBadRequest,
			string(addCategoryErr),
			err.Error(),
		).Res()
	}

	if len(category) == 0 {
		return entities.NewResponse(c).Error(
			fiber.StatusBadRequest,
			string(addCategoryErr),
			"categories request is empty",
		).Res()
	}

	err := h.appinfoUsecase.InsertCategory(category)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.StatusInternalServerError,
			string(addCategoryErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(
		fiber.StatusCreated,
		category,
	).Res()
}

func (h *appinfoHandler) DeleteCategory(c *fiber.Ctx) error {
	categoryId := c.Params("category_id")
	categoryIdInt, err := strconv.Atoi(categoryId)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.StatusBadRequest,
			string(deleteCategoryErr),
			err.Error(),
		).Res()
	}
	if categoryIdInt < 1 {
		return entities.NewResponse(c).Error(
			fiber.StatusBadRequest,
			string(deleteCategoryErr),
			"category id must be greater than zero",
		).Res()
	}
	err = h.appinfoUsecase.DeleteCategory(categoryId)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.StatusInternalServerError,
			string(deleteCategoryErr),
			err.Error(),
		).Res()
	}
	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		"category deleted successfully",
	).Res()
}
