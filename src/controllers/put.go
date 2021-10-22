package controllers

import (
	"net/http"

	dtos "github.com/nicolasolmos/agree-challenge/src/DTOS"
	entrypoints "github.com/nicolasolmos/agree-challenge/src/entrypoints"

	"github.com/gin-gonic/gin"
)

// PutPokemon godoc
// @Summary Modify an existing pokemon and store it on the database
// @Description Modify an existing pokemon and store it on the database. This overrides existing values.
// @Produce json
// @Accepts json
// @Success 200 {object} entities.Pokemon
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /pokemon/{id} [put]
// @Param id path string false "string that contains the ID of the pokemon to be modified"
func PutPokemonController(context *gin.Context) {
	var myPutPokemonDTO dtos.PutPokemonDTO
	var bindingError error

	bindingError = context.ShouldBindJSON(&myPutPokemonDTO)

	myPutPokemonDTO.Id = context.Param("id")

	if bindingError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request"})
		panic("ERROR binding JSON")
	}

	entrypoints.UpdatePokemon(myPutPokemonDTO)

}
