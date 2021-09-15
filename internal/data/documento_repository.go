package data

import (
	"context"

	"github.com/vermarycela/api-comunal/pkg/documento"
)

type DocumentoRepository struct {
	Data *Data
}

func (ur *DocumentoRepository) GetAll(ctx context.Context) ([]documento.Documento, error) {
	q := `
    SELECT codigo, tipo, dirigido, contenido, autorizado
        FROM documento;
    `

	rows, err := ur.Data.DB.QueryContext(ctx, q)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var documentos []documento.Documento
	for rows.Next() {
		var d documento.Documento
		rows.Scan(&d.Codigo, &d.Tipo, &d.Dirigido,
			&d.Contenido, &d.Tipo)
		documentos = append(documentos, d)
	}

	return documentos, nil
}

func (ur *DocumentoRepository) GetOne(ctx context.Context, id uint) (documento.Documento, error) {
	q := `
    SELECT codigo, tipo, dirigido, contenido, autorizado
	FROM documento WHERE codigo = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, id)

	var d documento.Documento
	err := row.Scan(&d.Codigo, &d.Tipo, &d.Dirigido, &d.Contenido, &d.Tipo)
	if err != nil {
		return documento.Documento{}, err
	}

	return d, nil
}

func (ur *DocumentoRepository) GetByDocumentoname(ctx context.Context, documentoname string) (documento.Documento, error) {
	q := `
    SELECT codigo, tipo, dirigido, contenido, autorizado
        FROM documento WHERE tipo = $1;
    `

	row := ur.Data.DB.QueryRowContext(ctx, q, documentoname)

	var d documento.Documento
	err := row.Scan(&d.Codigo, &d.Tipo, &d.Dirigido, &d.Contenido, &d.Tipo)
	if err != nil {
		return documento.Documento{}, err
	}

	return d, nil
}

func (ur *DocumentoRepository) Create(ctx context.Context, d *documento.Documento) error {
	q := `
    INSERT INTO documento (codigo, tipo, dirigido, contenido, autorizado)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING codigo;
    `

	// if u.Picture == "" {
	// 	u.Picture = "https://placekitten.com/g/300/300"
	// }

	// if err := u.HashPassword(); err != nil {
	// 	return err
	// }

	row := ur.Data.DB.QueryRowContext(
		ctx, q, &d.Codigo, &d.Tipo, &d.Dirigido, &d.Contenido, &d.Tipo,
	)

	err := row.Scan(&d.Codigo)
	if err != nil {
		return err
	}

	return nil
}

func (ur *DocumentoRepository) Update(ctx context.Context, id uint, d documento.Documento) error {
	q := `
    UPDATE documento set tipo=$1, dirigido=$2, contenido=$3, autorizado=$4
        WHERE codigo=$7;
    `

	stmt, err := ur.Data.DB.PrepareContext(ctx, q)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(
		ctx, &d.Codigo, &d.Tipo, &d.Dirigido, &d.Contenido, &d.Tipo,
	)
	if err != nil {
		return err
	}

	return nil
}

func (ur *DocumentoRepository) Delete(ctx context.Context, id uint) error {
	q := `DELETE FROM documento WHERE codigo=$1;`

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
