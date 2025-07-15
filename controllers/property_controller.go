package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jgargallo/test_app/models"
)

// PropertyController handles property-related routes
type PropertyController struct {
	propertyService *models.PropertyService
	engine          *html.Engine
}

// NewPropertyController creates a new property controller
func NewPropertyController(propertyService *models.PropertyService, engine *html.Engine) *PropertyController {
	return &PropertyController{
		propertyService: propertyService,
		engine:         engine,
	}
}

// PropertyIDMiddleware middleware para validar y extraer el ID de una propiedad
func (c *PropertyController) PropertyIDMiddleware(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	property, err := c.propertyService.GetPropertyByID(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error fetching property"})
	}

	if property == nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Property not found"})
	}

	ctx.Locals("property", property)
	return ctx.Next()
}

// Index renders the main properties page
// ListProperties renders the properties list page
func (c *PropertyController) ListProperties(ctx *fiber.Ctx) error {
	properties, err := c.propertyService.GetAllProperties()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Render("index", fiber.Map{
		"Properties": properties,
	})
}

// ShowProperty renders the property details page
func (c *PropertyController) ShowProperty(ctx *fiber.Ctx) error {
	property := ctx.Locals("property").(*models.Property)
	return ctx.Render("property_detail", fiber.Map{
		"Property": property,
	})
}

// SearchProperties handles property search requests
func (c *PropertyController) SearchProperties(ctx *fiber.Ctx) error {
	query := ctx.Query("q")
	properties, err := c.propertyService.SearchProperties(query)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(properties)
}

// GetProperty retrieves a single property
func (c *PropertyController) GetProperty(ctx *fiber.Ctx) error {
	property := ctx.Locals("property").(*models.Property)
	return ctx.JSON(property)
}

// CreateProperty creates a new property
func (c *PropertyController) CreateProperty(ctx *fiber.Ctx) error {
	var property models.Property
	if err := ctx.BodyParser(&property); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid property data",
			"details": err.Error(),
		})
	}

	// Validate property data
	if err := models.ValidateProperty(&property); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid property data",
			"details": err.Error(),
		})
	}

	// Create property
	if err := c.propertyService.CreateProperty(&property); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Error creating property",
			"details": err.Error(),
		})
	}

	// Return success response
	return ctx.JSON(fiber.Map{
		"message": "Property created successfully",
		"property": property,
	})
}

// UpdateProperty updates an existing property
func (c *PropertyController) UpdateProperty(ctx *fiber.Ctx) error {
	var property models.Property
	if err := ctx.BodyParser(&property); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid property data"})
	}

	// Validate property data
	if err := models.ValidateProperty(&property); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid property data",
			"details": err.Error(),
		})
	}

	// Get property ID from locals (middleware already validated it)
	existingProperty := ctx.Locals("property").(*models.Property)
	property.ID = existingProperty.ID

	// Update property
	if err := c.propertyService.UpdateProperty(property.ID, &property); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Property updated successfully"})
}

// DeleteProperty deletes a property
func (c *PropertyController) DeleteProperty(ctx *fiber.Ctx) error {
	// Get property ID from locals (middleware already validated it)
	property := ctx.Locals("property").(*models.Property)

	if err := c.propertyService.DeleteProperty(property.ID); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.JSON(fiber.Map{"message": "Property deleted successfully"})
}
