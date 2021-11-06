package arango

import (
	"context"
	"github.com/arangodb/go-driver"
	"github.com/suaas21/graphql-dummy/infra"
	"reflect"
)

type ArangoClient struct {
	db driver.Database
}

func NewArangoClient(db driver.Database) infra.ArangoDB {
	return &ArangoClient{db: db}
}

func (c *ArangoClient) ReadDocument(ctx context.Context, col, key string, result interface{}) error {
	panic("implement me")
}

func (c *ArangoClient) ReadDocuments(ctx context.Context, col string, key []string, results interface{}) error {
	panic("implement me")
}

func (c *ArangoClient) CreateDocument(ctx context.Context, col string, doc interface{}) error {
	panic("implement me")
}

func (c *ArangoClient) CreateDocuments(ctx context.Context, col string, docs interface{}) error {
	panic("implement me")
}

func (c *ArangoClient) UpdateDocument(ctx context.Context, col, key string, doc interface{}) error {
	panic("implement me")
}

func (c *ArangoClient) RemoveDocument(ctx context.Context, col, key string) error {
	panic("implement me")
}

func (c *ArangoClient) DocumentExists(ctx context.Context, col, key string) (bool, error) {
	panic("implement me")
}

func (c *ArangoClient) Query(ctx context.Context, query string, binVars map[string]interface{}) (data interface{}, count int64, err error) {
	ct := driver.WithQueryFullCount(ctx, true)
	cursor, err := c.db.Query(ct, query, binVars)

	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close()

	results := make([]interface{}, 0)
	for {
		var doc interface{}
		_, err := cursor.ReadDocument(ct, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, 0, err
		}
		if !reflect.ValueOf(doc).IsNil() {
			results = append(results, doc)
		}
	}

	return results, cursor.Statistics().FullCount(), nil
}
