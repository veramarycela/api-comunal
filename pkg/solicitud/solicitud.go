package solicitud

// User of the system.
type Solicitud struct {
	Codigo           string `json:"codigo,omitempty"`
	Cedula_residente string `json:"cedula_residente,omitempty"`
	Codigo_documento string `json:"codigo_documento,omitempty"`
	Fecha            string `json:"fecha,omitempty"`
}
