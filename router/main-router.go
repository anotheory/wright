package router

import "github.com/gofiber/fiber/v2"

func InitRouter(app *fiber.App) *fiber.App {
	HealthcheckRouter(app)
	EmployeeRouter(app)

	return app
}
