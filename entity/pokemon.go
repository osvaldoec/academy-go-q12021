package entity

// Pokemon - basic struct for a pokemon model
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience string `json:"base_experience"`
}
