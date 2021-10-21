package pokemon

import (
	"net/http"
	dtos "nicolas-olmos/agree-challenge/pokemon/DTOS"
	entrypoints "nicolas-olmos/agree-challenge/pokemon/entrypoints"

	"github.com/gin-gonic/gin"
)

func PutPokemonController(context *gin.Context) {
	var myPutPokemonDTO dtos.PutPokemonDTO
	var bindingError error

	bindingError = context.ShouldBindJSON(&myPutPokemonDTO)

	if bindingError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request"})
		panic("ERROR binding JSON")
	}

	entrypoints.CreatePokemonEntrypoint(myPutPokemonDTO)

}
