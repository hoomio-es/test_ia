package models

import (
	"database/sql"
	"time"
)

type Issue struct {
	ID          int
	PropertyID  int
	Description string
	Status      string
	Updates     string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type IssueService struct {
	DB *sql.DB
}

func NewIssueService(db *sql.DB) *IssueService {
	return &IssueService{DB: db}
}

// Get all issues for a property
func (s *IssueService) GetIssuesByProperty(propertyID int) ([]Issue, error) {
	rows, err := s.DB.Query(`SELECT id, property_id, description, status, updates, created_at, updated_at FROM issues WHERE property_id = $1 ORDER BY id DESC`, propertyID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var issues []Issue
	for rows.Next() {
		var issue Issue
		if err := rows.Scan(&issue.ID, &issue.PropertyID, &issue.Description, &issue.Status, &issue.Updates, &issue.CreatedAt, &issue.UpdatedAt); err != nil {
			return nil, err
		}
		issues = append(issues, issue)
	}
	return issues, nil
}

// Create a new issue
func (s *IssueService) CreateIssue(issue *Issue) error {
	err := s.DB.QueryRow(
		`INSERT INTO issues (property_id, description, status, updates) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`,
		issue.PropertyID, issue.Description, issue.Status, issue.Updates,
	).Scan(&issue.ID, &issue.CreatedAt, &issue.UpdatedAt)
	return err
}
