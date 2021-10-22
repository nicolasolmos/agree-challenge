package pokemon

import (
	entities "github.com/nicolasolmos/agree-challenge/pokemons/entities"
	repositories "github.com/nicolasolmos/agree-challenge/pokemons/repositories"
)

func GetPokemonById(paramId string, paramRepository repositories.Repository) entities.Pokemon {
	return paramRepository.SelectById(paramId)
}
