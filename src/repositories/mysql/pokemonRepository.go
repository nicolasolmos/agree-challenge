package repositories

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nicolasolmos/agree-challenge/src/controllers/config"
	"github.com/nicolasolmos/agree-challenge/src/entities"
)

type PokemonRepository struct {
	db *sql.DB
}

func NewPokemonRepository() (*PokemonRepository, error) {

	var db *sql.DB
	var connectionError error

	// I am perfectly concious that hardcoding a password in code is a terrible idea.
	// But given the fact that this is just a test I will leave it there anyway.
	// I am also concious that I should not use the root user and instad create a secondary one.
	// But this is a test and so on...

	db, connectionError = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE))

	if connectionError != nil {
		log.Print(connectionError)
		return nil, entities.NewDatabaseError("Can not establish a database connection")

	}

	return &PokemonRepository{
			db: db,
		},
		nil
}

func (base PokemonRepository) Insert(paramPokemon entities.Pokemon) error {
	var insertionError error

	insertion, insertionError := base.db.Query("INSERT INTO pokemon VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);", paramPokemon.Id, paramPokemon.Name, paramPokemon.Health, paramPokemon.IsFirstEdition, paramPokemon.ExpansionDeck, paramPokemon.PokemonType, paramPokemon.Oddity, paramPokemon.Price, paramPokemon.CardPicture, paramPokemon.CardCreationDate)

	if insertionError != nil {
		log.Print(insertionError)
		return entities.NewDatabaseError("Can not insert a new pokemon on database")
	}

	defer insertion.Close()
	return nil
}

func (base PokemonRepository) Delete(paramPokemon entities.Pokemon) error {

	deletion, databaseError := base.db.Query("DELETE FROM pokemon WHERE id = ?", paramPokemon.Id)

	if databaseError != nil {
		log.Print(databaseError)
		return entities.NewDatabaseError("Can not delete the selected pokemon on database")
	}

	defer deletion.Close()

	return nil
}

func (baseRepository PokemonRepository) SelectAll() (*[]entities.Pokemon, error) {
	var pokemonArray []entities.Pokemon
	rows, databaseError := baseRepository.db.Query("SELECT * FROM pokemon")

	if databaseError != nil {
		log.Print(databaseError)
		return nil, entities.NewDatabaseError("Error when trying to select all rows from pokemon table")
	}

	defer rows.Close()

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

	return &pokemonArray, nil
}

func (baseRepository PokemonRepository) SelectAndFilter() *[]entities.Pokemon {
	var pokemonArray []entities.Pokemon
	return &pokemonArray
}
func (baseRepository PokemonRepository) SelectById(paramId string) (entities.Pokemon, error) {
	var IsFirstEdition string
	var myPokemon entities.Pokemon
	var databaseError error

	databaseError = baseRepository.db.QueryRow("SELECT * FROM pokemon WHERE id = ?", paramId).Scan(&myPokemon.Id, &myPokemon.Name, &myPokemon.Health, &IsFirstEdition, &myPokemon.ExpansionDeck, &myPokemon.PokemonType, &myPokemon.Oddity, &myPokemon.Price, &myPokemon.CardPicture, &myPokemon.CardCreationDate)

	switch IsFirstEdition {
	case "\x00":
		myPokemon.IsFirstEdition = false
	case "\x01":
		myPokemon.IsFirstEdition = true
	}

	if databaseError != nil {
		return entities.Pokemon{}, entities.NewDatabaseError("Can not get the selected pokemon from database")
	}

	return myPokemon, nil
}

func (baseRepository PokemonRepository) Selectpage() *[]entities.Pokemon {
	var pokemonArray []entities.Pokemon
	return &pokemonArray
}

func (baseRepository PokemonRepository) UpdateAll(paramPokemon entities.Pokemon) error {
	var isFirstEdition string = "\x00"

	if paramPokemon.IsFirstEdition {
		isFirstEdition = "\x01"
	}
	row, databaseError := baseRepository.db.Query("UPDATE pokemon SET name = ?, health = ?, is_first_edition = ?,  expansion_deck = ?, pokemon_type = ?, oddity = ?, price  = ?, card_picture = ?, card_creation_date = ? WHERE id = ?;", paramPokemon.Name, paramPokemon.Health, isFirstEdition, paramPokemon.ExpansionDeck, paramPokemon.PokemonType, paramPokemon.Oddity, paramPokemon.Price, paramPokemon.CardPicture, paramPokemon.CardCreationDate, paramPokemon.Id)

	if databaseError != nil {
		log.Print(databaseError)
		return entities.NewDatabaseError("Can not update the selected pokemon on database")
	}

	defer row.Close()

	return nil
}
