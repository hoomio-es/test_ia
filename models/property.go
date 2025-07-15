package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var db *sql.DB

// InitDB initializes the database connection
func InitDB() (*sql.DB, error) {
	dsn := "postgresql://postgres:postgres@db:5432/test_app?sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	// Create tables if they don't exist
	createTables := `
	CREATE TABLE IF NOT EXISTS properties (
	    id SERIAL PRIMARY KEY,
	    nombre_alias VARCHAR(255) NOT NULL,
	    cedula_habitabilidad VARCHAR(255),
	    licencia_turistica VARCHAR(255),
	    direccion_completa TEXT NOT NULL,
	    tipo_propiedad VARCHAR(50) NOT NULL,
	    capacidad_maxima INTEGER NOT NULL,
	    numero_habitaciones INTEGER NOT NULL,
	    numero_banos INTEGER NOT NULL,
	    metros_cuadrados INTEGER NOT NULL,
	    descripcion_corta TEXT,
	    amenidades TEXT,
	    codigo_airbnb VARCHAR(255),
	    url_airbnb TEXT,
	    codigo_booking VARCHAR(255),
	    url_booking TEXT,
	    precio_base_noches DECIMAL(10,2) NOT NULL,
	    ibi_anual DECIMAL(10,2),
	    gastos_comunidad DECIMAL(10,2),
	    hipoteca_mensual DECIMAL(10,2),
	    coste_limpieza DECIMAL(10,2),
	    contacto_propietario_nombre VARCHAR(255),
	    contacto_propietario_telefono VARCHAR(255),
	    contacto_propietario_email VARCHAR(255),
	    limpiador_nombre VARCHAR(255),
	    limpiador_telefono VARCHAR(255),
	    fontanero_nombre VARCHAR(255),
	    fontanero_telefono VARCHAR(255),
	    seguro_hogar_aseguradora VARCHAR(255),
	    seguro_hogar_poliza VARCHAR(255),
	    seguro_hogar_fecha_venc TIMESTAMP,
	    seguro_hogar_telefono VARCHAR(255),
	    seguro_hogar_email VARCHAR(255),
	    seguro_rc_aseguradora VARCHAR(255),
	    seguro_rc_poliza VARCHAR(255),
	    seguro_rc_fecha_venc TIMESTAMP,
	    seguro_rc_telefono VARCHAR(255),
	    seguro_rc_email VARCHAR(255),
	    contrato_electricidad_empresa VARCHAR(255),
	    contrato_electricidad_contrato VARCHAR(255),
	    contrato_electricidad_fecha_venc TIMESTAMP,
	    contrato_electricidad_telefono VARCHAR(255),
	    contrato_agua_empresa VARCHAR(255),
	    contrato_agua_contrato VARCHAR(255),
	    contrato_agua_fecha_venc TIMESTAMP,
	    contrato_agua_telefono VARCHAR(255),
	    contrato_internet_empresa VARCHAR(255),
	    contrato_internet_contrato VARCHAR(255),
	    contrato_internet_fecha_venc TIMESTAMP,
	    contrato_internet_telefono VARCHAR(255),
	    contrato_gas_empresa VARCHAR(255),
	    contrato_gas_contrato VARCHAR(255),
	    contrato_gas_fecha_venc TIMESTAMP,
	    contrato_gas_telefono VARCHAR(255),
	    notas_adicionales TEXT,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	if _, err := db.Exec(createTables); err != nil {
		return nil, fmt.Errorf("error creating tables: %w", err)
	}

	return db, nil
}

// Property model
type Property struct {
	ID                    int            `json:"id"`
	NombreAlias           string         `json:"nombre_alias"`
	CedulaHabitabilidad   sql.NullString `json:"cedula_habitabilidad"`
	LicenciaTuristica     sql.NullString `json:"licencia_turistica"`
	DireccionCompleta     string         `json:"direccion_completa"`
	TipoPropiedad         string         `json:"tipo_propiedad"`
	CapacidadMaxima       int            `json:"capacidad_maxima"`
	NumeroHabitaciones    int            `json:"numero_habitaciones"`
	NumeroBanos           int            `json:"numero_banos"`
	MetrosCuadrados       int            `json:"metros_cuadrados"`
	DescripcionCorta      sql.NullString `json:"descripcion_corta"`
	Amenidades           sql.NullString `json:"amenidades"`
	CodigoAirbnb         sql.NullString `json:"codigo_airbnb"`
	UrlAirbnb            sql.NullString `json:"url_airbnb"`
	CodigoBooking        sql.NullString `json:"codigo_booking"`
	UrlBooking           sql.NullString `json:"url_booking"`
	PrecioBaseNoches     float64        `json:"precio_base_noches"`
	IbiAnual             sql.NullFloat64 `json:"ibi_anual"`
	GastosComunidad      sql.NullFloat64 `json:"gastos_comunidad"`
	HipotecaMensual      sql.NullFloat64 `json:"hipoteca_mensual"`
	CosteLimpieza        sql.NullFloat64 `json:"coste_limpieza"`
	ContactoPropietarioNombre sql.NullString `json:"contacto_propietario_nombre"`
	ContactoPropietarioTelefono sql.NullString `json:"contacto_propietario_telefono"`
	ContactoPropietarioEmail sql.NullString `json:"contacto_propietario_email"`
	LimpiadorNombre      sql.NullString `json:"limpiador_nombre"`
	LimpiadorTelefono    sql.NullString `json:"limpiador_telefono"`
	FontaneroNombre      sql.NullString `json:"fontanero_nombre"`
	FontaneroTelefono    sql.NullString `json:"fontanero_telefono"`
	SeguroHogarAseguradora sql.NullString `json:"seguro_hogar_aseguradora"`
	SeguroHogarPoliza    sql.NullString `json:"seguro_hogar_poliza"`
	SeguroHogarFechaVenc sql.NullTime   `json:"seguro_hogar_fecha_venc"`
	SeguroHogarTelefono  sql.NullString `json:"seguro_hogar_telefono"`
	SeguroHogarEmail     sql.NullString `json:"seguro_hogar_email"`
	SeguroRCAseguradora  sql.NullString `json:"seguro_rc_aseguradora"`
	SeguroRCPoliza       sql.NullString `json:"seguro_rc_poliza"`
	SeguroRCFechaVenc    sql.NullTime   `json:"seguro_rc_fecha_venc"`
	SeguroRCTelefono     sql.NullString `json:"seguro_rc_telefono"`
	SeguroRCEmail        sql.NullString `json:"seguro_rc_email"`
	ContratoElectricidadEmpresa sql.NullString `json:"contrato_electricidad_empresa"`
	ContratoElectricidadContrato sql.NullString `json:"contrato_electricidad_contrato"`
	ContratoElectricidadFechaVenc sql.NullTime   `json:"contrato_electricidad_fecha_venc"`
	ContratoElectricidadTelefono sql.NullString `json:"contrato_electricidad_telefono"`
	ContratoAguaEmpresa   sql.NullString `json:"contrato_agua_empresa"`
	ContratoAguaContrato  sql.NullString `json:"contrato_agua_contrato"`
	ContratoAguaFechaVenc sql.NullTime   `json:"contrato_agua_fecha_venc"`
	ContratoAguaTelefono  sql.NullString `json:"contrato_agua_telefono"`
	ContratoInternetEmpresa sql.NullString `json:"contrato_internet_empresa"`
	ContratoInternetContrato sql.NullString `json:"contrato_internet_contrato"`
	ContratoInternetFechaVenc sql.NullTime   `json:"contrato_internet_fecha_venc"`
	ContratoInternetTelefono sql.NullString `json:"contrato_internet_telefono"`
	ContratoGasEmpresa    sql.NullString `json:"contrato_gas_empresa"`
	ContratoGasContrato   sql.NullString `json:"contrato_gas_contrato"`
	ContratoGasFechaVenc  sql.NullTime   `json:"contrato_gas_fecha_venc"`
	ContratoGasTelefono   sql.NullString `json:"contrato_gas_telefono"`
	NotasAdicionales     sql.NullString `json:"notas_adicionales"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
}

