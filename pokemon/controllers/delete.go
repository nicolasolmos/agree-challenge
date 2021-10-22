package pokemon

import (
	"net/http"

	entrypoints "github.com/nicolasolmos/agree-challenge/pokemon/entrypoints"

	"github.com/gin-gonic/gin"
)

func DeletePokemonController(context *gin.Context) {
	var id string

	id = context.Param("id")

	if len(id) == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "description": "Missing id on URL"})
		panic("Missing id on URL")
	}

	entrypoints.DestroyPokemonEntrypoint(id)

	context.JSON(http.StatusOK, gin.H{"status": "Deleted"})

}
