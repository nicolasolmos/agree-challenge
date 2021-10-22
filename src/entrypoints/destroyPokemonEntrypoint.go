package entrypoints

import (
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories/mysql"
	usecases "github.com/nicolasolmos/agree-challenge/src/usecases"
)

func DestroyPokemonEntrypoint(paramId string) {
	var myPokemonRepository = repositories.NewPokemonRepository()
	usecases.DestroyPokemonUsecase(paramId, myPokemonRepository)
}
