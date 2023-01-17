package handler

import (
	"github.com/arvrao/project-management-app/database"
	"github.com/arvrao/project-management-app/model"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// get all projects
func GetAllProjects(c *fiber.Ctx) error {
	db := database.DB.Db
	var projects []model.Project
	// find all projects in the database
	db.Find(&projects)
	if len(projects) == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No projects available right now", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "All Projects are found", "data": projects})
}

func CreateProject(c *fiber.Ctx) error {
	// generate new Uid with that
	db := database.DB.Db
	project := new(model.Project)
	// Store the body in project and return error if encountered
	err := c.BodyParser(project)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&project).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create project", "data": err})
	}
	// Return the created project
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Project has created", "data": project})
}

func GetProjectById(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var user model.Project
	// find single user in the database by id
	db.Find(&user, "id = ?", id)
	if user.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Project not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Project Found", "data": user})
}

func UpdateProjectById(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	project := new(model.Project)
	if err := c.BodyParser(project); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&project)
	return c.Status(200).JSON("About value changed to -> " + project.About)
}

func DeleteProjectbyId(c *fiber.Ctx) error {
	db := database.DB.Db
	var project model.Project
	// get id params
	id := c.Params("id")
	db.Find(&project, "id = ?", id)
	if project.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Project not found", "data": nil})
	}
	err := db.Delete(&project, "id = ?", id).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete project", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Project deleted"})
}
