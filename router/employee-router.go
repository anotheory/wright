package router

import (
	"encoding/json"
	"fmt"
	"strconv"
	"wright/db"
	"wright/models/employee"
	"wright/models/exception"
	repo "wright/repository"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func EmployeeRouter(app *fiber.App) {
	employeeRouter := app.Group("/api/employees")

	employeeRouter.Get("/", func(c *fiber.Ctx) error {
		baseRepo := repo.BaseRepository{}
		employeeList := []employee.EmployeeEntity{}
		results := repo.IBaseRepository.Get(baseRepo, db.DbConnection, &employeeList)
		if results.Error != nil {
			c.SendString(results.Error.Error())
			return results.Error
		}
		c.JSON(employeeList)
		return nil
	})

	employeeRouter.Post("/", func(c *fiber.Ctx) error {
		request := employee.EmployeeRequest{}
		err := json.Unmarshal(c.Body(), &request)
		if err != nil {
			c.Status(500).JSON(exception.NewValidationException(err.Error()))
			return err
		}
		v := validator.New()
		err = v.Struct(request)
		if err != nil {
			errorMessage := ""
			validationErrors := err.(validator.ValidationErrors)
			for _, e := range validationErrors {
				errorMessage = fmt.Sprintf(`Validation failed => field: "%s", error: "%s"`, e.Field(), e.Tag())
			}
			c.Status(400).JSON(exception.NewValidationException(errorMessage))
			return err
		}
		toBeCreatedEmployee := employee.EmployeeEntity{
			PreName:   request.PreName,
			FirstName: *request.FirstName,
			LastName:  *request.LastName,
		}
		baseRepo := repo.BaseRepository{}
		results := repo.IBaseRepository.CreateAll(baseRepo, &toBeCreatedEmployee)
		if results.Error != nil {
			c.Status(500).JSON(exception.NewDbErrorException(results.Error.Error()))
			return err
		}
		c.JSON(toBeCreatedEmployee)
		return nil
	})

	employeeRouter.Patch("/:id", func(c *fiber.Ctx) error {
		request := employee.EmployeeRequest{}
		err := json.Unmarshal(c.Body(), &request)
		if err != nil {
			c.Status(500).JSON(exception.NewValidationException(err.Error()))
			return err
		}
		baseRepo := repo.BaseRepository{}
		toBeUpdatedEmployee := employee.EmployeeEntity{}
		selectedId, _ := strconv.Atoi(c.Params("id"))
		results := repo.IBaseRepository.GetById(baseRepo, uint(selectedId), &toBeUpdatedEmployee)
		if results.Error != nil {
			c.Status(500).JSON(exception.NewDbErrorException(results.Error.Error()))
			return err
		} else if results.RowsAffected == 0 {
			notFoundMessage := fmt.Sprintf("Selected ID [%d] not found", selectedId)
			c.Status(404).JSON(exception.NewNotFoundException(notFoundMessage))
			return nil
		}
		if request.PreName != nil {
			toBeUpdatedEmployee.PreName = request.PreName
		}
		if request.FirstName != nil {
			toBeUpdatedEmployee.FirstName = *request.FirstName
		}
		if request.LastName != nil {
			toBeUpdatedEmployee.LastName = *request.LastName
		}
		results = repo.IBaseRepository.UpdateAll(baseRepo, &toBeUpdatedEmployee)
		if results.Error != nil {
			c.Status(500).JSON(exception.NewDbErrorException(results.Error.Error()))
			return err
		}
		c.JSON(toBeUpdatedEmployee)
		return nil
	})

	employeeRouter.Delete("/:id", func(c *fiber.Ctx) error {
		baseRepo := repo.BaseRepository{}
		selectedId, _ := strconv.Atoi(c.Params("id"))
		toBeDeletedEmployee := employee.EmployeeEntity{}
		results := repo.IBaseRepository.GetById(baseRepo, uint(selectedId), &toBeDeletedEmployee)
		if results.Error != nil {
			c.Status(500).JSON(exception.NewDbErrorException(results.Error.Error()))
			return results.Error
		} else if results.RowsAffected == 0 {
			notFoundMessage := fmt.Sprintf("Selected ID [%d] not found", selectedId)
			c.Status(404).JSON(exception.NewNotFoundException(notFoundMessage))
			return nil
		}
		results = repo.IBaseRepository.DeleteAll(baseRepo, &toBeDeletedEmployee)
		if results.Error != nil {
			c.Status(500).JSON(exception.NewDbErrorException(results.Error.Error()))
			return results.Error
		}
		c.JSON(toBeDeletedEmployee)
		return nil
	})
}
