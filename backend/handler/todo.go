package handler

import (
	"fmt"

	"github.com/arvrao/project-management-app/database"
	"github.com/arvrao/project-management-app/model"
	"github.com/gofiber/fiber/v2"
)

// project FK table -> only get the project name
func GetAllTodosByProjectId(c *fiber.Ctx) error {
	db := database.DB.Db
	// var todos []model.Todo
	var todosAPI []model.TodoApiStruct
	id := c.Params("projectId")
	fmt.Println("id: ", id)
	todosVh := db.Debug().Model(&model.Todo{}).Joins("Project").Select("TodoName", "ProjectIdFk").Find(&todosAPI, "project_id_fk = ?", id)
	fmt.Println("todosVh: ", todosVh)

	if len(todosAPI) == 0 {
		return c.Status(404).JSON(fiber.Map{"message": "No todos available right now", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "All Todos are found", "data": todosAPI})
}

func GetTodobyId(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("id")
	var todo model.Todo
	// find single user in the database by id
	db.Find(&todo, "id = ?", id)
	if todo.ID != 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Project not found", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "Project Found", "data": todo})
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
	return c.Status(201).JSON(fiber.Map{"status": "success", "message": "Todo has been created", "data": todo})
}

func UpdateTodo(c *fiber.Ctx) error {
	db := database.DB.Db
	id := c.Params("todoId")
	todo := new(model.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Where("id = ?", id).Updates(&todo)
	return c.Status(200).JSON(todo)
}

func DeleteTodo(c *fiber.Ctx) error {
	db := database.DB.Db
	var todo model.Todo
	// get id params
	todoId := c.Params("todoId")
	fmt.Println("todo: ", todo)
	// if todo.ID == 0 {
	// 	return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Todo not found", "data": nil})
	// }
	err := db.Delete(&todo, "id = ?", todoId).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete project", "data": nil})
	}
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": `Given todo ${todoId} has been deleted`})
}
