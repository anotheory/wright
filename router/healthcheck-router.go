package router

import (
	"wright/models/healthcheck"
	repo "wright/repository"

	"github.com/gofiber/fiber/v2"
)

func HealthcheckRouter(app *fiber.App) {
	healthcheckRouter := app.Group("/healthcheck")

	healthcheckRouter.Get("/", func(c *fiber.Ctx) error {
		baseRepo := repo.BaseRepository{}
		healthcheckResponse := healthcheck.HealthcheckEntity{}
		results := repo.IBaseRepository.GetById(baseRepo, 1, &healthcheckResponse)
		if results.Error != nil {
			c.SendString(results.Error.Error())
			return results.Error
		}
		c.JSON(healthcheck.HealthcheckResponse{
			Version: "latest",
			Message: healthcheckResponse.Message,
		})
		return nil
	})
}
