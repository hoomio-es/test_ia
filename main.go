package main

import (
	"database/sql"
	"log"
	"os"
	"strings" // Importado para manipulación de strings
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jgargallo/property-manager/controllers" // Asegúrate que la ruta del módulo sea correcta
	"github.com/jgargallo/property-manager/models"      // Asegúrate que la ruta del módulo sea correcta
	_ "github.com/lib/pq"                               // PostgreSQL driver
)

var db *sql.DB

func init() {
	var err error
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	// Retry connecting to the database
	maxRetries := 10
	retryDelay := 2 * time.Second

	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("postgres", dsn)
		if err != nil {
			log.Printf("Error opening database connection: %v. Retrying in %v...", err, retryDelay)
			time.Sleep(retryDelay)
			continue
		}

		err = db.Ping()
		if err != nil {
			log.Printf("Error pinging database: %v. Retrying in %v...", err, retryDelay)
			time.Sleep(retryDelay)
			continue
		}

		log.Println("Successfully connected to PostgreSQL")
		break
	}

	if err != nil {
		log.Fatalf("Could not connect to the database after %d attempts", maxRetries)
	}

	// Create tables if they don't exist
	schemaBytes, err := os.ReadFile("db/schema.sql")
	if err != nil {
		log.Fatal("Error reading schema.sql file:", err)
	}

	// CORRECCIÓN: Hacer que la creación de la tabla sea idempotente
	// reemplazando "CREATE TABLE" por "CREATE TABLE IF NOT EXISTS".
	schema := string(schemaBytes)
	schema = strings.Replace(schema, "CREATE TABLE", "CREATE TABLE IF NOT EXISTS", 1)

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal("Error executing schema:", err)
	}
	log.Println("Database schema loaded successfully")
}

func main() {
	// Load templates
	engine := html.New("./templates", ".html")
	if err := engine.Load(); err != nil {
		log.Fatalf("error loading templates: %v", err)
	}

	// Create services and controllers
	propertyService := models.NewPropertyService(db)
	// CORRECTED: Remove the 'engine' argument from the call
	propertyController := controllers.NewPropertyController(propertyService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Static files
	app.Static("/", "./public")

	// Routes
	app.Get("/", propertyController.ListProperties)
	app.Get("/property/:id", propertyController.ShowProperty)

	// API Routes
	api := app.Group("/api")
	api.Post("/properties", propertyController.CreateProperty)
	api.Put("/properties/:id", propertyController.UpdateProperty)
	api.Delete("/properties/:id", propertyController.DeleteProperty)
	api.Get("/properties/:id", propertyController.GetProperty)
	api.Get("/search", propertyController.SearchProperties)

	// Start server
	log.Fatal(app.Listen(":8080"))
}