// ValidateProperty valida los campos requeridos y las restricciones de la propiedad
func ValidateProperty(p *Property) error {
	if p.NombreAlias == "" {
		return fmt.Errorf("nombre_alias es requerido")
	}
	if p.DireccionCompleta == "" {
		return fmt.Errorf("direccion_completa es requerido")
	}
	if p.TipoPropiedad == "" {
		return fmt.Errorf("tipo_propiedad es requerido")
	}
	if p.CapacidadMaxima <= 0 {
		return fmt.Errorf("capacidad_maxima debe ser mayor que 0")
	}
	if p.NumeroHabitaciones <= 0 {
		return fmt.Errorf("numero_habitaciones debe ser mayor que 0")
	}
	if p.NumeroBanos <= 0 {
		return fmt.Errorf("numero_banos debe ser mayor que 0")
	}
	if p.MetrosCuadrados <= 0 {
		return fmt.Errorf("metros_cuadrados debe ser mayor que 0")
	}
	if p.PrecioBaseNoches <= 0 {
		return fmt.Errorf("precio_base_noches debe ser mayor que 0")
	}
	return nil
}

// PropertyService provides methods for property operations
type PropertyService struct {
	db *sql.DB
}

// NewPropertyService creates a new property service
func NewPropertyService(db *sql.DB) *PropertyService {
	return &PropertyService{db: db}
}

