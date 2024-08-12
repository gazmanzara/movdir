package domain

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type DirectorRepositoryDB struct {
	db  *sql.DB
	ctx context.Context
}

func (d DirectorRepositoryDB) FindAll() ([]Director, error) {
	defer d.db.Close()
	query := `SELECT * FROM directors;`

	rows, err := d.db.QueryContext(d.ctx, query)
	defer rows.Close()

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	var directors []Director

	for rows.Next() {
		director := Director{}

		err := rows.Scan(
			&director.Id,
			&director.Name,
			&director.Gender,
		)

		if err != nil {
			return nil, err
		}

		directors = append(directors, director)
	}

	return directors, nil
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
