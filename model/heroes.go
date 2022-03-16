package model

type Heroes struct {
	ID         string `json:"id"`
	HeroesName string `json:"heroes_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Powers     string `json:"powers"`
	Equipment  string `json:"equipment"`
}
