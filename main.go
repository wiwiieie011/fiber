package main

import (
	"wiwieie011/base"
	"wiwieie011/rout"

	"github.com/gofiber/fiber/v2"
)



func main() {
	base.ConnectionDB()
	app:= fiber.New()
	rout.RoutGroup(app)
	app.Listen(":3000")
}