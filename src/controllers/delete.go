package controllers

import (
	"net/http"

	entrypoints "github.com/nicolasolmos/agree-challenge/src/entrypoints"

	"github.com/gin-gonic/gin"
)

// DeletePokemon godoc
// @Summary Delete an existing pokemon on the database
// @Description Delete an existing pokemon on the database by provinding an Id
// @Success 200
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /pokemon/{id} [delete]
// @Param id path string false "string that contains the ID of the pokemon to be deleted"
func DeletePokemonController(context *gin.Context) {
	var id string

	id = context.Param("id")

	if len(id) == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "description": "Missing id on URL"})
		panic("Missing id on URL")
	}

	error := entrypoints.DestroyPokemonEntrypoint(id)

	if error != nil {
		context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"description": error.Error()})
	}

	context.JSON(http.StatusOK, gin.H{"status": "Deleted"})

}
