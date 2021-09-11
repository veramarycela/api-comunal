package solicitud

import "context"

// Repository handle the CRUD operations with Solicituds.
type Repository interface {
	GetAll(ctx context.Context) ([]Solicitud, error)
	GetOne(ctx context.Context, id uint) (Solicitud, error)
	GetBySolicitudname(ctx context.Context, Solicitudname string) (Solicitud, error)
	Create(ctx context.Context, solicitud *Solicitud) error
	Update(ctx context.Context, id uint, solicitud Solicitud) error
	Delete(ctx context.Context, id uint) error
}
