package pokemon

import entities "nicolas-olmos/agree-challenge/pokemon/entities"

type Repository interface {
	Insert(entities.Pokemon)
	Delete(entities.Pokemon)
	SelectAll() *[]entities.Pokemon
	SelectAndFilter() *[]entities.Pokemon
	SelectById() *[]entities.Pokemon
	Selectpage() *[]entities.Pokemon
	UpdateAll(entities.Pokemon)
}