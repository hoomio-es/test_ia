package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/jgargallo/property-manager/models"
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

// Index renders the main properties page
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
	id, err := strconv.Atoi(ctx.Params("id"))
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
	query := ctx.Query("q")
	properties, err := c.propertyService.SearchProperties(query)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(properties)
}

// GetProperty retrieves a single property
func (c *PropertyController) GetProperty(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
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
	
	// Parse JSON body
	var property models.Property
	if err := ctx.BodyParser(&property); err != nil {
		log.Printf("[CreateProperty] Error parsing property data: %v", err)
		log.Printf("[CreateProperty] Raw body: %s", string(ctx.Body()))
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid property data",
			"details": err.Error(),
		})
	}

	log.Printf("[CreateProperty] Parsed property data: %+v", property)

	// Validar campos requeridos
	if property.NombreAlias == "" {
		log.Printf("[CreateProperty] Validation error: nombre_alias is empty")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid property data",
			"details": "nombre_alias is required",
		})
	}

	if property.DireccionCompleta == "" {
		log.Printf("[CreateProperty] Validation error: direccion_completa is empty")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid property data",
			"details": "direccion_completa is required",
		})
	}

	if property.TipoPropiedad == "" {
		log.Printf("[CreateProperty] Validation error: tipo_propiedad is empty")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid property data",
			"details": "tipo_propiedad is required",
		})
	}

	if property.CapacidadMaxima == "" {
		log.Printf("[CreateProperty] Validation error: capacidad_maxima is empty")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid property data",
			"details": "capacidad_maxima is required",
		})
	}

	// Convertir capacidad_maxima a int para validación
	capInt, err := strconv.Atoi(property.CapacidadMaxima)
	if err != nil {
		log.Printf("[CreateProperty] Error converting capacidad_maxima: %v", err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid property data",
			"details": "capacidad_maxima must be a valid number",
		})
	}

	if capInt <= 0 {
		log.Printf("[CreateProperty] Validation error: capacidad_maxima must be greater than 0")
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid property data",
			"details": "capacidad_maxima must be greater than 0",
		})
	}

	// La capacidad_maxima ya está en formato string, no necesitamos convertirla de vuelta
	log.Printf("[CreateProperty] Starting property creation in service")
	if err := c.propertyService.CreateProperty(&property); err != nil {
		log.Printf("[CreateProperty] Error creating property: %v", err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create property",
			"details": err.Error(),
		})
	}

	log.Printf("[CreateProperty] Property created successfully: %+v", property)
	return ctx.JSON(fiber.Map{
		"message": "Property created successfully",
		"property": property,
	})
}

// UpdateProperty updates an existing property
func (c *PropertyController) UpdateProperty(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var property models.Property
	if err := ctx.BodyParser(&property); err != nil {
		return ctx.Status(http.StatusBadRequest).SendString(err.Error())
	}

	if err := c.propertyService.UpdateProperty(id, &property); err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.JSON(fiber.Map{"message": "Property updated successfully"})
}

// DeleteProperty deletes a property
func (c *PropertyController) DeleteProperty(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := c.propertyService.DeleteProperty(id); err != nil {
		log.Printf("Error deleting property: %v", err)
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete property",
			"details": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{"message": "Property deleted successfully"})
}
