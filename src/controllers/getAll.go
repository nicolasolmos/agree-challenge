package controllers

import (
	"net/http"

	entities "github.com/nicolasolmos/agree-challenge/src/entities"
	entrypoints "github.com/nicolasolmos/agree-challenge/src/entrypoints"

	"github.com/gin-gonic/gin"
)

// GetAllPokemons godoc
// @Summary Get all pokemons stored on the database
// @Description Get all pokemons stored on the database
// @Produce  json
// @Success 200 {object} entities.Pokemon
/// @Failure 400
// @Failure 404
// @Failure 500
// @Router /pokemon/ [get]

func GetAllPokemons(context *gin.Context) {
	var pokemonArray *[]entities.Pokemon

	pokemonArray = entrypoints.GetAllPokemonsEntrypoint()
	context.JSON(http.StatusOK, gin.H{"data": pokemonArray})
}
