package arango

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/suaas21/graphql-dummy/config"
	"github.com/suaas21/graphql-dummy/infra"
)

// arangodb database driver to connect
type Arango struct {
	database driver.Database
}

// New returns a new instance of arangodb
func New(ctx context.Context, cfg *config.Arango) (driver.Database, error) {

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


func (d *Arango) ReadDocument(ctx context.Context, colName, key string, result interface{}) error {
	col, err := d.database.Collection(ctx, colName)
	if IsNotFound(err) != nil {
		return err
	}
	_, err = col.ReadDocument(ctx, key, &result)
	if err != nil {
		return err
	}
	return nil
}

func (d *Arango) ReadDocuments(ctx context.Context, colName string, keys []string, results interface{}) error {
	col, err := d.database.Collection(ctx, colName)
	if IsNotFound(err) != nil {
		return err
	}
	_, _, err = col.ReadDocuments(ctx, keys, &results)
	if err != nil {
		return err
	}
	return err
}

func (d *Arango) CreateDocument(ctx context.Context, colName string, doc interface{}) error {
	col, err := d.database.Collection(ctx, colName)
	if IsNotFound(err) != nil {
		return err
	}
	_, err = col.CreateDocument(ctx, doc)
	if err != nil {
		return err
	}
	return err
}

func (d *Arango) CreateDocuments(ctx context.Context, colName string, docs interface{}) error {
	col, err := d.database.Collection(ctx, colName)
	if IsNotFound(err) != nil {
		return err
	}
	_,_, err = col.CreateDocuments(ctx, docs)
	if err != nil {
		return err
	}
	return err
}

func (d *Arango) Query(ctx context.Context, query string, binVars map[string]interface{}) (interface{}, error) {
	cursor, err := d.database.Query(ctx, query, binVars)
	if err != nil {
		return nil, err
	}
	defer cursor.Close()
	results := make([]interface{}, 0)
	for cursor.HasMore() {
		var doc interface{}
		_, err = cursor.ReadDocument(ctx, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			return nil, err
		}
		results = append(results, doc)
	}
	return results, nil
}

func IsNotFound(err error) error {
	if driver.IsNotFound(err) {
		return infra.ErrNotFound
	} else if err != nil {
		return err
	}
	return nil
}
