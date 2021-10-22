package usecases

import (
	entities "github.com/nicolasolmos/agree-challenge/pokemon/entities"
	repositories "github.com/nicolasolmos/agree-challenge/pokemon/repositories"
)

func GetAllPokemons(paramRepository repositories.Repository) *[]entities.Pokemon {
	return paramRepository.SelectAll()
}
