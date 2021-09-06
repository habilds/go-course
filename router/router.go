package router

import (
	"github.com/habilds/go-course/handler"
	"github.com/habilds/go-course/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handler.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handler.Login)

	// User
	user := api.Group("/user")
	user.Get("/total", handler.GetTotalUser)
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", middleware.Protected(), handler.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handler.DeleteUser)

	// Course
	course := api.Group("/course")
	course.Get("/total", handler.GetTotalCourses)
	course.Get("/total-free", handler.GetTotalFreeCourses)
	course.Get("/", handler.GetAllCourses)
	course.Get("/:id", handler.GetCourse)
	course.Post("/", middleware.Protected(), handler.CreateCourse)
	course.Delete("/:id", middleware.Protected(), handler.DeleteCourse)
}
