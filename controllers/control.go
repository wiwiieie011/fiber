package controllers

import (
	"wiwieie011/base"
	"wiwieie011/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func GetUsers(c *fiber.Ctx) error {
	var userList []models.User
	if err := base.DB.Find(&userList).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(userList)
}

func GetUserByID(c *fiber.Ctx) error {
	var userID models.User
	if err := base.DB.Where("id = ?", c.Params("id")).First(&userID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})

	}

	return c.JSON(userID)
}

func CreateUser(c *fiber.Ctx) error {
	var userInput = new(models.InputUser)

	if err := c.BodyParser(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})

	}

	if err := validate.Struct(userInput); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	mainUser := models.User{Name: userInput.Name, Email: userInput.Email, Age: userInput.Age}

	base.DB.Create(&mainUser)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"created": true, "user": mainUser})

}

func PatchUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.UpdateUser
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid json",
		})
	}

	err := base.DB.Model(&models.User{}).Where("id = ?", id).Updates(data).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"updated": true})
}

func PutUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var data models.PutUpdateUser

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid json",
		})
	}

	if err := validate.Struct(data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := base.DB.Model(&models.User{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"updated": true})
}

func DeleteUserByID(c *fiber.Ctx) error {
	var userDelete models.User
	err := base.DB.Where("id = ?", c.Params("id")).Unscoped().Delete(&userDelete).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"deleted": true})
}
