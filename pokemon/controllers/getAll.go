package pokemon

import (
	"net/http"
	entities "nicolas-olmos/agree-challenge/pokemon/entities"
	entrypoints "nicolas-olmos/agree-challenge/pokemon/entrypoints"

	"github.com/gin-gonic/gin"
)

func GetAllPokemons(context *gin.Context) {
	var pokemonArray *[]entities.Pokemon

	pokemonArray = entrypoints.GetAllPokemonsEntrypoint()
	context.JSON(http.StatusOK, gin.H{"data": pokemonArray})
}
