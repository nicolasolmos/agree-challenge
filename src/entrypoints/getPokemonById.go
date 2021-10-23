package entrypoints

import (
	entities "github.com/nicolasolmos/agree-challenge/src/entities"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories/mysql"

	usecases "github.com/nicolasolmos/agree-challenge/src/usecases/getPokemon"
)

func GetPokemonEntrypoint(paramId string) (entities.Pokemon, error) {

	myRepository, error := repositories.NewPokemonRepository()

	if error != nil {
		return entities.Pokemon{}, error
	}

	return usecases.GetPokemonById(paramId, myRepository)
}
