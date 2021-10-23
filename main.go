package main

import (
	"strconv"
	"time"

	controllers "github.com/nicolasolmos/agree-challenge/src/controllers"
	"github.com/nicolasolmos/agree-challenge/src/controllers/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// gin-swagger middleware
// swagger embed files
// @title Nicolas Olmos Agree Backend Challenge
// @version 1.0
// @description Nicolas Olmos Agree Backend Challenge.

// @contact.name Nicolas Olmos
// @contact.url http://www.swagger.io/support
// @contact.email desa.nolmos@gmail.com

// @license.name GPLv3.0
// @license.url https://www.gnu.org/licenses/gpl-3.0.html

// @host 0.0.0.0:8080
// @BasePath /pokemon
// @query.collection.format multi

func main() {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://0.0.0.0:3000", "http://ec2-3-145-30-65.us-east-2.compute.amazonaws.com:3000"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Origin", "*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/pokemon", controllers.GetAllPokemons)
	router.GET("/pokemon/:id", controllers.GetPokemonByIdController)
	router.PUT("/pokemon/:id", controllers.PutPokemonController)
	router.DELETE("/pokemon/:id", controllers.DeletePokemonController)
	router.POST("/pokemon", controllers.PostPokemonController)

	// Swagger related stuff
	url := ginSwagger.URL("http://localhost:3000/docs/swagger.json")
	router.Static("/docs", "./docs")
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	router.Run("0.0.0.0:" + strconv.FormatInt(config.API_PORT, 10))
}
