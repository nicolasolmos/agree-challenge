package repositories

import entities "github.com/nicolasolmos/agree-challenge/src/entities"

type Repository interface {
	Insert(entities.Pokemon) error
	Delete(entities.Pokemon) error
	SelectAll() (*[]entities.Pokemon, error)
	SelectAndFilter() *[]entities.Pokemon
	SelectById(string) (entities.Pokemon, error)
	Selectpage() *[]entities.Pokemon
	UpdateAll(entities.Pokemon) error
}
