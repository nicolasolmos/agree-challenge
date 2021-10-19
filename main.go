package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func putPokemonController(context *gin.Context) {
	Pokemon json
	error = context.ShouldBindJSON(&json);

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusCreated, gin.H{"status": "created"});

}

func deletePokemonController(context *gin.Context) {

}

func getPokemonByIdController(context *gin.Context) {

}

func updatePokemonController(context *gin.Context) {

}

func getAllPokemonsController(context *gin.Context) {

}

func main() {
	router := gin.Default()
    router.GET("/pokemon", getAllPokemonsController)
	router.GET("/pokemon/:id",getPokemonByIdController)
	router.PUT("/pokemon", putPokemonController)
	router.DELETE("/pokemon/:id", deletePokemonController)
	router.UPDATE("/pokemon/:id", updatePokemonController)
    router.Run("localhost:8080")
}