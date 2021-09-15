package data

import (
	"context"

	"github.com/vermarycela/api-comunal/pkg/residente"
)

type ResidenteRepository struct {
	Data *Data
}

func (ur *ResidenteRepository) GetAll(ctx context.Context) ([]residente.Residente, error) {
	q := `
    SELECT cedula, nombre, apellido, fecha_nac, direccion, telefono,
        tipo
        FROM residente;
    `

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var residentes []residente.Residente
	for rows.Next() {
		var r residente.Residente
		rows.Scan(&r.Cedula, &r.Nombre, &r.Apellido, &r.Direccion,
			&r.Fecha_Nac, &r.Telefono, &r.Tipo)
		residentes = append(residentes, r)
	}

	return residentes, nil
}

func (ur *ResidenteRepository) GetOne(ctx context.Context, id uint) (residente.Residente, error) {
	q := `
    SELECT cedula, nombre, apellido, fecha_nac, direccion, telefono, tipo
	FROM residente WHERE cedula = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	var r residente.Residente
	err := row.Scan(&r.Cedula, &r.Nombre, &r.Apellido, &r.Fecha_Nac, &r.Direccion, &r.Telefono,
		&r.Tipo)
	if err != nil {
		return residente.Residente{}, err
	}

	return r, nil
}

func (ur *ResidenteRepository) GetByResidentename(ctx context.Context, residentename string) (residente.Residente, error) {
	q := `
    SELECT cedula, nombre, apellido, fecha_nac, direccion, telefono, tipo
        FROM residente WHERE nombre = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, residentename)

	var r residente.Residente
	err := row.Scan(&r.Cedula, &r.Nombre, &r.Apellido, &r.Fecha_Nac, &r.Direccion, &r.Telefono,
		&r.Tipo)
	if err != nil {
		return residente.Residente{}, err
	}

	return r, nil
}

func (ur *ResidenteRepository) Create(ctx context.Context, r *residente.Residente) error {
	q := `
    INSERT INTO residente (cedula, nombre, apellido, fecha_nac, direccion, telefono, tipo)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING cedula;
    `

	// if u.Picture == "" {
	// 	u.Picture = "https://placekitten.com/g/300/300"
	// }

	// if err := u.HashPassword(); err != nil {
	// 	return err
	// }

	row := ur.Data.DB.QueryRowContext(
		ctx, q, &r.Cedula, &r.Nombre, &r.Apellido, &r.Fecha_Nac, &r.Direccion, &r.Telefono,
		&r.Tipo,
	)

	err := row.Scan(&r.Cedula)
	if err != nil {
		return err
	}

	return nil
}

func (ur *ResidenteRepository) Update(ctx context.Context, id uint, r residente.Residente) error {
	q := `
    UPDATE residente set nombre=$1, apellido=$2, fecha_nac=$3, direccion=$4, telefono=$5, tipo:$6
        WHERE cedula=$7;
    `

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, &r.Cedula, &r.Nombre, &r.Apellido, &r.Fecha_Nac, &r.Direccion, &r.Telefono,
		&r.Tipo,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ur *ResidenteRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM residente WHERE cedula=$1;`

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
