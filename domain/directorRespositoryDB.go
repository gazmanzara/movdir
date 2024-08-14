package domain

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gazmanzara/movdir/app/errs"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type DirectorRepositoryDB struct {
	db  *sql.DB
	ctx context.Context
}

func (d DirectorRepositoryDB) FindAll() ([]Director, *errs.AppError) {
	query := `SELECT * FROM directors;`

	rows, err := d.db.Query(query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("No directors found")
		} else {
			println(err.Error())
			return nil, errs.NewInternalServerError("Unexpected error while querying the database")
		}
	}
	defer rows.Close()

	var directors []Director

	for rows.Next() {
		director := Director{}

		err := rows.Scan(
			&director.Id,
			&director.Name,
			&director.Gender,
		)

		if err != nil {
			return nil, errs.NewInternalServerError("Error while scanning directors")
		}

		directors = append(directors, director)
	}

	return directors, nil
}

func (d DirectorRepositoryDB) FindById(id string) (*Director, *errs.AppError) {
	query := `SELECT * FROM directors WHERE id = ?;`

	row := d.db.QueryRow(query, id)

	var director Director

	err := row.Scan(
		&director.Id,
		&director.Name,
		&director.Gender,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("Director not found")
		} else {
			return nil, errs.NewInternalServerError("Unexpected error while querying the database")
		}
	}

	return &director, nil
}

func NewDirectorRepositoryDB() DirectorRepositoryDB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/movdir")
	if err != nil {
		panic(err.Error())
	}

	ctx := context.Background()

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return DirectorRepositoryDB{
		db:  db,
		ctx: ctx,
	}
}
