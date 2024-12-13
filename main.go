package main

import(
	"github.com/gofiber/fiber/v2"
)


func main(){
	app := fiber.New()
	app.Get("/", handleFoo)
	app.Listen("localhost:8080")
}


func handleFoo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg":"welcome"})
}
