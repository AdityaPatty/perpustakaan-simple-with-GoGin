package repositorybuku

import (
	"context"
	"perpustakaan/modelbuku"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Ping(context.Context) error
	Close()
}

type Database struct {
	DB PgxIface
}

func NewDatabase(ds PgxIface) Database {
	return Database{DB: ds}
}

func (pool Database) Create(perpustakaan modelbuku.Perpustakaan) (*modelbuku.Perpustakaan, error) {

	query := `INSERT INTO perpustakaans (judulbuku,deskripsibuku,isbn,issn,bahasabuku)
          VALUES ($1,$2,$3,$4,$5) RETURNING id,judulbuku,deskripsibuku,isbn,issn,bahasabuku`

	row := pool.DB.QueryRow(context.Background(), query,
		perpustakaan.Judulbuku, perpustakaan.Deskripsibuku, perpustakaan.Isbn, perpustakaan.Issn, perpustakaan.Bahasabuku)

	buku := new(modelbuku.Perpustakaan)

	err := row.Scan(
		&buku.ID,
		&buku.Judulbuku,
		&buku.Deskripsibuku,
		&buku.Isbn,
		&buku.Issn,
		&buku.Bahasabuku,
	)

	if err != nil {
		return nil, err
	}

	return buku, nil
}

func (pool Database) Get(id int) (*modelbuku.Perpustakaan, error) {
	query := `SELECT * FROM perpustakaans WHERE id = $1`

	row := pool.DB.QueryRow(context.Background(), query, id)

	buku := new(modelbuku.Perpustakaan)

	err := row.Scan(
		&buku.ID,
		&buku.Judulbuku,
		&buku.Deskripsibuku,
		&buku.Isbn,
		&buku.Issn,
		&buku.Bahasabuku,
	)

	if err != nil {
		return nil, err
	}

	return buku, nil
}

func (pool Database) Gets() ([]*modelbuku.Perpustakaan, error) {
	query := `SELECT * FROM perpustakaans`

	rows, err := pool.DB.Query(context.Background(), query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var perpustakaans []*modelbuku.Perpustakaan
	if rows != nil {
		for rows.Next() {
			buku := new(modelbuku.Perpustakaan)

			err := rows.Scan(
				&buku.ID,
				&buku.Judulbuku,
				&buku.Deskripsibuku,
				&buku.Isbn,
				&buku.Issn,
				&buku.Bahasabuku,
			)

			if err != nil {
				return nil, err
			}

			perpustakaans = append(perpustakaans, buku)
		}
	}

	return perpustakaans, nil
}

func (pool Database) Update(id int, perpustakaan modelbuku.Perpustakaan) (*modelbuku.Perpustakaan, error) {
	query := `UPDATE perpustakaans SET 
            judulbuku = $2,
            deksripsibuku  = $3,
            isbn = $4,
            issn = $5,
			bahasabuku = $6
          WHERE id = $1
          RETURNING id, judubuku, deskripsibuku, isbn, issn, bahasabuku;
         `
	row := pool.DB.QueryRow(context.Background(), query, id,
		perpustakaan.Judulbuku, perpustakaan.Deskripsibuku, perpustakaan.Isbn, perpustakaan.Issn, perpustakaan.Bahasabuku)

	buku := new(modelbuku.Perpustakaan)

	if err := row.Scan(
		&buku.ID,
		&buku.Judulbuku,
		&buku.Deskripsibuku,
		&buku.Isbn,
		&buku.Issn,
		&buku.Bahasabuku,
	); err != nil {
		return nil, err
	}

	return buku, nil
}

func (pool Database) Delete(id int) (*modelbuku.Perpustakaan, error) {
	q := `DELETE FROM perpustakaans WHERE id = $1 RETURNING id,judulbuku,deskripsibuku,isbn,issn,bahasabuku;`

	row := pool.DB.QueryRow(context.Background(), q, id)

	buku := new(modelbuku.Perpustakaan)

	if err := row.Scan(
		&buku.ID,
		&buku.Judulbuku,
		&buku.Deskripsibuku,
		&buku.Isbn,
		&buku.Issn,
		&buku.Bahasabuku,
	); err != nil {
		return nil, err
	}

	return buku, nil

}
