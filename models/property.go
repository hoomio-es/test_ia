package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

// Property model
type Property struct {
	ID              int     `json:"id"`
	NombreAlias      string  `json:"nombre_alias"`
	DireccionCompleta string  `json:"direccion_completa"`
	TipoPropiedad    string  `json:"tipo_propiedad"`
	CapacidadMaxima  string  `json:"capacidad_maxima"`
}

// PropertyService provides methods for property operations
type PropertyService struct {
	db *sql.DB
}

// NewPropertyService creates a new property service
func NewPropertyService(db *sql.DB) *PropertyService {
	return &PropertyService{db: db}
}

// GetAllProperties retrieves all properties from the database
func (s *PropertyService) GetAllProperties() ([]Property, error) {
	rows, err := s.db.Query("SELECT id, nombre_alias, direccion_completa, tipo_propiedad, capacidad_maxima FROM properties")
	if err != nil {
		return nil, fmt.Errorf("error retrieving properties: %w", err)
	}
	defer rows.Close()

	properties := []Property{}
	for rows.Next() {
		var p Property
		if err := rows.Scan(&p.ID, &p.NombreAlias, &p.DireccionCompleta, &p.TipoPropiedad, &p.CapacidadMaxima); err != nil {
			return nil, fmt.Errorf("error scanning property: %w", err)
		}
		properties = append(properties, p)
	}

	return properties, nil
}

// GetPropertyByID retrieves a property by its ID
func (s *PropertyService) GetPropertyByID(id int) (*Property, error) {
	var p Property
	row := s.db.QueryRow("SELECT id, nombre_alias, direccion_completa, tipo_propiedad, capacidad_maxima FROM properties WHERE id = $1", id)
	if err := row.Scan(&p.ID, &p.NombreAlias, &p.DireccionCompleta, &p.TipoPropiedad, &p.CapacidadMaxima); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error retrieving property: %w", err)
	}

	return &p, nil
}

// SearchProperties searches for properties matching the query
func (s *PropertyService) SearchProperties(query string) ([]Property, error) {
	searchTerm := "%" + query + "%"
	rows, err := s.db.Query("SELECT id, nombre_alias, direccion_completa, tipo_propiedad, capacidad_maxima FROM properties WHERE nombre_alias ILIKE $1 OR direccion_completa ILIKE $1 OR tipo_propiedad ILIKE $1", searchTerm)
	if err != nil {
		return nil, fmt.Errorf("error searching properties: %w", err)
	}
	defer rows.Close()

	properties := []Property{}
	for rows.Next() {
		var p Property
		if err := rows.Scan(&p.ID, &p.NombreAlias, &p.DireccionCompleta, &p.TipoPropiedad, &p.CapacidadMaxima); err != nil {
			return nil, fmt.Errorf("error scanning property: %w", err)
		}
		properties = append(properties, p)
	}

	return properties, nil
}

// CreateProperty creates a new property
func (s *PropertyService) CreateProperty(p *Property) error {
	log.Printf("[PropertyService] Starting property creation with data: %+v", p)
	
	// Convertir capacidad_maxima a int para la base de datos
	capInt, err := strconv.Atoi(p.CapacidadMaxima)
	if err != nil {
		log.Printf("[PropertyService] Error converting capacidad_maxima: %v", err)
		return fmt.Errorf("error converting capacidad_maxima: %w", err)
	}

	query := "INSERT INTO properties (nombre_alias, direccion_completa, tipo_propiedad, capacidad_maxima) VALUES ($1, $2, $3, $4)"
	log.Printf("[PropertyService] Executing query: %s", query)
	
	_, err = s.db.Exec(query,
		p.NombreAlias, p.DireccionCompleta, p.TipoPropiedad, capInt)
	if err != nil {
		log.Printf("[PropertyService] Error executing query: %v", err)
		return fmt.Errorf("error creating property: %w", err)
	}

	log.Printf("[PropertyService] Property created successfully")
	return nil
}

// UpdateProperty updates an existing property
func (s *PropertyService) UpdateProperty(id int, p *Property) error {
	_, err := s.db.Exec("UPDATE properties SET nombre_alias = $1, direccion_completa = $2, tipo_propiedad = $3, capacidad_maxima = $4 WHERE id = $5",
		p.NombreAlias, p.DireccionCompleta, p.TipoPropiedad, p.CapacidadMaxima, id)
	if err != nil {
		return fmt.Errorf("error updating property: %w", err)
	}

	return nil
}

// DeleteProperty deletes a property by its ID
func (s *PropertyService) DeleteProperty(id int) error {
	_, err := s.db.Exec("DELETE FROM properties WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("error deleting property: %w", err)
	}

	return nil
}