// toSQLValues converts a Property to a slice of SQL values
func toSQLValues(p *Property) []interface{} {
	values := make([]interface{}, 0)
	values = append(values,
		p.NombreAlias,
		p.CedulaHabitabilidad,
		p.LicenciaTuristica,
		p.DireccionCompleta,
		p.TipoPropiedad,
		p.CapacidadMaxima,
		p.NumeroHabitaciones,
		p.NumeroBanos,
		p.MetrosCuadrados,
		p.DescripcionCorta,
		p.Amenidades,
		p.CodigoAirbnb,
		p.UrlAirbnb,
		p.CodigoBooking,
		p.UrlBooking,
		p.PrecioBaseNoches,
		p.IbiAnual,
		p.GastosComunidad,
		p.HipotecaMensual,
		p.CosteLimpieza,
		p.ContactoPropietarioNombre,
		p.ContactoPropietarioTelefono,
		p.ContactoPropietarioEmail,
		p.LimpiadorNombre,
		p.LimpiadorTelefono,
		p.FontaneroNombre,
		p.FontaneroTelefono,
		p.SeguroHogarAseguradora,
		p.SeguroHogarPoliza,
		p.SeguroHogarFechaVenc,
		p.SeguroHogarTelefono,
		p.SeguroHogarEmail,
		p.SeguroRCAseguradora,
		p.SeguroRCPoliza,
		p.SeguroRCFechaVenc,
		p.SeguroRCTelefono,
		p.SeguroRCEmail,
		p.ContratoElectricidadEmpresa,
		p.ContratoElectricidadContrato,
		p.ContratoElectricidadFechaVenc,
		p.ContratoElectricidadTelefono,
		p.ContratoAguaEmpresa,
		p.ContratoAguaContrato,
		p.ContratoAguaFechaVenc,
		p.ContratoAguaTelefono,
		p.ContratoInternetEmpresa,
		p.ContratoInternetContrato,
		p.ContratoInternetFechaVenc,
		p.ContratoInternetTelefono,
		p.ContratoGasEmpresa,
		p.ContratoGasContrato,
		p.ContratoGasFechaVenc,
		p.ContratoGasTelefono,
		p.NotasAdicionales,
	)

	return values
}

