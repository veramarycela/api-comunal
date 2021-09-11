package residente

// User of the system.
type Residente struct {
	Cedula    string `json:"cedula,omitempty"`
	Nombre    string `json:"nommbre,omitempty"`
	Apellido  string `json:"apellido,omitempty"`
	Fecha_Nac string `json:"fecha_nac,omitempty"`
	Direccion string `json:"direccion,omitempty"`
	Telefono  string `json:"telefono,omitempty"`
	Tipo      string `json:"tipo,omitempty"`
}
