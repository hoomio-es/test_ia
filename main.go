package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jgargallo/test_app/controllers"
	"github.com/jgargallo/test_app/models"
)

func main() {
	// Inicializar la base de datos
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}

	// Crear servicio de propiedades
	propertyService := models.NewPropertyService(db)

	// Configurar el template engine
	engine := html.New("./templates", ".html")
	engine.Reload(true)

	// Crear controlador de propiedades
	propertyController := controllers.NewPropertyController(propertyService, engine)

	// Crear la aplicación Fiber
	app := fiber.New(fiber.Config{
		AppName: "Property Manager",
	})

	// Middleware global
	app.Use(logger.New())

	// Middleware
	app.Use(func(c *fiber.Ctx) error {
		return c.Next()
	})

	// Rutas de API
	api := app.Group("/api")

	api.Post("/properties", propertyController.CreateProperty)
	api.Get("/properties/:id", propertyController.PropertyIDMiddleware, propertyController.GetProperty)
	api.Put("/properties/:id", propertyController.PropertyIDMiddleware, propertyController.UpdateProperty)
	api.Delete("/properties/:id", propertyController.PropertyIDMiddleware, propertyController.DeleteProperty)
	api.Get("/search", propertyController.SearchProperties)

	// Rutas de la interfaz web
	app.Get("/", propertyController.ListProperties)
	app.Get("/properties/:id", propertyController.PropertyIDMiddleware, propertyController.ShowProperty)

	// Servir archivos estáticos
	app.Static("/", "./public")

	// Iniciar servidor
	fmt.Println("Starting server on port :3000")
	fmt.Println("Press Ctrl+C to quit")
	app.Listen(":3000")
}
