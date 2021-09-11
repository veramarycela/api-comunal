package documento

import "context"

// Repository handle the CRUD operations with Residentes.
type Repository interface {
	GetAll(ctx context.Context) ([]Documento, error)
	GetOne(ctx context.Context, id uint) (Documento, error)
	GetByResidentename(ctx context.Context, Documentoname string) (Documento, error)
	Create(ctx context.Context, documento *Documento) error
	Update(ctx context.Context, id uint, documento Documento) error
	Delete(ctx context.Context, id uint) error
}
