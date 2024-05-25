package main

import (
	"wright/db"
	"wright/models/employee"
	"wright/models/healthcheck"
	"wright/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db.DbConnection.AutoMigrate(&healthcheck.HealthcheckEntity{}, &employee.EmployeeEntity{})
	app := fiber.New()
	app = router.InitRouter(app)
	app.Use(cors.New())
	app.Listen(":3000")
}
