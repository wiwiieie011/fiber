package main

import (
	"os"
	"wiwieie011/base"
	"wiwieie011/rout"

	"github.com/gofiber/fiber/v2"
)



func main() {
	base.LoadEnvVariables()
	base.ConnectionDB()
	app:= fiber.New()
	rout.RoutGroup(app)
	port :=os.Getenv("PORT")
	app.Listen(":"+port)
}