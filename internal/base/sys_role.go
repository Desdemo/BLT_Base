package base

type Role struct {
	Model
	Name        string `json:"name"`
	Description string `json:"description"`
}
