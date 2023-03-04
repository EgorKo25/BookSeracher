package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"BookSearcher/internal/config"
)

var (
	ErrConnectToDb  = errors.New("cannot connect to db")
	ErrCreatedTable = errors.New("cannot create table")
)

type Database struct {
	db  *sql.DB
	cfg *config.Config
}

func NewDatabase(cfg *config.Config) (*Database, error) {

	ctx := context.Background()

	db, err := sql.Open("pgx", cfg.DBAddr)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrConnectToDb, err)
	}

	err = createTable(ctx, db)
	if err != nil {
		return nil, fmt.Errorf("%s: %s", ErrCreatedTable, err)
	}

	return &Database{
		db:  db,
		cfg: cfg,
	}, nil
}

func createTable(ctx context.Context, db *sql.DB) error {

	childCtx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	_, err := db.ExecContext(childCtx,
		`CREATE TABLE 
    		IF NOT EXISTS books 
			(id SERIAL PRIMARY KEY,
			title VARCHAR(100),
			author VARCHAR(100));`)
	if err != nil {
		return err
	}

	return nil
}