// GetAllProperties retrieves all properties from the database
func (s *PropertyService) GetAllProperties() ([]Property, error) {
	query := "SELECT * FROM properties ORDER BY created_at DESC"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error retrieving properties: %w", err)
	}
	defer rows.Close()

	properties := []Property{}
	for rows.Next() {
		var p Property
		if err := rows.Scan(
			&p.ID,
			&p.NombreAlias,
			&p.CedulaHabitabilidad,
			&p.LicenciaTuristica,
			&p.DireccionCompleta,
			&p.TipoPropiedad,
			&p.CapacidadMaxima,
			&p.NumeroHabitaciones,
			&p.NumeroBanos,
			&p.MetrosCuadrados,
			&p.DescripcionCorta,
			&p.Amenidades,
			&p.CodigoAirbnb,
			&p.UrlAirbnb,
			&p.CodigoBooking,
			&p.UrlBooking,
			&p.PrecioBaseNoches,
			&p.IbiAnual,
			&p.GastosComunidad,
			&p.HipotecaMensual,
			&p.CosteLimpieza,
			&p.ContactoPropietarioNombre,
			&p.ContactoPropietarioTelefono,
			&p.ContactoPropietarioEmail,
			&p.LimpiadorNombre,
			&p.LimpiadorTelefono,
			&p.FontaneroNombre,
			&p.FontaneroTelefono,
			&p.SeguroHogarAseguradora,
			&p.SeguroHogarPoliza,
			&p.SeguroHogarFechaVenc,
			&p.SeguroHogarTelefono,
			&p.SeguroHogarEmail,
			&p.SeguroRCAseguradora,
			&p.SeguroRCPoliza,
			&p.SeguroRCFechaVenc,
			&p.SeguroRCTelefono,
			&p.SeguroRCEmail,
			&p.ContratoElectricidadEmpresa,
			&p.ContratoElectricidadContrato,
			&p.ContratoElectricidadFechaVenc,
			&p.ContratoElectricidadTelefono,
			&p.ContratoAguaEmpresa,
			&p.ContratoAguaContrato,
			&p.ContratoAguaFechaVenc,
			&p.ContratoAguaTelefono,
			&p.ContratoInternetEmpresa,
			&p.ContratoInternetContrato,
			&p.ContratoInternetFechaVenc,
			&p.ContratoInternetTelefono,
			&p.ContratoGasEmpresa,
			&p.ContratoGasContrato,
			&p.ContratoGasFechaVenc,
			&p.ContratoGasTelefono,
			&p.NotasAdicionales,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning property: %w", err)
		}
		properties = append(properties, p)
	}

	return properties, nil
}

