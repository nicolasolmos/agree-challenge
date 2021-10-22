package entrypoints

import (
	entities "github.com/nicolasolmos/agree-challenge/pokemon/entities"
	repositories "github.com/nicolasolmos/agree-challenge/pokemon/repositories/mysql"
	pokemon "github.com/nicolasolmos/agree-challenge/pokemon/usecases/getPokemon/getAll"
)

func GetAllPokemonsEntrypoint() *[]entities.Pokemon {
	var myRepository = repositories.NewPokemonRepository()
	return pokemon.GetAllPokemons(myRepository)
}
