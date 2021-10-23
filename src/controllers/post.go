package controllers

import (
	"net/http"

	dtos "github.com/nicolasolmos/agree-challenge/src/DTOS"
	entrypoints "github.com/nicolasolmos/agree-challenge/src/entrypoints"

	"github.com/gin-gonic/gin"
)

// PostPokemon godoc
// @Summary Create a new pokemon and store it on the database
// @Description Create a new pokemon and store it on the database
// @Produce json
// @Accepts json
// @Success 200 {object} entities.Pokemon
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /pokemon [post]
func PostPokemonController(context *gin.Context) {
	var myPostPokemonDTO dtos.PostPokemonDTO
	var bindingError error

	bindingError = context.ShouldBindJSON(&myPostPokemonDTO)

	if bindingError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request"})
		panic("ERROR binding JSON")
	}
	id, error := entrypoints.CreatePokemonEntrypoint(myPostPokemonDTO)

	if error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"description": error.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"id": id})

}
