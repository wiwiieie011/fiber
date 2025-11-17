package rout

import (
	"wiwieie011/controllers"

	"github.com/gofiber/fiber/v2"
)

func RoutGroup(app *fiber.App) {
	userGroup := app.Group("/userlist")                  // все маршруты начинаются с /userlist
	userGroup.Get("/", controllers.GetUsers)             // GET /userlist
	userGroup.Get("/:id", controllers.GetUserByID)       // GET /userlist/:id
	userGroup.Post("/", controllers.CreateUser)          // POST /userlist
	userGroup.Put("/:id", controllers.PutUser)           // PUT /userlist/:id
	userGroup.Patch("/:id", controllers.PatchUser)       // PATCH /userlist/:id
	userGroup.Delete("/:id", controllers.DeleteUserByID) // DELETE /userlist/:id
	app.Listen(":3000")
}
