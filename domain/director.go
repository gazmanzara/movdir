package domain

type Director struct {
	Id     int
	Name   string
	Gender int
}
type DirectorRepository interface {
	FindAll() ([]Director, error)
}
