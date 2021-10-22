package repositories

import entities "github.com/nicolasolmos/agree-challenge/src/entities"

type Repository interface {
	Insert(entities.Pokemon)
	Delete(entities.Pokemon)
	SelectAll() *[]entities.Pokemon
	SelectAndFilter() *[]entities.Pokemon
	SelectById(string) entities.Pokemon
	Selectpage() *[]entities.Pokemon
	UpdateAll(entities.Pokemon)
}
