package router

import (
	"github.com/arvrao/kanban-board-app/handler"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupRoutes(app *fiber.App) {
	// grouping
	api := app.Group("/api")
	// routes
	api.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON("Hello there @arvind")
	})

	api.Get("/getAllProjects", handler.GetAllProjects)
	api.Get("/getProject/:id", handler.GetProjectById)
	api.Post("/createNewProject", handler.CreateProject)
	api.Patch("/updateProjectById/:id", handler.UpdateProjectById)
	api.Delete("/deleteProject/:id", handler.DeleteProjectbyId)
}