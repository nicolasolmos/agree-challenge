package pokemon

import entities "nicolas-olmos/agree-challenge/pokemon/entities"

type Repository interface {
	Insert(entities.Pokemon)
	Delete()
	SelectAll()
	SelectAndFilter()
	SelectById()
	Selectpage()
	UpdateAll()
}
