package repo

import (
	"context"
	"github.com/suaas21/graphql-dummy/model"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book model.Book) error
	UpdateBook(ctx context.Context, book model.Book) error
	DeleteBook(ctx context.Context, id uint) error

	GetBook(ctx context.Context, id uint) (*model.Book, error)
}
