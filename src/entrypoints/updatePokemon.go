package entrypoints

import (
	dtos "github.com/nicolasolmos/agree-challenge/src/DTOS"
	"github.com/nicolasolmos/agree-challenge/src/entities"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories/mysql"
	"github.com/nicolasolmos/agree-challenge/src/usecases"
)

func UpdatePokemon(paramPokemon dtos.PutPokemonDTO) *entities.DatabaseError {
	myRepository, error := repositories.NewPokemonRepository()

	if error != nil {
		return error
	}

	return usecases.UpdatePokemon(paramPokemon, myRepository)
}
