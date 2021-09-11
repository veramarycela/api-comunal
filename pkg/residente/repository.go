package residente

import "context"

// Repository handle the CRUD operations with Residentes.
type Repository interface {
	GetAll(ctx context.Context) ([]Residente, error)
	GetOne(ctx context.Context, id uint) (Residente, error)
	GetByResidentename(ctx context.Context, Residentename string) (Residente, error)
	Create(ctx context.Context, residente *Residente) error
	Update(ctx context.Context, id uint, residente Residente) error
	Delete(ctx context.Context, id uint) error
}
