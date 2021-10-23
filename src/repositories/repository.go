package repositories

import entities "github.com/nicolasolmos/agree-challenge/src/entities"

type Repository interface {
	Insert(entities.Pokemon) *entities.DatabaseError
	Delete(entities.Pokemon) *entities.DatabaseError
	SelectAll() (*[]entities.Pokemon, *entities.DatabaseError)
	SelectAndFilter() *[]entities.Pokemon
	SelectById(string) (entities.Pokemon, *entities.DatabaseError)
	Selectpage() *[]entities.Pokemon
	UpdateAll(entities.Pokemon) *entities.DatabaseError
}
