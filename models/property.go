package models

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"
)

// Property model reflects the full schema
type Property struct {
	ID                            int             `json:"id"`
	NombreAlias                   string          `json:"nombre_alias"`
	CedulaHabitabilidad           sql.NullString  `json:"cedula_habitabilidad"`
	LicenciaTuristica             sql.NullString  `json:"licencia_turistica"`
	DireccionCompleta             string          `json:"direccion_completa"`
	TipoPropiedad                 string          `json:"tipo_propiedad"`
	CapacidadMaxima               int             `json:"capacidad_maxima"`
	NumeroHabitaciones            sql.NullInt64   `json:"numero_habitaciones"`
	NumeroBanos                   sql.NullInt64   `json:"numero_banos"`
	MetrosCuadrados               sql.NullInt64   `json:"metros_cuadrados"`
	DescripcionCorta              sql.NullString  `json:"descripcion_corta"`
	Amenidades                    sql.NullString  `json:"amenidades"`
	CodigoAirbnb                  sql.NullString  `json:"codigo_airbnb"`
	UrlAirbnb                     sql.NullString  `json:"url_airbnb"`
	CodigoBooking                 sql.NullString  `json:"codigo_booking"`
	UrlBooking                    sql.NullString  `json:"url_booking"`
	PrecioBaseNoches              sql.NullFloat64 `json:"precio_base_noches"`
	IbiAnual                      sql.NullFloat64 `json:"ibi_anual"`
	GastosComunidad               sql.NullFloat64 `json:"gastos_comunidad"`
	HipotecaMensual               sql.NullFloat64 `json:"hipoteca_mensual"`
	CosteLimpieza                 sql.NullFloat64 `json:"coste_limpieza"`
	ContactoPropietarioNombre     sql.NullString  `json:"contacto_propietario_nombre"`
	ContactoPropietarioTelefono   sql.NullString  `json:"contacto_propietario_telefono"`
	ContactoPropietarioEmail      sql.NullString  `json:"contacto_propietario_email"`
	LimpiadorNombre               sql.NullString  `json:"limpiador_nombre"`
	LimpiadorTelefono             sql.NullString  `json:"limpiador_telefono"`
	FontaneroNombre               sql.NullString  `json:"fontanero_nombre"`
	FontaneroTelefono             sql.NullString  `json:"fontanero_telefono"`
	SeguroHogarAseguradora        sql.NullString  `json:"seguro_hogar_aseguradora"`
	SeguroHogarPoliza             sql.NullString  `json:"seguro_hogar_poliza"`
	SeguroHogarFechaVenc          sql.NullTime    `json:"seguro_hogar_fecha_venc"`
	SeguroHogarTelefono           sql.NullString  `json:"seguro_hogar_telefono"`
	SeguroHogarEmail              sql.NullString  `json:"seguro_hogar_email"`
	SeguroRcAseguradora           sql.NullString  `json:"seguro_rc_aseguradora"`
	SeguroRcPoliza                sql.NullString  `json:"seguro_rc_poliza"`
	SeguroRcFechaVenc             sql.NullTime    `json:"seguro_rc_fecha_venc"`
	SeguroRcTelefono              sql.NullString  `json:"seguro_rc_telefono"`
	SeguroRcEmail                 sql.NullString  `json:"seguro_rc_email"`
	ContratoElectricidadEmpresa   sql.NullString  `json:"contrato_electricidad_empresa"`
	ContratoElectricidadContrato  sql.NullString  `json:"contrato_electricidad_contrato"`
	ContratoElectricidadFechaVenc sql.NullTime    `json:"contrato_electricidad_fecha_venc"`
	ContratoElectricidadTelefono  sql.NullString  `json:"contrato_electricidad_telefono"`
	ContratoAguaEmpresa           sql.NullString  `json:"contrato_agua_empresa"`
	ContratoAguaContrato          sql.NullString  `json:"contrato_agua_contrato"`
	ContratoAguaFechaVenc         sql.NullTime    `json:"contrato_agua_fecha_venc"`
	ContratoAguaTelefono          sql.NullString  `json:"contrato_agua_telefono"`
	ContratoInternetEmpresa       sql.NullString  `json:"contrato_internet_empresa"`
	ContratoInternetContrato      sql.NullString  `json:"contrato_internet_contrato"`
	ContratoInternetFechaVenc     sql.NullTime    `json:"contrato_internet_fecha_venc"`
	ContratoInternetTelefono      sql.NullString  `json:"contrato_internet_telefono"`
	ContratoGasEmpresa            sql.NullString  `json:"contrato_gas_empresa"`
	ContratoGasContrato           sql.NullString  `json:"contrato_gas_contrato"`
	ContratoGasFechaVenc          sql.NullTime    `json:"contrato_gas_fecha_venc"`
	ContratoGasTelefono           sql.NullString  `json:"contrato_gas_telefono"`
	NotasAdicionales              sql.NullString  `json:"notas_adicionales"`
	CreatedAt                     time.Time       `json:"created_at"`
	UpdatedAt                     time.Time       `json:"updated_at"`
}

