package main

import (
    "net/http"
	"fmt"
    "github.com/gin-gonic/gin"
	"database/sql"
    "github.com/go-sql-driver/mysql"
)

var db DB

func connectDB() {
	// I am perfectly concious that hardcoding a password in code is a terrible idea.
	// But given the fact that this is just a test I will leave it there anyway.
	// I am also concious that I should not use the root user and instad create a secondary one.
	// But this is a test and so on...
	db, err := sql.Open("mysql", "root:aa114477@tcp(127.0.0.1:3306)/agree")

    if err != nil {
		ftm.Println("Error opening the DB connection")
        panic(err.Error())
    }    
}

func putPokemonController(context *gin.Context) {
	Pokemon json
	error = context.ShouldBindJSON(&json);

	if error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// TODO: replace or format this line below correctly.
	insertionQuery := fmt.Sprintf("INSERT INTO pokemon VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s)", json.name, json.health, json.isFirstEdition, json.expasionDeck, json.pokemonType, json.oddity, json.price, json.cardPicture, json.cardCreationDate)
	insert, error := db.Query(insertionQuery)

	if error != nil {
		fmt.Println("ERROR inserting a pokemon on the database.")
		context.JSON(http.StatusInternalServerError, gin.H{"status": "ERROR inserting a pokemon on the database."});
        panic(err.Error())
    }
    
    defer insert.Close()

	context.JSON(http.StatusCreated, gin.H{"status": "created"})
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
	connectDB()
	router := gin.Default()
    router.GET("/pokemon", getAllPokemonsController)
	router.GET("/pokemon/:id",getPokemonByIdController)
	router.PUT("/pokemon", putPokemonController)
	router.DELETE("/pokemon/:id", deletePokemonController)
	router.UPDATE("/pokemon/:id", updatePokemonController)
    router.Run("localhost:8080")

	defer db.Close()
}