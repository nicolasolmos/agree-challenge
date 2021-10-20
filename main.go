package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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

//TODO: perform OWASP validations
func putPokemonController(context *gin.Context) {

	var newId string
	var json Pokemon
	var error error

	error = context.ShouldBindJSON(&json)
	newId = uuid.NewString()

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"status": "Bad Request"})
	}

	var isFirstEdition int

	isFirstEdition = 0

	if json.IsFirstEdition == true {
		isFirstEdition = 1
	}

	// TODO: replace or format this line below correctly.
	insertionQuery := fmt.Sprintf("INSERT INTO pokemon VALUES ('%s', '%s', %d, %b, '%s', '%s', '%s', %f, '%s', '%s')", newId, json.Name, json.Health, isFirstEdition, json.ExpansionDeck, json.PokemonType, json.Oddity, json.Price, json.CardPicture, json.CardCreationDate)
	insertion, error := db.Query(insertionQuery)

	if error != nil {
		fmt.Println("ERROR inserting a pokemon on the database.")
		context.JSON(http.StatusInternalServerError, gin.H{"status": "ERROR inserting a pokemon on the database."})
		panic(error.Error())
	}

	defer insertion.Close()

	context.JSON(http.StatusCreated, gin.H{"status": "created", "id": newId})
}

//TODO: Implement OWASP WebSec validations
func deletePokemonController(context *gin.Context) {

	var id string
	var databaseError error
	var deletionQuery string

	id = context.Param("id")

	if len(id) == 0 {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "description": "Missing id on URL"})
		panic(error.Error)
	}

	deletionQuery = fmt.Sprintf("DELETE FROM pokemon WHERE id = '%s'", id)

	deletion, databaseError := db.Query(deletionQuery)

	if databaseError != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"status": "Internal Server Error", "description": "Can not delete the specified id"})
		panic(error.Error)
	}

	defer deletion.Close()

	context.JSON(http.StatusOK, gin.H{"status": "Deleted"})
}

func getPokemonByIdController(context *gin.Context) {

}

func updatePokemonController(context *gin.Context) {

}

func getAllPokemonsController(context *gin.Context) {

}

func main() {
	connectDB()
	router := gin.Default()
	router.GET("/pokemon", getAllPokemonsController)
	router.GET("/pokemon/:id", getPokemonByIdController)
	router.PUT("/pokemon", putPokemonController)
	router.DELETE("/pokemon/:id", deletePokemonController)
	router.POST("/pokemon/:id", updatePokemonController)
	router.Run("localhost:8080")

	defer db.Close()
}
