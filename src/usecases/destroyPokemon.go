package usecases

import (
	entities "github.com/nicolasolmos/agree-challenge/src/entities"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories"
)

func DestroyPokemonUsecase(paramId string, paramRepository repositories.Repository) {
	var newPokemon = entities.Pokemon{
		Id: paramId,
	}

	paramRepository.Delete(newPokemon)
}
