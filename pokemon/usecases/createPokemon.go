package usecases

import (
	dtos "github.com/nicolasolmos/agree-challenge/pokemon/DTOS"
	entities "github.com/nicolasolmos/agree-challenge/pokemon/entities"
	repositories "github.com/nicolasolmos/agree-challenge/pokemon/repositories"

	uuid "github.com/google/uuid"
)

func CreatePokemonUseCase(paramPokemon dtos.PutPokemonDTO, paramRepository repositories.Repository) {
	var newPokemon entities.Pokemon

	newPokemon.Id = uuid.NewString()
	newPokemon.Name = paramPokemon.Name
	newPokemon.Health = paramPokemon.Health
	newPokemon.IsFirstEdition = paramPokemon.IsFirstEdition
	newPokemon.ExpansionDeck = paramPokemon.ExpansionDeck
	newPokemon.PokemonType = paramPokemon.PokemonType
	newPokemon.Oddity = paramPokemon.Oddity
	newPokemon.Price = paramPokemon.Price
	newPokemon.CardPicture = paramPokemon.CardPicture
	newPokemon.CardCreationDate = paramPokemon.CardCreationDate

	paramRepository.Insert(newPokemon)

}
