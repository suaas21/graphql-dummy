package arango

import (
	"crypto/tls"
	"fmt"
	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
	"github.com/suaas21/graphql-dummy/config"
	"golang.org/x/net/context"
	"reflect"
)

// arango client structure
type Arango struct {
	ctx context.Context
	db  driver.Database
	col driver.Collection
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

<<<<<<< HEAD
// NewArangoClient returns Arango database and collection
func NewArangoClient(ctx context.Context, db driver.Database, col driver.Collection) *Arango {
	return &Arango{
		ctx: ctx,
		db:  db,
		col: col,
=======
func (d *Arango) ReadDocument(ctx context.Context, colName, key string, result interface{}) error {
	col, err := d.database.Collection(ctx, colName)
	if IsNotFound(err) != nil {
		return err
>>>>>>> format code
	}
}

func (a *Arango) ReadDocument(ctx context.Context, key string, result interface{}) error {
	_, err := a.col.ReadDocument(ctx, key, result)
	if err != nil {
		return err
	}
	return nil
}

func (a *Arango) ReadDocuments(ctx context.Context, keys []string, results interface{}) error {
	_, _, err := a.col.ReadDocuments(ctx, keys, results)
	if err != nil {
		return err
	}
	return err
}

func (a *Arango) CreateDocument(ctx context.Context, doc interface{}) error {
	_, err := a.col.CreateDocument(ctx, doc)
	if err != nil {
		return err
	}
	return err
}

<<<<<<< HEAD
func (a *Arango) CreateDocuments(ctx context.Context, docs interface{}) error {
	_, _, err := a.col.CreateDocuments(ctx, docs)
=======
func (d *Arango) CreateDocuments(ctx context.Context, colName string, docs interface{}) error {
	col, err := d.database.Collection(ctx, colName)
	if IsNotFound(err) != nil {
		return err
	}
	_, _, err = col.CreateDocuments(ctx, docs)
>>>>>>> format code
	if err != nil {
		return err
	}
	return err
}

func (a *Arango) Query(ctx context.Context, query string, binVars map[string]interface{}) (interface{}, error) {
	cursor, err := a.db.Query(ctx, query, binVars)
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
