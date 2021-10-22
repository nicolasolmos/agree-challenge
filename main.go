package main

import (
	controllers "github.com/nicolasolmos/agree-challenge/src/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/pokemon", controllers.GetAllPokemons)
	router.GET("/pokemon/:id", controllers.GetPokemonByIdController)
	router.PUT("/pokemon/:id", controllers.PutPokemonController)
	router.DELETE("/pokemon/:id", controllers.DeletePokemonController)
	router.POST("/pokemon", controllers.PostPokemonController)
	router.Run("0.0.0.0:3000")
}
