package healthcheckhandlers

import (
	"github.com/IzePhanthakarn/go-phanthakarn-shop/config"
	"github.com/IzePhanthakarn/go-phanthakarn-shop/modules/entities"
	"github.com/IzePhanthakarn/go-phanthakarn-shop/modules/healthcheck"
	"github.com/gofiber/fiber/v2"
)

type IHealthcheckHandler interface {
	Healthcheck(c *fiber.Ctx) error
}

type healthcheckHandler struct {
	cfg config.IConfig
}

func HealthcheckHandler(ctg config.IConfig) IHealthcheckHandler {
	return &healthcheckHandler{
		cfg: ctg,
	}
}

func (h *healthcheckHandler) Healthcheck(c *fiber.Ctx) error {
	res := &healthcheck.Healthcheck{
		Name:    h.cfg.App().Name(),
		Version: h.cfg.App().Version(),
	}
	return entities.NewResponse(c).Success(fiber.StatusOK, res).Res()
}