// GetPropertyByID retrieves a property by its ID
func (s *PropertyService) GetPropertyByID(id int) (*Property, error) {
	query := "SELECT * FROM properties WHERE id = $1"
	var p Property
	row := s.db.QueryRow(query, id)
	if err := row.Scan(
		&p.ID,
		&p.NombreAlias,
		&p.CedulaHabitabilidad,
		&p.LicenciaTuristica,
		&p.DireccionCompleta,
		&p.TipoPropiedad,
		&p.CapacidadMaxima,
		&p.NumeroHabitaciones,
		&p.NumeroBanos,
		&p.MetrosCuadrados,
		&p.DescripcionCorta,
		&p.Amenidades,
		&p.CodigoAirbnb,
		&p.UrlAirbnb,
		&p.CodigoBooking,
		&p.UrlBooking,
		&p.PrecioBaseNoches,
		&p.IbiAnual,
		&p.GastosComunidad,
		&p.HipotecaMensual,
		&p.CosteLimpieza,
		&p.ContactoPropietarioNombre,
		&p.ContactoPropietarioTelefono,
		&p.ContactoPropietarioEmail,
		&p.LimpiadorNombre,
		&p.LimpiadorTelefono,
		&p.FontaneroNombre,
		&p.FontaneroTelefono,
		&p.SeguroHogarAseguradora,
		&p.SeguroHogarPoliza,
		&p.SeguroHogarFechaVenc,
		&p.SeguroHogarTelefono,
		&p.SeguroHogarEmail,
		&p.SeguroRCAseguradora,
		&p.SeguroRCPoliza,
		&p.SeguroRCFechaVenc,
		&p.SeguroRCTelefono,
		&p.SeguroRCEmail,
		&p.ContratoElectricidadEmpresa,
		&p.ContratoElectricidadContrato,
		&p.ContratoElectricidadFechaVenc,
		&p.ContratoElectricidadTelefono,
		&p.ContratoAguaEmpresa,
		&p.ContratoAguaContrato,
		&p.ContratoAguaFechaVenc,
		&p.ContratoAguaTelefono,
		&p.ContratoInternetEmpresa,
		&p.ContratoInternetContrato,
		&p.ContratoInternetFechaVenc,
		&p.ContratoInternetTelefono,
		&p.ContratoGasEmpresa,
		&p.ContratoGasContrato,
		&p.ContratoGasFechaVenc,
		&p.ContratoGasTelefono,
		&p.NotasAdicionales,
		&p.CreatedAt,
		&p.UpdatedAt,
	); err != nil {
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
	query := "SELECT * FROM properties WHERE nombre_alias ILIKE $1 OR direccion_completa ILIKE $1 OR tipo_propiedad ILIKE $1"
	rows, err := s.db.Query(query, searchTerm)
	if err != nil {
		return nil, fmt.Errorf("error searching properties: %w", err)
	}
	defer rows.Close()

	properties := []Property{}
	for rows.Next() {
		var p Property
		if err := rows.Scan(
			&p.ID,
			&p.NombreAlias,
			&p.CedulaHabitabilidad,
			&p.LicenciaTuristica,
			&p.DireccionCompleta,
			&p.TipoPropiedad,
			&p.CapacidadMaxima,
			&p.NumeroHabitaciones,
			&p.NumeroBanos,
			&p.MetrosCuadrados,
			&p.DescripcionCorta,
			&p.Amenidades,
			&p.CodigoAirbnb,
			&p.UrlAirbnb,
			&p.CodigoBooking,
			&p.UrlBooking,
			&p.PrecioBaseNoches,
			&p.IbiAnual,
			&p.GastosComunidad,
			&p.HipotecaMensual,
			&p.CosteLimpieza,
			&p.ContactoPropietarioNombre,
			&p.ContactoPropietarioTelefono,
			&p.ContactoPropietarioEmail,
			&p.LimpiadorNombre,
			&p.LimpiadorTelefono,
			&p.FontaneroNombre,
			&p.FontaneroTelefono,
			&p.SeguroHogarAseguradora,
			&p.SeguroHogarPoliza,
			&p.SeguroHogarFechaVenc,
			&p.SeguroHogarTelefono,
			&p.SeguroHogarEmail,
			&p.SeguroRCAseguradora,
			&p.SeguroRCPoliza,
			&p.SeguroRCFechaVenc,
			&p.SeguroRCTelefono,
			&p.SeguroRCEmail,
			&p.ContratoElectricidadEmpresa,
			&p.ContratoElectricidadContrato,
			&p.ContratoElectricidadFechaVenc,
			&p.ContratoElectricidadTelefono,
			&p.ContratoAguaEmpresa,
			&p.ContratoAguaContrato,
			&p.ContratoAguaFechaVenc,
			&p.ContratoAguaTelefono,
			&p.ContratoInternetEmpresa,
			&p.ContratoInternetContrato,
			&p.ContratoInternetFechaVenc,
			&p.ContratoInternetTelefono,
			&p.ContratoGasEmpresa,
			&p.ContratoGasContrato,
			&p.ContratoGasFechaVenc,
			&p.ContratoGasTelefono,
			&p.NotasAdicionales,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("error scanning property: %w", err)
		}
		properties = append(properties, p)
	}

	return properties, nil
}

// CreateProperty creates a new property
func (s *PropertyService) CreateProperty(p *Property) error {
	if err := ValidateProperty(p); err != nil {
		return fmt.Errorf("invalid property data: %w", err)
	}

	query := "INSERT INTO properties (nombre_alias, cedula_habitabilidad, licencia_turistica, direccion_completa, tipo_propiedad, capacidad_maxima, numero_habitaciones, numero_banos, metros_cuadrados, descripcion_corta, amenidades, codigo_airbnb, url_airbnb, codigo_booking, url_booking, precio_base_noches, ibi_anual, gastos_comunidad, hipoteca_mensual, coste_limpieza, contacto_propietario_nombre, contacto_propietario_telefono, contacto_propietario_email, limpiador_nombre, limpiador_telefono, fontanero_nombre, fontanero_telefono, seguro_hogar_aseguradora, seguro_hogar_poliza, seguro_hogar_fecha_venc, seguro_hogar_telefono, seguro_hogar_email, seguro_rc_aseguradora, seguro_rc_poliza, seguro_rc_fecha_venc, seguro_rc_telefono, seguro_rc_email, contrato_electricidad_empresa, contrato_electricidad_contrato, contrato_electricidad_fecha_venc, contrato_electricidad_telefono, contrato_agua_empresa, contrato_agua_contrato, contrato_agua_fecha_venc, contrato_agua_telefono, contrato_internet_empresa, contrato_internet_contrato, contrato_internet_fecha_venc, contrato_internet_telefono, contrato_gas_empresa, contrato_gas_contrato, contrato_gas_fecha_venc, contrato_gas_telefono, notas_adicionales) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36, $37, $38, $39, $40, $41, $42, $43, $44, $45, $46, $47, $48, $49, $50, $51, $52, $53, $54, $55, $56, $57, $58, $59, $60, $61, $62, $63) RETURNING id"

	values := toSQLValues(p)
	var id int
	if err := s.db.QueryRow(query, values...).Scan(&id); err != nil {
		return fmt.Errorf("error inserting property: %w", err)
	}

	p.ID = id
	return nil
}

// UpdateProperty updates an existing property
func (s *PropertyService) UpdateProperty(id int, p *Property) error {
	if err := ValidateProperty(p); err != nil {
		return fmt.Errorf("invalid property data: %w", err)
	}

	query := "UPDATE properties SET nombre_alias = $1, cedula_habitabilidad = $2, licencia_turistica = $3, direccion_completa = $4, tipo_propiedad = $5, capacidad_maxima = $6, numero_habitaciones = $7, numero_banos = $8, metros_cuadrados = $9, descripcion_corta = $10, amenidades = $11, codigo_airbnb = $12, url_airbnb = $13, codigo_booking = $14, url_booking = $15, precio_base_noches = $16, ibi_anual = $17, gastos_comunidad = $18, hipoteca_mensual = $19, coste_limpieza = $20, contacto_propietario_nombre = $21, contacto_propietario_telefono = $22, contacto_propietario_email = $23, limpiador_nombre = $24, limpiador_telefono = $25, fontanero_nombre = $26, fontanero_telefono = $27, seguro_hogar_aseguradora = $28, seguro_hogar_poliza = $29, seguro_hogar_fecha_venc = $30, seguro_hogar_telefono = $31, seguro_hogar_email = $32, seguro_rc_aseguradora = $33, seguro_rc_poliza = $34, seguro_rc_fecha_venc = $35, seguro_rc_telefono = $36, seguro_rc_email = $37, contrato_electricidad_empresa = $38, contrato_electricidad_contrato = $39, contrato_electricidad_fecha_venc = $40, contrato_electricidad_telefono = $41, contrato_agua_empresa = $42, contrato_agua_contrato = $43, contrato_agua_fecha_venc = $44, contrato_agua_telefono = $45, contrato_internet_empresa = $46, contrato_internet_contrato = $47, contrato_internet_fecha_venc = $48, contrato_internet_telefono = $49, contrato_gas_empresa = $50, contrato_gas_contrato = $51, contrato_gas_fecha_venc = $52, contrato_gas_telefono = $53, notas_adicionales = $54 WHERE id = $55"

	values := toSQLValues(p)
	values = append(values, id)
	_, err := s.db.Exec(query, values...)
	return err
}

// DeleteProperty deletes a property by its ID
func (s *PropertyService) DeleteProperty(id int) error {
	query := "DELETE FROM properties WHERE id = $1"
	_, err := s.db.Exec(query, id)
	return err
}
