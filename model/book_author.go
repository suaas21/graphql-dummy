package model

type Book struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	AuthorIDs   []uint `json:"author_ids"`
}

type Author struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	BookIDs []uint `json:"book_ids"`
}
