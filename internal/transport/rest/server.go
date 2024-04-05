package rest

import (
	"github.com/Alieksieiev0/user-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type RESTServer struct {
	app *fiber.App
}

func NewServer(app *fiber.App) *RESTServer {
	return &RESTServer{
		app: app,
	}
}

func (us *RESTServer) Start(addr string, service services.UserService) error {
	us.app.Use(logger.New(logger.Config{
		Format: "${time} | ${status} | ${latency} | ${method} | ${path} | ${error}\nResponse Body: ${resBody}\n",
	}))
	us.app.Use(cors.New())

	us.app.Get("/users/:id", getById(service))
	us.app.Get("/users", create(service))
	go us.app.Listen(addr)

	return us.app.Listen(addr)
}
