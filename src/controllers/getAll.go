package controllers

import (
	"net/http"

	entities "github.com/nicolasolmos/agree-challenge/src/entities"
	entrypoints "github.com/nicolasolmos/agree-challenge/src/entrypoints"

	"github.com/gin-gonic/gin"
)

func GetAllPokemons(context *gin.Context) {
	var pokemonArray *[]entities.Pokemon

	pokemonArray = entrypoints.GetAllPokemonsEntrypoint()
	context.JSON(http.StatusOK, gin.H{"data": pokemonArray})
}
