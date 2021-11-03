package arango

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/suaas21/graphql-dummy/config"
	"reflect"
)

type Client struct {
	db driver.Database
}

func NewClient(db driver.Database) *Client {
	return &Client{db: db}
}

// NewArangoDB returns a new instance of arangodb
func NewArangoDB(ctx context.Context, cfg *config.Arango) (driver.Database, error) {
	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{fmt.Sprintf("http://%s:%s", cfg.Host, cfg.Port)},
		TLSConfig: &tls.Config{ /*...*/ },
	})
	if err != nil {
		return nil, err
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(cfg.DBUser, cfg.DBPassword),
	})
	if err != nil {
		return nil, err
	}

	db, err := c.Database(ctx, cfg.DBName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (c *Client) ReadDocument(ctx context.Context, col, key string, result interface{}) error {
	collection, err := c.db.Collection(ctx, col)
	if err != nil {
		return err
	}
	_, err = collection.ReadDocument(ctx, key, result)
	return err
}

func (c *Client) ReadDocuments(ctx context.Context, col string, keys []string, results interface{}) error {
	collection, err := c.db.Collection(ctx, col)
	if err != nil {
		return err
	}
	_, _, err = collection.ReadDocuments(ctx, keys, results)
	return err
}

func (c *Client) CreateDocument(ctx context.Context, col string, doc interface{}) error {
	collection, err := c.db.Collection(ctx, col)
	if err != nil {
		return err
	}
	_, err = collection.CreateDocument(ctx, doc)
	return err
}

func (c *Client) CreateDocuments(ctx context.Context, col string, docs interface{}) error {
	collection, err := c.db.Collection(ctx, col)
	if err != nil {
		return err
	}
	_, _, err = collection.CreateDocuments(ctx, docs)
	return err
}

func (c *Client) UpdateDocument(ctx context.Context, col, key string, doc interface{}) error {
	collection, err := c.db.Collection(ctx, col)
	if err != nil {
		return err
	}
	_, err = collection.UpdateDocument(ctx, key, doc)
	return err
}

func (c *Client) RemoveDocument(ctx context.Context, col, key string) error {
	collection, err := c.db.Collection(ctx, col)
	if err != nil {
		return err
	}
	_, err = collection.RemoveDocument(ctx, key)
	return err
}

func (c *Client) DocumentExists(ctx context.Context, col, key string) (bool, error) {
	collection, err := c.db.Collection(ctx, col)
	if err != nil {
		return false, err
	}
	return collection.DocumentExists(ctx, key)
}

func (c *Client) Query(ctx context.Context, query string, binVars map[string]interface{}) (interface{}, error) {
	cursor, err := c.db.Query(ctx, query, binVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()

	results := make([]interface{}, 0)
	for {
		var doc interface{}
		_, err := cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}
		if !reflect.ValueOf(doc).IsNil() {
			results = append(results, doc)
		}
	}

	return results, nil
}
