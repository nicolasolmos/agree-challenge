package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nicolasolmos/agree-challenge/src/entrypoints"
)

func GetPokemonByIdController(context *gin.Context) {
	id := context.Param("id")

	if len(id) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"description": "Missing id on URL"})
		panic(error.Error)
	}

	myPokemon := entrypoints.GetPokemonEntrypoint(id)

	context.JSON(http.StatusOK, myPokemon)
}
