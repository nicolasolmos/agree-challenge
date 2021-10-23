package entrypoints

import (
	dtos "github.com/nicolasolmos/agree-challenge/src/DTOS"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories/mysql"
	"github.com/nicolasolmos/agree-challenge/src/usecases"
)

func UpdatePokemon(paramPokemon dtos.PutPokemonDTO) error {
	myRepository, error := repositories.NewPokemonRepository()

	if error != nil {
		return error
	}

	return usecases.UpdatePokemon(paramPokemon, myRepository)
}
