package pokemon

import (
	repositories "nicolas-olmos/agree-challenge/pokemon/repositories/mysql"
	usecases "nicolas-olmos/agree-challenge/pokemon/usecases"
)

func DestroyPokemonEntrypoint(paramId string) {
	var myPokemonRepository = repositories.NewPokemonRepository()
	usecases.DestroyPokemonUsecase(paramId, myPokemonRepository)
}
