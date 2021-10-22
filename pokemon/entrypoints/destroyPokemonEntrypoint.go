package entrypoints

import (
	repositories "github.com/nicolasolmos/agree-challenge/pokemon/repositories/mysql"
	usecases "github.com/nicolasolmos/agree-challenge/pokemon/usecases"
)

func DestroyPokemonEntrypoint(paramId string) {
	var myPokemonRepository = repositories.NewPokemonRepository()
	usecases.DestroyPokemonUsecase(paramId, myPokemonRepository)
}
