package author

import (
	"context"
	"encoding/json"
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/model"
	"github.com/suaas21/graphql-dummy/repo"
)

const CollectionAuthor = "Author"

type authorArangoRepository struct {
	db  infra.ArangoDB
	log logger.StructLogger
}

func NewArangoAuthorRepository(db infra.ArangoDB, lgr logger.StructLogger) repo.AuthorRepository {
	return &authorArangoRepository{
		db:  db,
		log: lgr,
	}
}

func (a *authorArangoRepository) CreateAuthor(ctx context.Context, author model.Author) error {
	return a.db.CreateDocument(ctx, CollectionAuthor, &author)
}

func (a *authorArangoRepository) UpdateAuthor(ctx context.Context, author model.Author) error {
	return a.db.UpdateDocument(ctx, CollectionAuthor, author.ID, &author)
}

func (a *authorArangoRepository) DeleteAuthor(ctx context.Context, id string) error {
	return a.db.RemoveDocument(ctx, CollectionAuthor, id)
}

func (a *authorArangoRepository) GetAuthor(ctx context.Context, id string) (*model.Author, error) {
	var author model.Author
	if err := a.db.ReadDocument(ctx, CollectionAuthor, id, &author); err != nil {
		return nil, err
	}
	return &author, nil
}

func (a *authorArangoRepository) QueryAuthors(ctx context.Context, query string, binVars map[string]interface{}) ([]model.Author, error) {
	res, err := a.db.Query(ctx, query, binVars)
	if err != nil {
		return nil, err
	}

	authors := make([]model.Author, 0)
	dataBytes, err := json.Marshal(res)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(dataBytes, &authors)
	if err != nil {
		return nil, err
	}

	return authors, nil
}
