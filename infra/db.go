package infra

import (
	"context"
)

// ArangoDB interface wraps the database method
type ArangoDB interface {
	ReadDocument(ctx context.Context, col, key string, result interface{}) error
	ReadDocuments(ctx context.Context, col string, key []string, results interface{}) error
	CreateDocument(ctx context.Context, col string, doc interface{}) error
	CreateDocuments(ctx context.Context, col string, docs interface{}) error
	UpdateDocument(ctx context.Context, col, key string, doc interface{}) error
	RemoveDocument(ctx context.Context, col, key string) error
	DocumentExists(ctx context.Context, col, key string) (bool, error)

	Query(ctx context.Context, query string, binVars map[string]interface{}) (interface{}, error)
}
