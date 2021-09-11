package documento

// User of the system.
type Documento struct {
	Codigo       string `json:"codigo,omitempty"`
	Tipo         string `json:"tipo,omitempty"`
	Dirigido     string `json:"dirigido,omitempty"`
	Contenido    string `json:"contenido,omitempty"`
	Autorizacion string `json:"autorizacion,omitempty"`
}