// PropertyService provides methods for property operations
type PropertyService struct {
	db *sql.DB
}

// NewPropertyService creates a new property service
func NewPropertyService(db *sql.DB) *PropertyService {
	return &PropertyService{db: db}
}

// rowScanner defines an interface that can be satisfied by *sql.Row and *sql.Rows.
type rowScanner interface {
	Scan(dest ...interface{}) error
}

// scanProperty is a helper function to scan a row into the Property struct.
func scanProperty(row rowScanner, p *Property) error {
	return row.Scan(
		&p.ID, &p.NombreAlias, &p.CedulaHabitabilidad, &p.LicenciaTuristica, &p.DireccionCompleta,
		&p.TipoPropiedad, &p.CapacidadMaxima, &p.NumeroHabitaciones, &p.NumeroBanos, &p.MetrosCuadrados,
		&p.DescripcionCorta, &p.Amenidades, &p.CodigoAirbnb, &p.UrlAirbnb, &p.CodigoBooking,
		&p.UrlBooking, &p.PrecioBaseNoches, &p.IbiAnual, &p.GastosComunidad, &p.HipotecaMensual,
		&p.CosteLimpieza, &p.ContactoPropietarioNombre, &p.ContactoPropietarioTelefono,
		&p.ContactoPropietarioEmail, &p.LimpiadorNombre, &p.LimpiadorTelefono, &p.FontaneroNombre,
		&p.FontaneroTelefono, &p.SeguroHogarAseguradora, &p.SeguroHogarPoliza, &p.SeguroHogarFechaVenc,
		&p.SeguroHogarTelefono, &p.SeguroHogarEmail, &p.SeguroRcAseguradora, &p.SeguroRcPoliza,
		&p.SeguroRcFechaVenc, &p.SeguroRcTelefono, &p.SeguroRcEmail, &p.ContratoElectricidadEmpresa,
		&p.ContratoElectricidadContrato, &p.ContratoElectricidadFechaVenc, &p.ContratoElectricidadTelefono,
		&p.ContratoAguaEmpresa, &p.ContratoAguaContrato, &p.ContratoAguaFechaVenc, &p.ContratoAguaTelefono,
		&p.ContratoInternetEmpresa, &p.ContratoInternetContrato, &p.ContratoInternetFechaVenc,
		&p.ContratoInternetTelefono, &p.ContratoGasEmpresa, &p.ContratoGasContrato,
		&p.ContratoGasFechaVenc, &p.ContratoGasTelefono, &p.NotasAdicionales, &p.CreatedAt, &p.UpdatedAt,
	)
}

// GetAllProperties retrieves all properties from the database
func (s *PropertyService) GetAllProperties() ([]Property, error) {
	rows, err := s.db.Query("SELECT * FROM properties ORDER BY id ASC")
	if err != nil {
		return nil, fmt.Errorf("error retrieving properties: %w", err)
	}
	defer rows.Close()

	var properties []Property
	for rows.Next() {
		var p Property
		if err := scanProperty(rows, &p); err != nil {
			return nil, fmt.Errorf("error scanning property: %w", err)
		}
		properties = append(properties, p)
	}

	return properties, nil
}

// GetPropertyByID retrieves a property by its ID
func (s *PropertyService) GetPropertyByID(id int) (*Property, error) {
	var p Property
	row := s.db.QueryRow("SELECT * FROM properties WHERE id = $1", id)
	if err := scanProperty(row, &p); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found is not an application error
		}
		return nil, fmt.Errorf("error retrieving property: %w", err)
	}

	return &p, nil
}

