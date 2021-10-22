package dtos

type PutPokemonDTO struct {
	Name             string  `json:"name"`
	Health           int     `json:"health"`
	IsFirstEdition   bool    `json:"isFirstEdition"`
	ExpansionDeck    string  `json:"expansionDeck"`
	PokemonType      string  `json:"pokemonType"`
	Oddity           string  `json:"oddity"`
	Price            float32 `json:"price"`
	CardPicture      string  `json:"cardPicture"`
	CardCreationDate string  `json:"creationDate"`
}
