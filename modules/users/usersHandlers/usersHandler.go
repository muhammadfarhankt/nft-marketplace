package usersHandlers

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/muhammadfarhankt/nft-marketplace/config"
	"github.com/muhammadfarhankt/nft-marketplace/modules/entities"
	"github.com/muhammadfarhankt/nft-marketplace/modules/users"
	"github.com/muhammadfarhankt/nft-marketplace/modules/users/usersUsecases"
	"github.com/muhammadfarhankt/nft-marketplace/pkg/nftauth"
)

type usersHandlersErrCode string

const (
	signUpCustomerErr     usersHandlersErrCode = "users-error-001"
	signInErr             usersHandlersErrCode = "users-error-002"
	refreshErr            usersHandlersErrCode = "users-error-003"
	logoutErr             usersHandlersErrCode = "users-error-004"
	signUpAdminErr        usersHandlersErrCode = "users-error-005"
	generateAdminTokenErr usersHandlersErrCode = "users-error-006"
	getUserProfileErr     usersHandlersErrCode = "users-error-007"
)

type IUsersHandler interface {
	SignUpCustomer(c *fiber.Ctx) error
	SignIn(c *fiber.Ctx) error
	RefreshPassport(c *fiber.Ctx) error
	SignOut(c *fiber.Ctx) error
	SignUpAdmin(c *fiber.Ctx) error
	GenerateAdminToken(c *fiber.Ctx) error
	GetUserProfile(c *fiber.Ctx) error
}

type usersHandler struct {
	cfg         config.IConfig
	userUsecase usersUsecases.IUsersUsecase
}

func UsersHandler(cfg config.IConfig, usersUsecase usersUsecases.IUsersUsecase) IUsersHandler {
	return &usersHandler{
		cfg:         cfg,
		userUsecase: usersUsecase,
	}
}

func (u *usersHandler) SignUpCustomer(c *fiber.Ctx) error {
	// Request body parser to get data
	req := new(users.UserRegisterReq)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signUpCustomerErr),
			err.Error(),
		).Res()
	}

	//email validation
	if !req.IsEmail() {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signUpCustomerErr),
			"Invalid Email pattern",
		).Res()
	}

	// Insertion
	result, err := u.userUsecase.InsertCustomer(req)
	if err != nil {
		switch err.Error() {
		case "username already exists":
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(signUpCustomerErr),
				err.Error(),
			).Res()
		case "email already exists":
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(signUpCustomerErr),
				err.Error(),
			).Res()
		default:
			return entities.NewResponse(c).Error(
				fiber.ErrInternalServerError.Code,
				string(signUpCustomerErr),
				err.Error(),
			).Res()

		}
	}
	//201 Created status code means that the request was successfully fulfilled and resulted in one or possibly multiple new resources being created.
	return entities.NewResponse(c).Success(fiber.StatusCreated, result).Res()
}

func (h *usersHandler) SignIn(c *fiber.Ctx) error {
	req := new(users.UserCredential)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signInErr),
			err.Error(),
		).Res()
	}

	passport, err := h.userUsecase.GetPassport(req)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signInErr),
			err.Error(),
		).Res()
	}
	return entities.NewResponse(c).Success(fiber.StatusOK, passport).Res()
}

func (u *usersHandler) SignUpAdmin(c *fiber.Ctx) error {
	// Request body parser to get data
	req := new(users.UserRegisterReq)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signUpAdminErr),
			err.Error(),
		).Res()
	}

	//email validation
	if !req.IsEmail() {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(signUpAdminErr),
			"Invalid Email pattern",
		).Res()
	}

	// Insertion
	result, err := u.userUsecase.InsertAdmin(req)
	if err != nil {
		switch err.Error() {
		case "username already exists":
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(signUpAdminErr),
				err.Error(),
			).Res()
		case "email already exists":
			return entities.NewResponse(c).Error(
				fiber.ErrBadRequest.Code,
				string(signUpAdminErr),
				err.Error(),
			).Res()
		default:
			return entities.NewResponse(c).Error(
				fiber.ErrInternalServerError.Code,
				string(signUpAdminErr),
				err.Error(),
			).Res()

		}
	}
	//201 Created status code means that the request was successfully fulfilled and resulted in one or possibly multiple new resources being created.
	return entities.NewResponse(c).Success(fiber.StatusCreated, result).Res()
}

func (u *usersHandler) GenerateAdminToken(c *fiber.Ctx) error {
	adminToken, err := nftauth.NewAuth(
		nftauth.Admin,
		u.cfg.Jwt(),
		nil,
	)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrInternalServerError.Code,
			string(generateAdminTokenErr),
			err.Error(),
		).Res()
	}
	return entities.NewResponse(c).Success(
		fiber.StatusOK,
		&struct {
			Token string `json:"token"`
		}{
			Token: adminToken.SignToken(),
		},
	).Res()
}

func (h *usersHandler) RefreshPassport(c *fiber.Ctx) error {
	req := new(users.UserRefreshCredential)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(refreshErr),
			err.Error(),
		).Res()
	}

	passport, err := h.userUsecase.RefreshPassport(req)
	if err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(refreshErr),
			err.Error(),
		).Res()
	}
	return entities.NewResponse(c).Success(fiber.StatusOK, passport).Res()
}

func (h *usersHandler) SignOut(c *fiber.Ctx) error {

	req := new(users.UserRemoveCredential)
	if err := c.BodyParser(req); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(logoutErr),
			err.Error(),
		).Res()
	}
	fmt.Println("Users Handler req : ", req)
	if err := h.userUsecase.DeleteOauth(req.OauthId); err != nil {
		return entities.NewResponse(c).Error(
			fiber.ErrBadRequest.Code,
			string(logoutErr),
			err.Error(),
		).Res()
	}

	return entities.NewResponse(c).Success(fiber.StatusOK, "logout sucess").Res()
}

func (h *usersHandler) GetUserProfile(c *fiber.Ctx) error {
	// Get user_id from params
	userID := strings.Trim(c.Params("user_id"), " ")
	profile, err := h.userUsecase.GetUserProfile(userID)
	if err != nil {
		switch err.Error() {
		case "get user failed: sql: no rows in result set":
			return entities.NewResponse(c).Error(
				fiber.ErrNotFound.Code,
				string(getUserProfileErr),
				err.Error(),
			).Res()
		default:
			return entities.NewResponse(c).Error(
				fiber.ErrInternalServerError.Code,
				string(getUserProfileErr),
				err.Error(),
			).Res()
		}
	}
	return entities.NewResponse(c).Success(fiber.StatusOK, profile).Res()
}
