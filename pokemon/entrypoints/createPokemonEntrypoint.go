package pokemon

import (
	dtos "nicolas-olmos/agree-challenge/pokemon/DTOS"
	repositories "nicolas-olmos/agree-challenge/pokemon/repositories/mysql"
	usecases "nicolas-olmos/agree-challenge/pokemon/usecases"
)

func CreatePokemonEntrypoint(paramPokemon dtos.PutPokemonDTO) {
	var myPokemonRepository = repositories.NewPokemonRepository()
	usecases.CreatePokemonUseCase(paramPokemon, myPokemonRepository)
}
