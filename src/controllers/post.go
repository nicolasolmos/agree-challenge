package controllers

import (
	"net/http"

	dtos "github.com/nicolasolmos/agree-challenge/src/DTOS"
	entrypoints "github.com/nicolasolmos/agree-challenge/src/entrypoints"

	"github.com/gin-gonic/gin"
)

func PostPokemonController(context *gin.Context) {
	var myPostPokemonDTO dtos.PostPokemonDTO
	var bindingError error

	bindingError = context.ShouldBindJSON(&myPostPokemonDTO)

	if bindingError != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request"})
		panic("ERROR binding JSON")
	}

	context.JSON(http.StatusOK, gin.H{"id": entrypoints.CreatePokemonEntrypoint(myPostPokemonDTO)})

}
