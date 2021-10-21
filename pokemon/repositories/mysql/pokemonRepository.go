package pokemon

import (
	"database/sql"
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

func (baseRepository PokemonRepository) SelectAll()       {}
func (baseRepository PokemonRepository) SelectAndFilter() {}
func (baseRepository PokemonRepository) SelectById()      {}
func (baseRepository PokemonRepository) Selectpage()      {}
func (baseRepository PokemonRepository) UpdateAll()       {}
