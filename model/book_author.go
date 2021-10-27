package model

type Book struct {
	Xkey        string   `json:"_key"`
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	AuthorIDs   []string `json:"author_ids"`
}

type Author struct {
	Xkey    string   `json:"_key"`
	ID      string   `json:"id"`
	Name    string   `json:"name"`
	BookIDs []string `json:"book_ids"`
}
