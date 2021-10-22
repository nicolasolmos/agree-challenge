package usecases

import (
	entities "github.com/nicolasolmos/agree-challenge/src/entities"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories"
)

func GetPokemonById(paramId string, paramRepository repositories.Repository) entities.Pokemon {
	return paramRepository.SelectById(paramId)
}
