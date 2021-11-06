package arango

import (
	"context"
	"github.com/arangodb/go-driver"
	"reflect"
)

type Client struct {
	db driver.Database
}

func NewClient(db driver.Database) *Client {
	return &Client{db: db}
}

func (c *Client) ReadDocument(ctx context.Context, col, key string, result interface{}) error {
	panic("implement me")
}

func (c *Client) ReadDocuments(ctx context.Context, col string, key []string, results interface{}) error {
	panic("implement me")
}

func (c *Client) CreateDocument(ctx context.Context, col string, doc interface{}) error {
	panic("implement me")
}

func (c *Client) CreateDocuments(ctx context.Context, col string, docs interface{}) error {
	panic("implement me")
}

func (c *Client) UpdateDocument(ctx context.Context, col, key string, doc interface{}) error {
	panic("implement me")
}

func (c *Client) RemoveDocument(ctx context.Context, col, key string) error {
	panic("implement me")
}

func (c *Client) DocumentExists(ctx context.Context, col, key string) (bool, error) {
	panic("implement me")
}

func (c *Client) Query(ctx context.Context, query string, binVars map[string]interface{}) (data interface{}, count int64, err error) {
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
