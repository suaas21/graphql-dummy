package book

import (
	"context"
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo"
)

type bookArangoRepository struct {
	db  infra.DB
	ctx context.Context
	log logger.StructLogger
}

func NewArangoBookRepository(ctx context.Context, db infra.DB, lgr logger.StructLogger) repo.BookRepository {
	return &bookArangoRepository{
		ctx: ctx,
		db:  db,
		log: lgr,
	}
}

func (b *bookArangoRepository) CreateBook(book *model.Book) error {
	return b.db.CreateDocument(b.ctx, book)
}

func (a *bookArangoRepository) GetBookByID(key string) (*model.Book, error) {
	var book model.Book
	err := a.db.ReadDocument(a.ctx, key, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (a *bookArangoRepository) ListBookByIDs(keys []string) ([]model.Book, error) {
	books := make([]model.Book, len(keys))
	err := a.db.ReadDocuments(a.ctx, keys, books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (b *bookArangoRepository) UpdateBook(book model.Book) error {
	panic("implement me")
}

func (b *bookArangoRepository) DeleteBook(id uint) error {
	panic("implement me")
}
