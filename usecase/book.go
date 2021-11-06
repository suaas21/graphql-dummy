package usecase

import (
	"context"
	"github.com/suaas21/graphql-dummy/model"
)

/*
type BooleanSelectionEnum int32

const (
	BooleanSelectionEnum_UNSELECTED BooleanSelectionEnum = 0
	BooleanSelectionEnum_TRUE       BooleanSelectionEnum = 1
	BooleanSelectionEnum_FALSE      BooleanSelectionEnum = 2
)
*/

type GetBooksRequestParams struct {
	Name        string
	Description string
	AuthorIds   []string
	SortBy      string
}

type BookUsageCase interface {
	CreateBook(ctx context.Context, book *model.Book) error
	UpdateBook(ctx context.Context, book *model.Book) error
	DeleteBook(ctx context.Context, id string) error
	GetBook(ctx context.Context, id string) (*model.Book, error)

	GetBooks(ctx context.Context, offset, limit int64, params GetBooksRequestParams) (result []*model.Book, count int64, err error)
}
