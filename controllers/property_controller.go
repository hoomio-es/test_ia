package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jgargallo/property-manager/models" // Asegúrate que la ruta del módulo sea correcta
)

// PropertyController handles property-related routes
type PropertyController struct {
	propertyService *models.PropertyService
}

// NewPropertyController creates a new property controller
func NewPropertyController(propertyService *models.PropertyService) *PropertyController {
	return &PropertyController{
		propertyService: propertyService,
	}
}

// getPropertyID is a helper to extract and convert the ID from the request parameters.
func getPropertyID(ctx *fiber.Ctx) (int, error) {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

// ListProperties renders the properties list page
func (c *PropertyController) ListProperties(ctx *fiber.Ctx) error {
	properties, err := c.propertyService.GetAllProperties()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.Render("index", fiber.Map{
		"Properties": properties,
	})
}

// ShowProperty renders the property details page
func (c *PropertyController) ShowProperty(ctx *fiber.Ctx) error {
	id, err := getPropertyID(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).SendString("Invalid ID")
	}

	property, err := c.propertyService.GetPropertyByID(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	if property == nil {
		return ctx.Status(http.StatusNotFound).SendString("Property not found")
	}

	return ctx.Render("property_detail", fiber.Map{
		"Property": property,
	})
}

// SearchProperties handles property search requests
func (c *PropertyController) SearchProperties(ctx *fiber.Ctx) error {
	// CORRECTED: The method is not .Queries, we need to use PeekMulti from the underlying context.
	qBytes := ctx.Context().QueryArgs().PeekMulti("q")
	queries := make([]string, len(qBytes))
	for i, b := range qBytes {
		queries[i] = string(b)
	}

	properties, err := c.propertyService.SearchProperties(queries)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(properties)
}

// GetProperty retrieves a single property as JSON
func (c *PropertyController) GetProperty(ctx *fiber.Ctx) error {
	id, err := getPropertyID(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	property, err := c.propertyService.GetPropertyByID(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	if property == nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Property not found"})
	}

	return ctx.JSON(property)
}

// CreateProperty creates a new property
func (c *PropertyController) CreateProperty(ctx *fiber.Ctx) error {
	log.Printf("[CreateProperty] Received request")

	var property models.Property
	if err := ctx.BodyParser(&property); err != nil {
		log.Printf("[CreateProperty] Error parsing property data: %v", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid property data",
			"details": err.Error(),
		})
	}

	log.Printf("[CreateProperty] Parsed property data: %+v", property)

	// Use the centralized validation function from the model
	if err := models.ValidateProperty(&property); err != nil {
		log.Printf("[CreateProperty] Validation error: %v", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid property data",
			"details": err.Error(),
		})
	}

	if err := c.propertyService.CreateProperty(&property); err != nil {
		log.Printf("[CreateProperty] Error creating property: %v", err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to create property",
			"details": err.Error(),
		})
	}

	log.Printf("[CreateProperty] Property created successfully: %+v", property)
	return ctx.Status(http.StatusCreated).JSON(fiber.Map{
		"message":  "Property created successfully",
		"property": property,
	})
}

// UpdateProperty updates an existing property
func (c *PropertyController) UpdateProperty(ctx *fiber.Ctx) error {
	id, err := getPropertyID(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var property models.Property
	if err := ctx.BodyParser(&property); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	// Use the centralized validation function from the model
	if err := models.ValidateProperty(&property); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid property data",
			"details": err.Error(),
		})
	}

	if err := c.propertyService.UpdateProperty(id, &property); err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(fiber.Map{"message": "Property updated successfully"})
}

// DeleteProperty deletes a property
func (c *PropertyController) DeleteProperty(ctx *fiber.Ctx) error {
	id, err := getPropertyID(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := c.propertyService.DeleteProperty(id); err != nil {
		log.Printf("Error deleting property: %v", err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error":   "Failed to delete property",
			"details": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{"message": "Property deleted successfully"})
}
