package monitorHandlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/muhammadfarhankt/nft-marketplace/config"
	"github.com/muhammadfarhankt/nft-marketplace/modules/monitor"
)

type IMontitorHandler interface {
	HealthCheck(c *fiber.Ctx) error
}

type monitorHandler struct {
	cfg config.IConfig
}

func MonitorHandler(cfg config.IConfig) IMontitorHandler {
	return &monitorHandler{
		cfg: cfg,
	}
}

func (h *monitorHandler) HealthCheck(c *fiber.Ctx) error {
	res := &monitor.Monitor{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	return c.Status(fiber.StatusOK).JSON(res)
}