// SearchProperties searches for properties matching ALL of the query terms across multiple fields.
func (s *PropertyService) SearchProperties(queries []string) ([]Property, error) {
	if len(queries) == 0 {
		return s.GetAllProperties()
	}

	var conditions []string
	var args []interface{}
	paramIndex := 1

	// Lista de todos los campos de texto en los que se quiere buscar.
	searchFields := []string{
		"nombre_alias", "direccion_completa", "tipo_propiedad",
		"cedula_habitabilidad", "licencia_turistica", "descripcion_corta",
		"amenidades", "codigo_airbnb", "codigo_booking",
		"contacto_propietario_nombre", "contacto_propietario_email",
		"limpiador_nombre", "fontanero_nombre", "seguro_hogar_aseguradora",
		"seguro_hogar_poliza", "seguro_rc_aseguradora", "seguro_rc_poliza",
		"notas_adicionales",
	}

	for _, q := range queries {
		if q == "" {
			continue
		}

		// Para cada término de búsqueda, crea un grupo de condiciones OR.
		var orConditions []string
		for _, field := range searchFields {
			orConditions = append(orConditions, fmt.Sprintf("%s ILIKE $%d", field, paramIndex))
		}
		conditions = append(conditions, "("+strings.Join(orConditions, " OR ")+")")

		// Añade el término de búsqueda a los argumentos de la consulta.
		args = append(args, "%"+q+"%")
		paramIndex++
	}

	if len(conditions) == 0 {
		return s.GetAllProperties()
	}

	// Une todos los grupos de condiciones con AND para una búsqueda estricta.
	query := "SELECT * FROM properties WHERE " + strings.Join(conditions, " AND ") + " ORDER BY id ASC"

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error searching properties: %w", err)
	}
	defer rows.Close()

	var properties []Property
	for rows.Next() {
		var p Property
		if err := scanProperty(rows, &p); err != nil {
			return nil, fmt.Errorf("error scanning property: %w", err)
		}
		properties = append(properties, p)
	}
	return properties, nil
}

// toSQLValues converts a Property struct to a slice of interfaces for DB operations.
func toSQLValues(p *Property) []interface{} {
	return []interface{}{
		p.NombreAlias, p.CedulaHabitabilidad, p.LicenciaTuristica, p.DireccionCompleta, p.TipoPropiedad,
		p.CapacidadMaxima, p.NumeroHabitaciones, p.NumeroBanos, p.MetrosCuadrados, p.DescripcionCorta,
		p.Amenidades, p.CodigoAirbnb, p.UrlAirbnb, p.CodigoBooking, p.UrlBooking, p.PrecioBaseNoches,
		p.IbiAnual, p.GastosComunidad, p.HipotecaMensual, p.CosteLimpieza, p.ContactoPropietarioNombre,
		p.ContactoPropietarioTelefono, p.ContactoPropietarioEmail, p.LimpiadorNombre, p.LimpiadorTelefono,
		p.FontaneroNombre, p.FontaneroTelefono, p.SeguroHogarAseguradora, p.SeguroHogarPoliza,
		p.SeguroHogarFechaVenc, p.SeguroHogarTelefono, p.SeguroHogarEmail, p.SeguroRcAseguradora,
		p.SeguroRcPoliza, p.SeguroRcFechaVenc, p.SeguroRcTelefono, p.SeguroRcEmail,
		p.ContratoElectricidadEmpresa, p.ContratoElectricidadContrato, p.ContratoElectricidadFechaVenc,
		p.ContratoElectricidadTelefono, p.ContratoAguaEmpresa, p.ContratoAguaContrato,
		p.ContratoAguaFechaVenc, p.ContratoAguaTelefono, p.ContratoInternetEmpresa,
		p.ContratoInternetContrato, p.ContratoInternetFechaVenc, p.ContratoInternetTelefono,
		p.ContratoGasEmpresa, p.ContratoGasContrato, p.ContratoGasFechaVenc, p.ContratoGasTelefono,
		p.NotasAdicionales,
	}
}

