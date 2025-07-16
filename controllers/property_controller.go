package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jgargallo/property-manager/models"
)

type PropertyController struct {
	propertyService *models.PropertyService
}

func NewPropertyController(propertyService *models.PropertyService) *PropertyController {
	return &PropertyController{
		propertyService: propertyService,
	}
}

func getPropertyID(ctx *fiber.Ctx) (int, error) {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (c *PropertyController) ListProperties(ctx *fiber.Ctx) error {
	properties, err := c.propertyService.GetAllProperties()
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).SendString(err.Error())
	}
	return ctx.Render("index", fiber.Map{"Properties": properties})
}

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
	return ctx.Render("property_detail", fiber.Map{"Property": property})
}

func (c *PropertyController) SearchProperties(ctx *fiber.Ctx) error {
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

func (c *PropertyController) GetProperty(ctx *fiber.Ctx) error {
	id, err := getPropertyID(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	property, err := c.propertyService.GetPropertyByID(id)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	if property == nil {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{"error": "Property not found"})
	}
	return ctx.JSON(property)
}

func (c *PropertyController) CreateProperty(ctx *fiber.Ctx) error {
	var property models.Property
	if err := ctx.BodyParser(&property); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid property data", "details": err.Error()})
	}
	if err := models.ValidateProperty(&property); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid property data", "details": err.Error()})
	}
	if err := c.propertyService.CreateProperty(&property); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create property", "details": err.Error()})
	}
	return ctx.Status(http.StatusCreated).JSON(property)
}

func (c *PropertyController) UpdateProperty(ctx *fiber.Ctx) error {
	id, err := getPropertyID(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var property models.Property
	if err := ctx.BodyParser(&property); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid property data", "details": err.Error()})
	}
	if err := models.ValidateProperty(&property); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid property data", "details": err.Error()})
	}
	if err := c.propertyService.UpdateProperty(id, &property); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update property", "details": err.Error()})
	}
	return ctx.JSON(fiber.Map{"message": "Property updated successfully"})
}

func (c *PropertyController) DeleteProperty(ctx *fiber.Ctx) error {
	id, err := getPropertyID(ctx)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := c.propertyService.DeleteProperty(id); err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete property", "details": err.Error()})
	}
	return ctx.JSON(fiber.Map{"message": "Property deleted successfully"})
}
