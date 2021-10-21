package main

import (
	"database/sql"
	"fmt"
	"net/http"
	controllers "nicolas-olmos/agree-challenge/pokemon/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Relocate this code to pokemon.go file
type Pokemon struct {
	Name             string  `json:"name" binding: "required"`
	Health           int     `json:"health" binding: "required"`
	IsFirstEdition   bool    `json:"isFirstEdition" binding: "required"`
	ExpansionDeck    string  `json:"expansionDeck" binding: "required"`
	PokemonType      string  `json:"pokemonType" binding: "required"`
	Oddity           string  `json:"oddity" binding: "required"`
	Price            float32 `json:"price" binding: "required"`
	CardPicture      string  `json:"cardPicture" binding: "required"`
	CardCreationDate string  `json:"cardCreationDate" binding: "required"`
}

var db *sql.DB

func connectDB() {
	// I am perfectly concious that hardcoding a password in code is a terrible idea.
	// But given the fact that this is just a test I will leave it there anyway.
	// I am also concious that I should not use the root user and instad create a secondary one.
	// But this is a test and so on...
	var error error
	db, error = sql.Open("mysql", "root:aa114477@tcp(127.0.0.1:3306)/agree")

	if error != nil {
		fmt.Println("Error opening the DB connection")
		panic(error.Error())
	}
}

//TODO: Implement OWASP WebSec validations
func getPokemonByIdController(context *gin.Context) {
	var id string
	var requestedPokemon Pokemon
	var IsFirstEdition string

	id = context.Param("id")

	if len(id) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"description": "Missing id on URL"})
		panic(error.Error)
	}

	databaseError := db.QueryRow("SELECT * FROM pokemon WHERE id = ?", id).Scan(&id, &requestedPokemon.Name, &requestedPokemon.Health, &IsFirstEdition, &requestedPokemon.ExpansionDeck, &requestedPokemon.PokemonType, &requestedPokemon.Oddity, &requestedPokemon.Price, &requestedPokemon.CardPicture, &requestedPokemon.CardCreationDate)

	switch IsFirstEdition {
	case "\x00":
		requestedPokemon.IsFirstEdition = false
	case "\x01":
		requestedPokemon.IsFirstEdition = true
	}

	if databaseError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"description": "Could not get the specified pokemon"})
		panic(error.Error)
	}

	context.JSON(http.StatusOK, gin.H{"data": requestedPokemon})
}

func updatePokemonController(context *gin.Context) {

}

func getAllPokemonsController(context *gin.Context) {
	var pokemonArray []Pokemon

	rows, databaseError := db.Query("SELECT * FROM pokemon")

	if databaseError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"description": "Can not get all pokemons from database"})
		panic(error.Error)
	}

	for rows.Next() {
		var myPokemon Pokemon
		var id string
		var IsFirstEdition string
		rows.Scan(&id, &myPokemon.Name, &myPokemon.Health, &IsFirstEdition, &myPokemon.ExpansionDeck, &myPokemon.PokemonType, &myPokemon.Oddity, &myPokemon.Price, &myPokemon.CardPicture, &myPokemon.CardCreationDate)

		switch IsFirstEdition {
		case "\x00":
			myPokemon.IsFirstEdition = false
		case "\x01":
			myPokemon.IsFirstEdition = true
		}

		pokemonArray = append(pokemonArray, myPokemon)
	}

	context.JSON(http.StatusOK, gin.H{"data": pokemonArray})
}

func main() {
	connectDB()
	router := gin.Default()
	router.GET("/pokemon", getAllPokemonsController)
	router.GET("/pokemon/:id", getPokemonByIdController)
	router.PUT("/pokemon", controllers.PutPokemonController)
	router.DELETE("/pokemon/:id", controllers.DeletePokemonController)
	router.POST("/pokemon/:id", updatePokemonController)
	router.Run("localhost:8080")

	defer db.Close()
}
