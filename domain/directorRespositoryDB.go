package domain

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gazmanzara/movdir/app/errs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

type DirectorRepositoryDB struct {
	db  *sqlx.DB
	ctx context.Context
}

func (d DirectorRepositoryDB) FindAll() ([]Director, *errs.AppError) {
	var directors []Director

	query := `SELECT * FROM directors;`
	err := d.db.Select(&directors, query)
	if err != nil {
		return nil, errs.NewInternalServerError("Unexpected database error!")
	}

	return directors, nil
}

func (d DirectorRepositoryDB) FindById(id string) (*Director, *errs.AppError) {
	var director Director

	query := `SELECT * FROM directors WHERE id = ?;`
	err := d.db.Get(&director, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("Director not found")
		} else {
			println(err.Error())
			return nil, errs.NewInternalServerError("Unexpected error while querying the database")
		}
	}

	return &director, nil
}

func NewDirectorRepositoryDB() DirectorRepositoryDB {
	db, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/movdir")
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return DirectorRepositoryDB{
		db: db,
	}
}
