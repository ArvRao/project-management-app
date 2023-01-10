package handler

import (
	"fmt"

	"github.com/arvrao/project-management-app/database"
	"github.com/arvrao/project-management-app/model"
	"github.com/gofiber/fiber/v2"
)

// get all projects
func GetAllTodosByProjectId(c *fiber.Ctx) error {
	db := database.DB.Db
	var todos []model.Todo
	// var projects []model.Project	// not required
	// type Todo model.Todo
	// var Project model.Project
	id := c.Params("projectId")
	fmt.Println("id: ", id)
	// db.Find(&projects, "ID = ?", id) 	// not required
	todosVh := db.Debug().Joins("Project").Find(&todos, "project_id_fk = ?", id)
	fmt.Println("todosVh: ", todosVh)

	if len(todos) == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No todos available right now", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "All Todos are found", "data": todos})
}

func CreateTodo(c *fiber.Ctx) error {
	db := database.DB.Db
	todo := new(model.Todo)
	err := c.BodyParser(todo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Something's wrong with your input", "data": err})
	}
	err = db.Create(&todo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create todo", "data": err})
	}
	// Return the created todo
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Project has created", "data": todo})
}
