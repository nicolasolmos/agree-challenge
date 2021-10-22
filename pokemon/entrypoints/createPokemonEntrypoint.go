package entrypoints

import (
	dtos "github.com/nicolasolmos/agree-challenge/pokemon/DTOS"
	repositories "github.com/nicolasolmos/agree-challenge/pokemon/repositories/mysql"
	usecases "github.com/nicolasolmos/agree-challenge/pokemon/usecases"
)

func CreatePokemonEntrypoint(paramPokemon dtos.PutPokemonDTO) {
	var myPokemonRepository = repositories.NewPokemonRepository()
	usecases.CreatePokemonUseCase(paramPokemon, myPokemonRepository)
}
