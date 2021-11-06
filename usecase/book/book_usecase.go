package book

import (
	"context"
	"fmt"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo"
	"github.com/suaas21/graphql-dummy/usecase"
	"strings"
)

type bookUseCase struct {
	repository repo.BookRepository
}

func NewBookUseCase(repository repo.BookRepository) usecase.BookUsageCase {
	return &bookUseCase{repository: repository}
}

func (b *bookUseCase) CreateBook(ctx context.Context, book *model.Book) error {
	return b.repository.CreateBook(ctx, book)
}

func (b *bookUseCase) UpdateBook(ctx context.Context, book *model.Book) error {
	return b.repository.UpdateBook(ctx, book)
}

func (b *bookUseCase) DeleteBook(ctx context.Context, id string) error {
	return b.repository.DeleteBook(ctx, id)
}

func (b *bookUseCase) GetBook(ctx context.Context, id string) (*model.Book, error) {
	return b.repository.GetBook(ctx, id)
}

func (b *bookUseCase) GetBooks(ctx context.Context, offset, limit int64, params usecase.GetBooksRequestParams) (result []*model.Book, count int64, err error) {
	query, err := buildQuery(offset, limit, params)
	if err != nil {
		return nil, 0, err
	}

	return b.repository.QueryBooks(ctx, query, nil)
}

func limitBuilder(offset, limit int64) string {
	if offset < 0 {
		offset = 0
	}

	if limit < 0 || limit > 100 {
		limit = 100
	}

	return fmt.Sprintf(`%d, %d`, offset, limit)
}

func buildQuery(offset, limit int64, params usecase.GetBooksRequestParams) (query string, err error) {
	var extraFilter []string

	if params.Name != "" {
		extraFilter = append(extraFilter, fmt.Sprintf(`LOWER(x.name) LIKE '%%%s%%'`, params.Name))
	}

	if params.Description != "" {
		extraFilter = append(extraFilter, fmt.Sprintf(`LOWER(x.description) LIKE '%%%s%%'`, params.Description))
	}

	if len(params.AuthorIds) > 0 {
		extraFilter = append(extraFilter, fmt.Sprintf(`COUNT(INTERSECTION(author_ids, ['%s']))`, strings.Join(params.AuthorIds, "','")))
	}

	filter := strings.Join(extraFilter, " AND ")

	sortBy := "x.id"
	if params.SortBy != "" {
		sortBy = params.SortBy
	}

	limitQuery := limitBuilder(offset, limit)

	query = fmt.Sprintf(`
	for x in books 
		%s
		SORT %s
		Limit %s
		return x
	`, filter, sortBy, limitQuery)

	return
}
