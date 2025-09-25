package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/jgargallo/property-manager/models"
)

type IssueController struct {
	Service *models.IssueService
}

func NewIssueController(service *models.IssueService) *IssueController {
	return &IssueController{Service: service}
}

func (ic *IssueController) ListIssues(c *fiber.Ctx) error {
	propertyID, _ := strconv.Atoi(c.Params("id"))
	issues, err := ic.Service.GetIssuesByProperty(propertyID)
	if err != nil {
		return c.Status(500).SendString("Error fetching issues")
	}
	return c.Render("issues", fiber.Map{
		"PropertyID": propertyID,
		"Issues":     issues,
	})
}

func (ic *IssueController) CreateIssue(c *fiber.Ctx) error {
	var issue models.Issue
	if err := c.BodyParser(&issue); err != nil {
		return c.Status(400).SendString("Invalid input")
	}
	if err := ic.Service.CreateIssue(&issue); err != nil {
		return c.Status(500).SendString("Error creating issue")
	}
	return c.Redirect("/property/" + strconv.Itoa(issue.PropertyID) + "/issues")
}
