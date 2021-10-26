package repo

import (
<<<<<<< HEAD
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/model"
	"golang.org/x/net/context"
)

type Book struct {
	db  infra.DB
	ctx context.Context
	log logger.StructLogger
}

func NewBook(ctx context.Context, db infra.DB, lgr logger.StructLogger) *Book {
	return &Book{
		ctx: ctx,
		db:  db,
		log: lgr,
	}
}

func (b *Book) CreateBookDocument(book *model.Book) error {
	err := b.db.CreateDocument(b.ctx, book)
	if err != nil {
		return err
	}
	return err
}

func (a *Book) GetBookDocument(key string) (*model.Book, error) {
	var book model.Book
	err := a.db.ReadDocument(a.ctx, key, &book)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (a *Book) GetBookDocuments(keys []string) ([]model.Book, error) {
	books := make([]model.Book, len(keys))
	err := a.db.ReadDocuments(a.ctx, keys, books)
	if err != nil {
		return nil, err
	}
	return books, nil
}
=======
	"github.com/suaas21/graphql-dummy/model"
)

type BookRepository interface {
	CreateBook(book model.Book) error
	UpdateBook(book model.Book) error
	DeleteBook(id uint) error
}
>>>>>>> inroduce new structure for repository
