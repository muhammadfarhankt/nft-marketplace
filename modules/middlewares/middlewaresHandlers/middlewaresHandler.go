package middlewareHandlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/muhammadfarhankt/nft-marketplace/config"

	"github.com/muhammadfarhankt/nft-marketplace/modules/entities"
	"github.com/muhammadfarhankt/nft-marketplace/modules/middlewares/middlewaresUsecases"

	"github.com/muhammadfarhankt/nft-marketplace/pkg/nftauth"
	"github.com/muhammadfarhankt/nft-marketplace/pkg/utils"
)

type middlewareHandlersErrCode string

const (
	routerCheckErr middlewareHandlersErrCode = "middleware-001"
	jwtAuthErr     middlewareHandlersErrCode = "middleware-002"
	paramsCheckErr middlewareHandlersErrCode = "middleware-003"
	auhorizeErr    middlewareHandlersErrCode = "middleware-004"
	apiKeyErr      middlewareHandlersErrCode = "middleware-005"
)

type NMiddlewaresHandler interface {
	Cors() fiber.Handler
	RouterCheck() fiber.Handler
	Logger() fiber.Handler
	JwtAuth() fiber.Handler
	ParamsCheck() fiber.Handler
	Authorize(expectedRoleId ...int) fiber.Handler
	ApiKeyAuth() fiber.Handler
}

type middlewaresHandler struct {
	cfg                config.IConfig
	middlewaresUsecase middlewaresUsecases.NMiddlewaresUsecase
}

func MiddlewaresHandler(cfg config.IConfig, middlewaresUsecase middlewaresUsecases.NMiddlewaresUsecase) NMiddlewaresHandler {
	return &middlewaresHandler{
		cfg:                cfg,
		middlewaresUsecase: middlewaresUsecase,
	}
}

func (h *middlewaresHandler) Cors() fiber.Handler {
	return cors.New(cors.Config{
		Next:             cors.ConfigDefault.Next,
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,HEAD",
		AllowHeaders:     "",
		AllowCredentials: false,
		ExposeHeaders:    "",
		MaxAge:           0,
	})
}

func (h *middlewaresHandler) RouterCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return entities.NewResponse(c).Error(
			fiber.ErrNotFound.Code,
			string(routerCheckErr),
			"router not found",
		).Res()
	}
}

func (h *middlewaresHandler) Logger() fiber.Handler {
	return logger.New(logger.Config{
		Format:     "${time} [${ip}] ${status} - ${method} ${path} ${latency} \n",
		TimeFormat: "02-01-2006 15:04:05",
		TimeZone:   "Asia/Mumbai",
	})
}

func (h *middlewaresHandler) JwtAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := strings.TrimPrefix(c.Get("Authorization"), "Bearer ")
		result, err := nftauth.ParseToken(h.cfg.Jwt(), token)
		if err != nil {
			return entities.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(jwtAuthErr),
				err.Error(),
			).Res()
		}
		claims := result.Claims
		if !h.middlewaresUsecase.FindAccessToken(claims.Id, token) {
			return entities.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(jwtAuthErr),
				"no permission to access token / invalid access token",
			).Res()
		}
		//set UserId
		c.Locals("userId", claims.Id)
		c.Locals("userRoleId", claims.RoleId)
		return c.Next()
	}
}

func (h *middlewaresHandler) ParamsCheck() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userId := c.Locals("userId").(string)
		if userId != c.Params("user_id") {
			return entities.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(paramsCheckErr),
				"no permission to access this user profile",
			).Res()
		}
		return c.Next()
	}
}

func (h *middlewaresHandler) Authorize(expectedRoleId ...int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userRoleId, ok := c.Locals("userRoleId").(int)
		if !ok {
			return entities.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(auhorizeErr),
				"userRoleId not int type",
			).Res()
		}

		roles, err := h.middlewaresUsecase.FindRole()
		fmt.Println("Roles : ", roles)
		if err != nil {
			return entities.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(auhorizeErr),
				err.Error(),
			).Res()
		}

		sum := 0
		for _, v := range expectedRoleId {
			sum += v
		}

		expectedValueBinary := utils.BinaryConverter(sum, len(roles))
		userValueBinary := utils.BinaryConverter(userRoleId, len(roles))

		for i := range userValueBinary {
			if userValueBinary[i]&expectedValueBinary[i] == 1 {
				return c.Next()
			}
		}
		return entities.NewResponse(c).Error(
			fiber.ErrUnauthorized.Code,
			string(auhorizeErr),
			"no permission to access this route",
		).Res()
	}
}

func (h *middlewaresHandler) ApiKeyAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		apiKey := c.Get("X-Api-Key")
		if _, err := nftauth.ParseApiKey(h.cfg.Jwt(), apiKey); err != nil {
			return entities.NewResponse(c).Error(
				fiber.ErrUnauthorized.Code,
				string(apiKeyErr),
				"api key required or invalid",
			).Res()
		}
		return c.Next()
	}
}
