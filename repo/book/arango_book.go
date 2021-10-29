package book

import (
	"context"
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

type bookArangoRepository struct {
	db         infra.ArangoDB
	collection string
	ctx        context.Context
	log        logger.StructLogger
}

func NewArangoBookRepository(ctx context.Context, db infra.ArangoDB, collection string, lgr logger.StructLogger) repo.BookRepository {
	return &bookArangoRepository{
		ctx:        ctx,
		db:         db,
		collection: collection,
		log:        lgr,
	}
}

func (b *bookArangoRepository) CreateBook(book model.Book) error {
	if err := b.db.CreateDocument(b.ctx, b.collection, &book); err != nil {
		return err
	}

	for _, authorId := range book.AuthorIDs {
		if err := b.upsertBookAuthorEdge(fmt.Sprintf("%d", book.ID), fmt.Sprintf("%d", authorId)); err != nil {
			return err
		}
	}

	return nil
}

func (b *bookArangoRepository) UpdateBook(book model.Book) error {
	if err := b.db.UpdateDocument(b.ctx, b.collection, fmt.Sprintf("%d", book.ID), &book); err != nil {
		return err
	}

	for _, authorId := range book.AuthorIDs {
		if err := b.upsertBookAuthorEdge(fmt.Sprintf("%d", book.ID), fmt.Sprintf("%d", authorId)); err != nil {
			return err
		}
	}

	return nil
}

func (b *bookArangoRepository) DeleteBook(id uint) error {
	return b.db.RemoveDocument(b.ctx, b.collection, fmt.Sprintf("%d", id))

}

func (b *bookArangoRepository) upsertBookAuthorEdge(bookId, authorId string) error {
	key := fmt.Sprintf("%s-%s", bookId, authorId)

	// look for existence
	if exist, err := b.db.DocumentExists(b.ctx, "book_author_edges", key); err != nil {
		return err
	} else if exist {
		// no need to create edge
		return nil
	}

	// insert story place relation
	relation := edgeRelation{Key: key, From: "author", To: "book", Relation: "has_many",
		XFrom: fmt.Sprintf("authors/%s", authorId),
		XTo:   fmt.Sprintf("%s/%s", b.collection, bookId), CreatedAt: time.Now().Format(time.RFC3339)}

	if err := b.db.CreateDocument(b.ctx, "book_author_edges", &relation); err != nil {
		return err
	}

	return nil
}
