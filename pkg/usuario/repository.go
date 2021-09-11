package usuario

import "context"

// Repository handle the CRUD operations with Usuarios.
type Repository interface {
	GetAll(ctx context.Context) ([]Usuario, error)
	GetOne(ctx context.Context, id uint) (Usuario, error)
	GetByUsuarioname(ctx context.Context, Usuarioname string) (Usuario, error)
	Create(ctx context.Context, usuario *Usuario) error
	Update(ctx context.Context, id uint, usuario Usuario) error
	Delete(ctx context.Context, id uint) error
}
