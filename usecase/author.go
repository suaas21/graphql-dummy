package usecase

import (
	"context"
	"github.com/suaas21/graphql-dummy/model"
)

type AuthorUsageCase interface {
	CreateAuthor(ctx context.Context, author *model.Author) error
	UpdateAuthor(ctx context.Context, author *model.Author) error
	DeleteAuthor(ctx context.Context, id string) error
	GetAuthor(ctx context.Context, id string) error

	GetAuthors(ctx context.Context, offset, limit int64, params GetBooksRequestParams) (result []*model.Author, count int64, err error)
}
