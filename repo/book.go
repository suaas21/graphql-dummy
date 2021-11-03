package repo

import (
	"context"
	"github.com/suaas21/graphql-dummy/model"
)

type BookRepository interface {
	CreateBook(ctx context.Context, book model.Book) error
	UpdateBook(ctx context.Context, book model.Book) error
	DeleteBook(ctx context.Context, id string) error
	GetBook(ctx context.Context, id string) (*model.Book, error)

	QueryBooks(ctx context.Context, query string, binVars map[string]interface{}) ([]model.Book, error)
}
