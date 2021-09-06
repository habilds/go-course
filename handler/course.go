package handler

import (
	"fmt"
	"strings"

	"github.com/habilds/go-course/database"
	"github.com/habilds/go-course/models"

	"github.com/gofiber/fiber/v2"
)

// GetAllCourses query all courses
func GetAllCourses(c *fiber.Ctx) error {
	db := database.DB.Db
	var courses []models.Course
	q := db
	if c.Query("q") != "" {
		searchQ := fmt.Sprintf("%%%s%%", c.Query("q"))
		q = q.Where("title LIKE ?", searchQ)
	}
	if c.Query("sort") != "" {
		sortQ := strings.Replace(c.Query("sort"), "_", " ", -1)
		q = q.Order(sortQ)
	}
	if c.Query("free") == "true" || c.Query("free") == "1" {
		q = q.Where("price = ?", 0)
	}
	q = q.Find(&courses)
	return c.JSON(fiber.Map{"status": "success", "message": "All courses", "data": courses})
}

// GetTotalCourses query all courses
func GetTotalCourses(c *fiber.Ctx) error {
	db := database.DB.Db
	var courses []models.Course
	var count int64
	db.Find(&courses).Count(&count)
	return c.JSON(fiber.Map{"status": "success", "message": "All courses", "data": count})
}

// GetTotalFreeCourses query all courses
func GetTotalFreeCourses(c *fiber.Ctx) error {
	db := database.DB.Db
	var courses []models.Course
	var count int64
	db.Where("price = ?", 0).Find(&courses).Count(&count)
	return c.JSON(fiber.Map{"status": "success", "message": "All courses", "data": count})
}

// GetCourse query course
func GetCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var course models.Course
	db.Db.Find(&course, id)
	if course.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No course found with ID", "data": nil})

	}
	return c.JSON(fiber.Map{"status": "success", "message": "Course found", "data": course})
}

// CreateCourse new course
func CreateCourse(c *fiber.Ctx) error {
	db := database.DB
	course := new(models.Course)
	if err := c.BodyParser(course); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create course", "data": err})
	}
	db.Db.Create(&course)
	return c.JSON(fiber.Map{"status": "success", "message": "Created course", "data": course})
}

// DeleteCourse delete course
func DeleteCourse(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB

	var course models.Course
	db.Db.First(&course, id)
	if course.Title == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No course found with ID", "data": nil})

	}
	db.Db.Delete(&course)
	return c.JSON(fiber.Map{"status": "success", "message": "Course successfully deleted", "data": nil})
}
