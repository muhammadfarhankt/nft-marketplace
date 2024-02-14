package servers

import (
	"github.com/gofiber/fiber/v2"
	middlewareHandlers "github.com/muhammadfarhankt/nft-marketplace/modules/middlewares/middlewaresHandlers"
	"github.com/muhammadfarhankt/nft-marketplace/modules/middlewares/middlewaresRepositories"
	"github.com/muhammadfarhankt/nft-marketplace/modules/middlewares/middlewaresUsecases"
	"github.com/muhammadfarhankt/nft-marketplace/modules/monitor/monitorHandlers"
	"github.com/muhammadfarhankt/nft-marketplace/modules/users/usersHandlers"
	"github.com/muhammadfarhankt/nft-marketplace/modules/users/usersRepositories"
	"github.com/muhammadfarhankt/nft-marketplace/modules/users/usersUsecases"
)

type IModuleFactory interface {
	MonitorModule()
	UserModule()
}

type moduleFactory struct {
	r   fiber.Router
	s   *server
	mid middlewareHandlers.NMiddlewaresHandler
}

func InitModule(r fiber.Router, s *server, mid middlewareHandlers.NMiddlewaresHandler) IModuleFactory {
	return &moduleFactory{
		r:   r,
		s:   s,
		mid: mid,
	}
}

func InitMiddlewares(s *server) middlewareHandlers.NMiddlewaresHandler {
	repository := middlewaresRepositories.MiddlewaresRepository(s.db)
	usecase := middlewaresUsecases.MiddlewaresUsecase(repository)
	return middlewareHandlers.MiddlewaresHandler(s.cfg, usecase)
}

func (m *moduleFactory) MonitorModule() {
	handler := monitorHandlers.MonitorHandler(m.s.cfg)

	m.r.Get("/", handler.HealthCheck)
}

func (m *moduleFactory) UserModule() {
	repository := usersRepositories.UsersRepository(m.s.db)
	usecase := usersUsecases.UsersUsecase(m.s.cfg, repository)
	handler := usersHandlers.UsersHandler(m.s.cfg, usecase)

	router := m.r.Group("/users")
	router.Post("/signup", handler.SignUpCustomer)
	router.Post("/signin", handler.SignIn)
	router.Post("/refresh", handler.RefreshPassport)
	router.Post("/signout", handler.SignOut)
}
