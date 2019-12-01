package domain

type Book struct {
	ID     string `json:"id"`
	Tittle string `json:"title"`
	Year   int    `json:"year"`
	Author string `json:"author"`
}
