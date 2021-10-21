package pokemon

import entities "nicolas-olmos/agree-challenge/pokemon/entities"

type Repository interface {
	Insert(entities.Pokemon)
	Delete(entities.Pokemon)
	SelectAll()
	SelectAndFilter()
	SelectById()
	Selectpage()
	UpdateAll()
}
