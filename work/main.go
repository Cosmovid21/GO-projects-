package main

import (
	"log"

	database "github.com/cosmovid21/fiber/database"
	"github.com/cosmovid21/fiber/routes"
	"github.com/gofiber/fiber/v2"
)



func welcome(c *fiber.Ctx) error {
	return c.SendString("hello whateverr")
}

func setupRoutes(app *fiber.App) {

	app.Get("/api", welcome)
	//user enpoints
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.Updateuser)
	app.Delete("/api/users/:id",routes.DeleteUser)
	//product endpoints 
	//app.Post("/api/Products", routes.CreateProduct)
	//app.Get("/api/Products", routes.GetProducts)

}



func main() {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
