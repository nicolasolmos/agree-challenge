package pokemon

import (
	entities "nicolas-olmos/agree-challenge/pokemon/entities"
	repositories "nicolas-olmos/agree-challenge/pokemon/repositories/mysql"
	pokemon "nicolas-olmos/agree-challenge/pokemon/usecases/getPokemon/getAll"
)

func GetAllPokemonsEntrypoint() *[]entities.Pokemon {
	var myRepository = repositories.NewPokemonRepository()
	return pokemon.GetAllPokemons(myRepository)
}
