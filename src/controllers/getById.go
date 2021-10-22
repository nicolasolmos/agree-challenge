package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicolasolmos/agree-challenge/src/entrypoints"
)

// GetPokemonById godoc
// @Summary Obtain a pokemons from the database selecting it by its ID
// @Description Obtain a pokemons from the database selecting it by its ID
// @Produce  json
// @Param id path string false "string that contains the ID of the pokemon to be selected"
// @Success 200 {object} entities.Pokemon
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /pokemon/{id} [get]
func GetPokemonByIdController(context *gin.Context) {
	id := context.Param("id")

	if len(id) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"description": "Missing id on URL"})
		panic(error.Error)
	}

	myPokemon := entrypoints.GetPokemonEntrypoint(id)

	context.JSON(http.StatusOK, myPokemon)
}
