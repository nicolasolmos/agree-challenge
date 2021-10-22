package usecases

import (
	dtos "github.com/nicolasolmos/agree-challenge/src/DTOS"
	"github.com/nicolasolmos/agree-challenge/src/entities"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories"
)

func UpdatePokemon(paramPokemonDTO dtos.PutPokemonDTO, paramRepository repositories.Repository) {
	var myPokemon = entities.Pokemon{
		Name:             paramPokemonDTO.Name,
		Health:           paramPokemonDTO.Health,
		IsFirstEdition:   paramPokemonDTO.IsFirstEdition,
		ExpansionDeck:    paramPokemonDTO.ExpansionDeck,
		PokemonType:      paramPokemonDTO.PokemonType,
		Oddity:           paramPokemonDTO.Oddity,
		CardPicture:      paramPokemonDTO.CardPicture,
		CardCreationDate: paramPokemonDTO.CardCreationDate,
	}

	paramRepository.UpdateAll(myPokemon)

}
