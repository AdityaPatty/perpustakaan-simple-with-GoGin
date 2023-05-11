package servicebuku

// PACKAGE BUKU UNTUK FILE DB.GO
// BERISIKAN PERINTAH KONEKSI KE DATABASE

// YANG NANTINYA AKAN DI PANGGIL KE Func Main

// UNTUK

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type DatabaseConfig struct {
	Username string
	Password string
	Hostname string
	Port     string
	DBName   string
}

func (db DatabaseConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		db.Username, db.Password, db.Hostname, db.Port, db.DBName)
}

type Data struct {
	dbPool *pgxpool.Pool
}

func NewData(pool *pgxpool.Pool) Data {
	return Data{dbPool: pool}
}

func (ds Data) Pool() *pgxpool.Pool {
	return ds.dbPool
}

func NewDBPool(dbConfig DatabaseConfig) (*pgxpool.Pool, func(), error) {

	f := func() {}

	pool, err := pgxpool.Connect(context.Background(), dbConfig.DSN())

	if err != nil {
		return nil, f, errors.New("Koneksi ke Database Erorr")
	}

	err = validateDBPool(pool)

	if err != nil {
		return nil, f, err
	}

	return pool, func() { pool.Close() }, nil
}

func validateDBPool(pool *pgxpool.Pool) error {

	err := pool.Ping(context.Background())

	if err != nil {
		return errors.New("Koneksi ke Database Erorr")
	}

	var (
		currentDatabase string
		currentUser     string
		dbVersion       string
	)

	sqlStatement := `select current_database(), current_user, version();`
	row := pool.QueryRow(context.Background(), sqlStatement)
	err = row.Scan(&currentDatabase, &currentUser, &dbVersion)

	switch {
	case err == sql.ErrNoRows:
		return errors.New("no rows were returned")
	case err != nil:
		return errors.New("Koneksi ke Database Erorr")
	default:
		log.Printf("database version: %s\n", dbVersion)
		log.Printf("current database user: %s\n", currentUser)
		log.Printf("current database: %s\n", currentDatabase)
	}

	return nil
}
