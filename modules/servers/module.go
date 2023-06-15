package servers

import (
	"github.com/IzePhanthakarn/go-phanthakarn-shop/modules/healthcheck/healthcheckHandlers"
	"github.com/gofiber/fiber/v2"
)

type IModuleFactory interface {
	HealthcheckModule()
}

type moduleFactory struct {
	router fiber.Router
	server *server
}

func InitModule(r fiber.Router, s *server) IModuleFactory {
	return &moduleFactory{
		router: r,
		server: s,
	}
}

func (m *moduleFactory) HealthcheckModule() {
	handler := healthcheckhandlers.HealthcheckHandler(m.server.cfg)

	m.router.Get("/", handler.Healthcheck)

}
