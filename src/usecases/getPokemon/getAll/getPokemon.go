package usecases

import (
	entities "github.com/nicolasolmos/agree-challenge/src/entities"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories"
)

func GetAllPokemons(paramRepository repositories.Repository) *[]entities.Pokemon {
	return paramRepository.SelectAll()
}
