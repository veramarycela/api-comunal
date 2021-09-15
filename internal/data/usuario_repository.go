package data

import (
	"context"

	"github.com/vermarycela/api-comunal/pkg/usuario"
)

type UsuarioRepository struct {
	Data *Data
}

func (ur *UsuarioRepository) GetAll(ctx context.Context) ([]usuario.Usuario, error) {
	q := `
    SELECT id, username, password, privilegios, rol, cedula_residente
        FROM usuario;
    `

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var usuarios []usuario.Usuario
	for rows.Next() {
		var u usuario.Usuario
		rows.Scan(&u.ID, &u.Username, &u.Password, &u.Privilegios,
			&u.Rol, &u.Cedula_residente)
		usuarios = append(usuarios, u)
	}

	return usuarios, nil
}

func (ur *UsuarioRepository) GetOne(ctx context.Context, id uint) (usuario.Usuario, error) {
	q := `
    SELECT id, username, password, privilegios, rol, cedula_residente
	FROM usuario WHERE id = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	var u usuario.Usuario
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Privilegios,
		&u.Rol, &u.Cedula_residente)
	if err != nil {
		return usuario.Usuario{}, err
	}

	return u, nil
}

func (ur *UsuarioRepository) GetByUsuarioname(ctx context.Context, usuarioname string) (usuario.Usuario, error) {
	q := `
    SELECT id, username, password, privilegios, rol, cedula_residente
        FROM usuario WHERE username = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, usuarioname)

	var u usuario.Usuario
	err := row.Scan(&u.ID, &u.Username, &u.Password, &u.Privilegios,
		&u.Rol, &u.Cedula_residente)
	if err != nil {
		return usuario.Usuario{}, err
	}

	return u, nil
}

func (ur *UsuarioRepository) Create(ctx context.Context, u *usuario.Usuario) error {
	q := `
    INSERT INTO usuario (id, username, password, privilegios, rol, cedula_residente)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id;
    `
	// if u.Picture == "" {
	// 	u.Picture = "https://placekitten.com/g/300/300"
	// }
	// if err := u.HashPassword(); err != nil {
	// 	return err
	// }
	row := ur.Data.DB.QueryRowContext(
		ctx, q, &u.ID, &u.Username, &u.Password, &u.Privilegios,
		&u.Rol, &u.Cedula_residente,
	)

	err := row.Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UsuarioRepository) Update(ctx context.Context, id uint, u usuario.Usuario) error {
	q := `
    UPDATE usuario set username=$1, password=$2, privilegios=$3, rol=$4, cedula_residente=$5
        WHERE id=$6;
    `

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, &u.ID, &u.Username, &u.Password, &u.Privilegios,
		&u.Rol, &u.Cedula_residente,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ur *UsuarioRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM usuario WHERE id=$1;`

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
