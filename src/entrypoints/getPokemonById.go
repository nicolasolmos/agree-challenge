package entrypoints

import (
	entities "github.com/nicolasolmos/agree-challenge/src/entities"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories/mysql"

	usecases "github.com/nicolasolmos/agree-challenge/src/usecases/getPokemon"
)

func GetPokemonEntrypoint(paramId string) entities.Pokemon {

	var myRepository = repositories.NewPokemonRepository()

	return usecases.GetPokemonById(paramId, myRepository)

}
