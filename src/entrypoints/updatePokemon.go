package entrypoints

import (
	dtos "github.com/nicolasolmos/agree-challenge/src/DTOS"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories/mysql"
	"github.com/nicolasolmos/agree-challenge/src/usecases"
)

func UpdatePokemon(paramPokemon dtos.PutPokemonDTO) {
	var myRepository = repositories.NewPokemonRepository()
	usecases.UpdatePokemon(paramPokemon, myRepository)
}
