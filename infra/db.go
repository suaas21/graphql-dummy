package infra

import (
	"golang.org/x/net/context"
)

// DB interface wraps the database method
type DB interface {
	ReadDocument(ctx context.Context, key string, result interface{}) error
	ReadDocuments(ctx context.Context, key []string, results interface{}) error
	CreateDocument(ctx context.Context, doc interface{}) error
	CreateDocuments(ctx context.Context, docs interface{}) error
	Query(ctx context.Context, query string, binVars map[string]interface{}) (interface{}, error)
}
