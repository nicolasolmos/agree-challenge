package entrypoints

import (
	entities "github.com/nicolasolmos/agree-challenge/src/entities"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories/mysql"
	pokemon "github.com/nicolasolmos/agree-challenge/src/usecases/getPokemon/getAll"
)

func GetAllPokemonsEntrypoint() (*[]entities.Pokemon, *entities.DatabaseError) {
	myRepository, error := repositories.NewPokemonRepository()

	if error != nil {
		return nil, error
	}

	return pokemon.GetAllPokemons(myRepository)
}
