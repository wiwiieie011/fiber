package controllers

import (
	"wiwieie011/base"
	"wiwieie011/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func GetUsers(c *fiber.Ctx) error {
	var s []models.User
	if err := base.DB.Find(&s).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(s)
}

func GetUserByID(c *fiber.Ctx) error {
	var s models.User
	if err := base.DB.Where("id = ?", c.Params("id")).First(&s).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})

	}

	return c.JSON(s)
}

func CreateUser(c *fiber.Ctx) error {
	var s = new(models.InputUser)

	if err := c.BodyParser(s); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err})

	}

	 if err := validate.Struct(s); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
    }

	a := models.User{Name: s.Name, Email: s.Email, Age: s.Age}

	base.DB.Create(&a)
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"created": true, "user": a})

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
	var s models.User
	err := base.DB.Where("id = ?", c.Params("id")).Unscoped().Delete(&s).Error

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"deleted": true})
}
