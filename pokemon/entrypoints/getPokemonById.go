package entrypoints

import (
	entities "github.com/nicolasolmos/agree-challenge/pokemon/entities"
	repositories "github.com/nicolasolmos/agree-challenge/pokemon/repositories/mysql"

	usecases "github.com/nicolasolmos/agree-challenge/pokemon/usecases/getPokemon"
)

func GetPokemonEntrypoint(paramId string) entities.Pokemon {

	var myRepository = repositories.NewPokemonRepository()

	return usecases.GetPokemonById(paramId, myRepository)

}
