package book

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo"
	"time"
)

type edgeRelation struct {
	XFrom     string      `json:"_from,omitempty"`
	XTo       string      `json:"_to,omitempty"`
	Key       string      `json:"_key,omitempty"`
	Relation  string      `json:"relation,omitempty"`
	From      string      `json:"from,omitempty"`
	To        string      `json:"to,omitempty"`
	CreatedAt string      `json:"created_at,omitempty"`
	UpdatedAt string      `json:"updated_at,omitempty"`
	Data      interface{} `json:"data,omitempty"`
	Type      string      `json:"type,omitempty"`
}

const CollectionBook = "Book"

type bookArangoRepository struct {
	db  infra.ArangoDB
	log logger.StructLogger
}

func NewArangoBookRepository(db infra.ArangoDB, lgr logger.StructLogger) repo.BookRepository {
	return &bookArangoRepository{
		db:  db,
		log: lgr,
	}
}

func (b *bookArangoRepository) CreateBook(ctx context.Context, book model.Book) error {
	if err := b.db.CreateDocument(ctx, CollectionBook, &book); err != nil {
		return err
	}

	for _, authorId := range book.AuthorIDs {
		if err := b.upsertBookAuthorEdge(ctx, book.ID, authorId); err != nil {
			return err
		}
	}

	return nil
}

func (b *bookArangoRepository) UpdateBook(ctx context.Context, book model.Book) error {
	if err := b.db.UpdateDocument(ctx, CollectionBook, book.ID, &book); err != nil {
		return err
	}

	for _, authorId := range book.AuthorIDs {
		if err := b.upsertBookAuthorEdge(ctx, book.ID, authorId); err != nil {
			return err
		}
	}

	return nil
}

func (b *bookArangoRepository) DeleteBook(ctx context.Context, id string) error {
	return b.db.RemoveDocument(ctx, CollectionBook, id)

}

func (b *bookArangoRepository) GetBook(ctx context.Context, id string) (*model.Book, error) {
	var book model.Book
	if err := b.db.ReadDocument(ctx, CollectionBook, id, &book); err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *bookArangoRepository) upsertBookAuthorEdge(ctx context.Context, bookId, authorId string) error {
	key := fmt.Sprintf("%s-%s", bookId, authorId)

	// look for existence
	if exist, err := b.db.DocumentExists(ctx, "book_author_edges", key); err != nil {
		return err
	} else if exist {
		// no need to create edge
		return nil
	}

	// insert story place relation
	relation := edgeRelation{Key: key, From: "author", To: "book", Relation: "has_many",
		XFrom: fmt.Sprintf("authors/%s", authorId),
		XTo:   fmt.Sprintf("%s/%s", CollectionBook, bookId), CreatedAt: time.Now().Format(time.RFC3339)}

	if err := b.db.CreateDocument(ctx, "book_author_edges", &relation); err != nil {
		return err
	}

	return nil
}

func (b *bookArangoRepository) QueryBooks(ctx context.Context, query string, binVars map[string]interface{}) ([]model.Book, error) {
	res, err := b.db.Query(ctx, query, binVars)
	if err != nil {
		return nil, err
	}

	books := make([]model.Book, 0)
	dataBytes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataBytes, &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}
