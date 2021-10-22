package pokemon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entities "github.com/nicolasolmos/agree-challenge/pokemon/entities"
	entrypoints "github.com/nicolasolmos/agree-challenge/pokemons/entrypoints"
)

func GetPokemonByIdController(context *gin.Context) {
	id := context.Param("id")

	if len(id) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"description": "Missing id on URL"})
		panic(error.Error)
	}

	myPokemon := entrypoints.GetPokemonByIdEntrypoint(entities.PokemonDTO)

	context.JSON(http.StatusOK, myPokemon)
}
