package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jgargallo/property-manager/controllers"
	"github.com/jgargallo/property-manager/models"
	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

func init() {
	var err error
	dsn := "postgresql://postgres:postgres@db:5432/test_app?sslmode=disable"

	// Intentar conectar con un máximo de 10 intentos
	maxRetries := 10
	retryDelay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		if i > 0 {
			log.Printf("Intento de conexión %d/%d...", i+1, maxRetries)
			time.Sleep(retryDelay)
		}

		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("Error opening database connection: %v", err)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Error pinging database: %v", err)
			continue
		}

		log.Printf("Conexión exitosa con PostgreSQL")
		break
	}

	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos después de %d intentos", maxRetries)
	}

	// Create tables if they don't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS properties (
		    id SERIAL PRIMARY KEY,
		    nombre_alias VARCHAR(255) NOT NULL,
		    direccion_completa TEXT NOT NULL,
		    tipo_propiedad VARCHAR(50) NOT NULL,
		    capacidad_maxima INTEGER NOT NULL
		);
	`)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}

	// Debug: Check if templates directory exists
	if _, err := os.Stat("/app/templates"); err != nil {
		log.Printf("WARNING: Templates directory not found: %v", err)
	}

	// Debug: Check if index.html exists
	if _, err := os.Stat("/app/templates/index.html"); err != nil {
		log.Printf("WARNING: index.html not found: %v", err)
	}

	// Debug: List files in templates directory
	files, err := os.ReadDir("/app/templates")
	if err != nil {
		log.Printf("Error reading templates directory: %v", err)
	} else {
		log.Printf("Files in templates directory:")
		for _, file := range files {
			log.Printf("- %s", file.Name())
		}
	}

	// Debug: Test template engine
	engine := html.New("/app/templates", ".html")
	if err := engine.Load(); err != nil {
		log.Printf("Error loading templates: %v", err)
	}
}

func main() {
	// Load templates
	engine := html.New("/app/templates", ".html")

	// Create property service and controller
	propertyService := models.NewPropertyService(db)
	propertyController := controllers.NewPropertyController(propertyService, engine)

	// Create Fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Static files
	app.Static("/", "/app/public")

	// Routes
	app.Get("/", propertyController.ListProperties)
	app.Get("/property/:id", propertyController.ShowProperty)
	app.Post("/api/properties", propertyController.CreateProperty)
	app.Put("/api/properties/:id", propertyController.UpdateProperty)
	app.Delete("/api/properties/:id", propertyController.DeleteProperty)
	app.Get("/api/properties/:id", propertyController.GetProperty)
	app.Get("/api/search", propertyController.SearchProperties)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
