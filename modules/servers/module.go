package servers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/muhammadfarhankt/nft-marketplace/modules/appinfo/appinfoHandlers"
	"github.com/muhammadfarhankt/nft-marketplace/modules/appinfo/appinfoRepositories"
	"github.com/muhammadfarhankt/nft-marketplace/modules/appinfo/appinfoUsecases"

	filesUsecases "github.com/muhammadfarhankt/nft-marketplace/modules/files/fileUsecases"
	"github.com/muhammadfarhankt/nft-marketplace/modules/files/filesHandlers"

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
	AppinfoModule()
	FilesModule()
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
	router.Post("/signup", m.mid.ApiKeyAuth(), handler.SignUpCustomer)
	router.Post("/signin", m.mid.ApiKeyAuth(), handler.SignIn)
	router.Post("/refresh", m.mid.ApiKeyAuth(), handler.RefreshPassport)
	router.Post("/signout", m.mid.ApiKeyAuth(), handler.SignOut)

	router.Get("/:user_id", m.mid.JwtAuth(), m.mid.ParamsCheck(), handler.GetUserProfile)

	router.Get("/admin/generate-token", m.mid.JwtAuth(), m.mid.Authorize(2), handler.GenerateAdminToken)
	router.Post("/signup-admin", m.mid.JwtAuth(), m.mid.Authorize(2), handler.SignUpAdmin)
}

func (m *moduleFactory) AppinfoModule() {
	repository := appinfoRepositories.AppinfoRepository(m.s.db)
	usecase := appinfoUsecases.AppinfoUsecase(repository)
	handler := appinfoHandlers.AppinfoHandler(m.s.cfg, usecase)

	router := m.r.Group("/appinfo")

	router.Get("/apikey", m.mid.JwtAuth(), m.mid.Authorize(2), handler.GenerateApiKey)

	router.Get("/categories", m.mid.ApiKeyAuth(), handler.FindCategory)
	router.Post("/categories", m.mid.JwtAuth(), m.mid.Authorize(2), handler.InsertCategory)
	router.Delete("/delete-category/:category_id", m.mid.JwtAuth(), m.mid.Authorize(2), handler.DeleteCategory)
}

func (m *moduleFactory) FilesModule() {
	usecase := filesUsecases.FilesUsecase(m.s.cfg)
	handler := filesHandlers.FilesHandler(m.s.cfg, usecase)

	router := m.r.Group("/files")

	//_ = handler
	//_ = router

	router.Post("/upload", m.mid.JwtAuth(), m.mid.Authorize(2), handler.UploadToGCP)

	router.Patch("/delete", m.mid.JwtAuth(), m.mid.Authorize(2), handler.DeleteFromGCP)
}
