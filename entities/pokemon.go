package main

type Pokemon struct {
	name string `json:"name"`
	health int `json:"health"`
	isFirstEdition bool `json:"isFirstEdition"`
	expansionDeck string `json:"expansionDeck"`
	pokemonType string `json:"pokemonType"`
	oddity string `json:"oddity"`
	price float32 `json:"price"`
	cardPicture string `json:"cardPicture"`
	cardCreationDate string `json:"creationDate"`
}
