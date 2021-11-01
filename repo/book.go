package repo

import (
	"github.com/suaas21/graphql-dummy/model"
)

type BookRepository interface {
	CreateBook(book *model.Book) error
	GetBookByID(key string) (*model.Book, error)
	ListBookByIDs(keys []string) ([]model.Book, error)
	UpdateBook(book model.Book) error
	DeleteBook(id uint) error
}
