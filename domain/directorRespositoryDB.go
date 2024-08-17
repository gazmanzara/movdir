package domain

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gazmanzara/movdir/errs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
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
		println(err.Error())
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

func (d DirectorRepositoryDB) Save(director Director) (*Director, *errs.AppError) {
	query := `INSERT INTO directors (id, name, gender) VALUES (?, ?, ?);`
	exec, err := d.db.Exec(query, director.Id, director.Name, director.Gender)
	if err != nil {
		return nil, errs.NewInternalServerError("Unexpected error while saving the director")
	}

	_, err = exec.LastInsertId()
	if err != nil {
		return nil, errs.NewInternalServerError("Unexpected error while getting the director id")
	}

	return &director, nil
}

func NewDirectorRepositoryDB() DirectorRepositoryDB {
	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Open("mysql", dataSource)
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
