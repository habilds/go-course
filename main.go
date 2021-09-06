package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/habilds/go-course/database"
	"github.com/habilds/go-course/router"

	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Use(cors.New())

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
