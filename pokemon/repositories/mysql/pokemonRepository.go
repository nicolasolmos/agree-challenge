package pokemon

import (
	"database/sql"
	"log"
	entities "nicolas-olmos/agree-challenge/pokemon/entities"
)

type PokemonRepository struct {
	db *sql.DB
}

func NewPokemonRepository() *PokemonRepository {

	var db *sql.DB
	var connectionError error

	// I am perfectly concious that hardcoding a password in code is a terrible idea.
	// But given the fact that this is just a test I will leave it there anyway.
	// I am also concious that I should not use the root user and instad create a secondary one.
	// But this is a test and so on...

	db, connectionError = sql.Open("mysql", "root:aa114477@tcp(127.0.0.1:3306)/agree")

	if connectionError != nil {
		panic("ERROR opening the DB connection")
	}
	return &PokemonRepository{
		db: db,
	}
}

func (base PokemonRepository) Insert(paramPokemon entities.Pokemon) {
	var insertionError error

	insertion, insertionError := base.db.Query("INSERT INTO pokemon VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", paramPokemon.Id, paramPokemon.Name, paramPokemon.Health, paramPokemon.IsFirstEdition, paramPokemon.ExpansionDeck, paramPokemon.PokemonType, paramPokemon.Oddity, paramPokemon.Price, paramPokemon.CardPicture, paramPokemon.CardCreationDate)

	if insertionError != nil {
		log.Fatal(insertionError)
		panic("ERROR when trying to insert pokemon on MySQL table")
	}

	defer insertion.Close()
}

func (base PokemonRepository) Delete(paramPokemon entities.Pokemon) {

	deletion, databaseError := base.db.Query("DELETE FROM pokemon WHERE id = ?", paramPokemon.Id)
	defer deletion.Close()

	if databaseError != nil {
		panic("ERROR when deleting spcified id on DB")
	}

}

func (baseRepository PokemonRepository) SelectAll() *[]entities.Pokemon {
	var pokemonArray []entities.Pokemon
	rows, databaseError := baseRepository.db.Query("SELECT * FROM pokemon")

	if databaseError != nil {
		panic("Error when trying to select all rows from pokemon table")
	}

	for rows.Next() {
		var myPokemon entities.Pokemon
		var IsFirstEdition string
		rows.Scan(&myPokemon.Id, &myPokemon.Name, &myPokemon.Health, &IsFirstEdition, &myPokemon.ExpansionDeck, &myPokemon.PokemonType, &myPokemon.Oddity, &myPokemon.Price, &myPokemon.CardPicture, &myPokemon.CardCreationDate)

		switch IsFirstEdition {
		case "\x00":
			myPokemon.IsFirstEdition = false
		case "\x01":
			myPokemon.IsFirstEdition = true
		}

		pokemonArray = append(pokemonArray, myPokemon)
	}

	return &pokemonArray
}

func (baseRepository PokemonRepository) SelectAndFilter() *[]entities.Pokemon {
	var pokemonArray []entities.Pokemon
	return &pokemonArray
}
func (baseRepository PokemonRepository) SelectById() *[]entities.Pokemon {
	var pokemonArray []entities.Pokemon
	return &pokemonArray
}
func (baseRepository PokemonRepository) Selectpage() *[]entities.Pokemon {
	var pokemonArray []entities.Pokemon
	return &pokemonArray
}
func (baseRepository PokemonRepository) UpdateAll(entities.Pokemon) {}
