package usuario

// User of the system.
type Usuario struct {
	ID               string `json:"id,omitempty"`
	Username         string `json:"username,omitempty"`
	Password         string `json:"password,omitempty"`
	Privilegios      string `json:"privilegios,omitempty"`
	Rol              string `json:"rol,omitempty"`
	Cedula_residente string `json:"cedula_residente,omitempty"`
}