// CreateProperty creates a new property
func (s *PropertyService) CreateProperty(p *Property) error {
	query := `
        INSERT INTO properties (
            nombre_alias, cedula_habitabilidad, licencia_turistica, direccion_completa, tipo_propiedad,
            capacidad_maxima, numero_habitaciones, numero_banos, metros_cuadrados, descripcion_corta,
            amenidades, codigo_airbnb, url_airbnb, codigo_booking, url_booking, precio_base_noches,
            ibi_anual, gastos_comunidad, hipoteca_mensual, coste_limpieza, contacto_propietario_nombre,
            contacto_propietario_telefono, contacto_propietario_email, limpiador_nombre, limpiador_telefono,
            fontanero_nombre, fontanero_telefono, seguro_hogar_aseguradora, seguro_hogar_poliza,
            seguro_hogar_fecha_venc, seguro_hogar_telefono, seguro_hogar_email, seguro_rc_aseguradora,
            seguro_rc_poliza, seguro_rc_fecha_venc, seguro_rc_telefono, seguro_rc_email,
            contrato_electricidad_empresa, contrato_electricidad_contrato, contrato_electricidad_fecha_venc,
            contrato_electricidad_telefono, contrato_agua_empresa, contrato_agua_contrato,
            contrato_agua_fecha_venc, contrato_agua_telefono, contrato_internet_empresa,
            contrato_internet_contrato, contrato_internet_fecha_venc, contrato_internet_telefono,
            contrato_gas_empresa, contrato_gas_contrato, contrato_gas_fecha_venc, contrato_gas_telefono,
            notas_adicionales
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20,
            $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38,
            $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54
        )
    `
	_, err := s.db.Exec(query, toSQLValues(p)...)
	if err != nil {
		log.Printf("Error executing create query: %v", err)
		return fmt.Errorf("error creating property: %w", err)
	}
	return nil
}

// UpdateProperty updates an existing property
func (s *PropertyService) UpdateProperty(id int, p *Property) error {
	query := `
        UPDATE properties SET
            nombre_alias = $1, cedula_habitabilidad = $2, licencia_turistica = $3, direccion_completa = $4,
            tipo_propiedad = $5, capacidad_maxima = $6, numero_habitaciones = $7, numero_banos = $8,
            metros_cuadrados = $9, descripcion_corta = $10, amenidades = $11, codigo_airbnb = $12,
            url_airbnb = $13, codigo_booking = $14, url_booking = $15, precio_base_noches = $16,
            ibi_anual = $17, gastos_comunidad = $18, hipoteca_mensual = $19, coste_limpieza = $20,
            contacto_propietario_nombre = $21, contacto_propietario_telefono = $22,
            contacto_propietario_email = $23, limpiador_nombre = $24, limpiador_telefono = $25,
            fontanero_nombre = $26, fontanero_telefono = $27, seguro_hogar_aseguradora = $28,
            seguro_hogar_poliza = $29, seguro_hogar_fecha_venc = $30, seguro_hogar_telefono = $31,
            seguro_hogar_email = $32, seguro_rc_aseguradora = $33, seguro_rc_poliza = $34,
            seguro_rc_fecha_venc = $35, seguro_rc_telefono = $36, seguro_rc_email = $37,
            contrato_electricidad_empresa = $38, contrato_electricidad_contrato = $39,
            contrato_electricidad_fecha_venc = $40, contrato_electricidad_telefono = $41,
            contrato_agua_empresa = $42, contrato_agua_contrato = $43, contrato_agua_fecha_venc = $44,
            contrato_agua_telefono = $45, contrato_internet_empresa = $46, contrato_internet_contrato = $47,
            contrato_internet_fecha_venc = $48, contrato_internet_telefono = $49, contrato_gas_empresa = $50,
            contrato_gas_contrato = $51, contrato_gas_fecha_venc = $52, contrato_gas_telefono = $53,
            notas_adicionales = $54, updated_at = NOW()
        WHERE id = $55
    `
	args := append(toSQLValues(p), id)
	_, err := s.db.Exec(query, args...)
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

// ValidateProperty checks if the property data is valid
func ValidateProperty(p *Property) error {
	if p.NombreAlias == "" {
		return fmt.Errorf("nombre_alias is required")
	}
	if p.DireccionCompleta == "" {
		return fmt.Errorf("direccion_completa is required")
	}
	if p.TipoPropiedad == "" {
		return fmt.Errorf("tipo_propiedad is required")
	}
	if p.CapacidadMaxima <= 0 {
		return fmt.Errorf("capacidad_maxima must be greater than 0")
	}
	return nil
}
