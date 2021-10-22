package controllers

import (
	"net/http"

	dtos "github.com/nicolasolmos/agree-challenge/src/DTOS"
	entrypoints "github.com/nicolasolmos/agree-challenge/src/entrypoints"

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
