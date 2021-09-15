package data

import (
	"context"

	"github.com/vermarycela/api-comunal/pkg/solicitud"
)

type SolicitudRepository struct {
	Data *Data
}

func (ur *SolicitudRepository) GetAll(ctx context.Context) ([]solicitud.Solicitud, error) {
	q := `
    SELECT codigo, cedula_residente, codigo_documento, fecha        
        FROM solicitud;
    `

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var solicituds []solicitud.Solicitud
	for rows.Next() {
		var s solicitud.Solicitud
		rows.Scan(&s.Codigo, &s.Cedula_residente, &s.Codigo_documento, &s.Fecha)
		solicituds = append(solicituds, s)
	}

	return solicituds, nil
}

func (ur *SolicitudRepository) GetOne(ctx context.Context, id uint) (solicitud.Solicitud, error) {
	q := `
    SELECT codigo, cedula_residente, codigo_documento, fecha
	FROM solicitud WHERE codigo = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	var s solicitud.Solicitud
	err := row.Scan(s.Codigo, &s.Cedula_residente, &s.Codigo_documento, &s.Fecha)
	if err != nil {
		return solicitud.Solicitud{}, err
	}

	return s, nil
}

func (ur *SolicitudRepository) GetBySolicitudname(ctx context.Context, solicitudname string) (solicitud.Solicitud, error) {
	q := `
    SELECT codigo, cedula_residente, codigo_documento, fecha
        FROM solicitud WHERE fecha = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, solicitudname)

	var s solicitud.Solicitud
	err := row.Scan(s.Codigo, &s.Cedula_residente, &s.Codigo_documento, &s.Fecha)
	if err != nil {
		return solicitud.Solicitud{}, err
	}

	return s, nil
}

func (ur *SolicitudRepository) Create(ctx context.Context, s *solicitud.Solicitud) error {
	q := `
    INSERT INTO solicitud (codigo, cedula_residente, codigo_documento, fecha)
        VALUES ($1, $2, $3, $4)
        RETURNING codigo;
    `
	row := ur.Data.DB.QueryRowContext(
		ctx, q, &s.Codigo, &s.Cedula_residente, &s.Codigo_documento, &s.Fecha,
	)

	err := row.Scan(&s.Codigo)
	if err != nil {
		return err
	}

	return nil
}

func (ur *SolicitudRepository) Update(ctx context.Context, id uint, s solicitud.Solicitud) error {
	q := `
    UPDATE solicitud set codigo=$1, cedula_residente=$2, codigo_documento=$3, fecha=$4
        WHERE cedula=$5;
    `

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, &s.Codigo, &s.Cedula_residente, &s.Codigo_documento, &s.Fecha,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ur *SolicitudRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM solicitud WHERE codigo=$1;`

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
