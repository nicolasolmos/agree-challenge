package entrypoints

import (
	dtos "github.com/nicolasolmos/agree-challenge/src/DTOS"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories/mysql"
	usecases "github.com/nicolasolmos/agree-challenge/src/usecases"
)

func CreatePokemonEntrypoint(paramPokemon dtos.PostPokemonDTO) (string, error) {
	var myPokemonRepository, error = repositories.NewPokemonRepository()

	if error != nil {
		return "", error
	}

	return usecases.CreatePokemonUseCase(paramPokemon, myPokemonRepository)
}
