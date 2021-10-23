package entrypoints

import (
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories/mysql"
	usecases "github.com/nicolasolmos/agree-challenge/src/usecases"
)

func DestroyPokemonEntrypoint(paramId string) error {
	myPokemonRepository, error := repositories.NewPokemonRepository()

	if error != nil {
		return error
	}

	return usecases.DestroyPokemonUsecase(paramId, myPokemonRepository)
}
