package entrypoints

import (
	dtos "github.com/nicolasolmos/agree-challenge/src/DTOS"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories/mysql"
	usecases "github.com/nicolasolmos/agree-challenge/src/usecases"
)

func CreatePokemonEntrypoint(paramPokemon dtos.PostPokemonDTO) {
	var myPokemonRepository = repositories.NewPokemonRepository()
	usecases.CreatePokemonUseCase(paramPokemon, myPokemonRepository)
}
