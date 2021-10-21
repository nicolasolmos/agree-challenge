package pokemon

import (
	entities "nicolas-olmos/agree-challenge/pokemon/entities"
	repositories "nicolas-olmos/agree-challenge/pokemon/repositories"
)

func DestroyPokemonUsecase(paramId string, paramRepository repositories.Repository) {
	var newPokemon = entities.Pokemon{
		Id: paramId,
	}

	paramRepository.Delete(newPokemon)
}
