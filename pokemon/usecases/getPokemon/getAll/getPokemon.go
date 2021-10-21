package pokemon

import (
	entities "nicolas-olmos/agree-challenge/pokemon/entities"
	repositories "nicolas-olmos/agree-challenge/pokemon/repositories"
)

func GetAllPokemons(paramRepository repositories.Repository) *[]entities.Pokemon {
	return paramRepository.SelectAll()
}
