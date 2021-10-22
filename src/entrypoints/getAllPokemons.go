package entrypoints

import (
	entities "github.com/nicolasolmos/agree-challenge/src/entities"
	repositories "github.com/nicolasolmos/agree-challenge/src/repositories/mysql"
	pokemon "github.com/nicolasolmos/agree-challenge/src/usecases/getPokemon/getAll"
)

func GetAllPokemonsEntrypoint() *[]entities.Pokemon {
	var myRepository = repositories.NewPokemonRepository()
	return pokemon.GetAllPokemons(myRepository)
}
