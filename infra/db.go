package infra

import (
	"golang.org/x/net/context"
)
<<<<<<< HEAD
=======

// temp DB interface
>>>>>>> format code

// DB interface wraps the database method
type DB interface {
<<<<<<< HEAD
	ReadDocument(ctx context.Context, key string, result interface{}) error
	ReadDocuments(ctx context.Context, key []string, results interface{}) error
	CreateDocument(ctx context.Context, doc interface{}) error
	CreateDocuments(ctx context.Context, docs interface{}) error
	Query(ctx context.Context, query string, binVars map[string]interface{}) (interface{}, error)
}
=======
}

// <------------------------------------arango----------------------------------->

// DB interface wraps the database

//type DB interface {
//	Query(ctx context.Context, query string, bindVars map[string]interface{}) (driver.Cursor, error)
//	ReadDocument(ctx context.Context, collectionName, key string, result interface{}) error
//	ReadDocuments(ctx context.Context, collectionName string, key []string, results interface{}) error
//	CreateDocument(ctx context.Context, colName string, doc interface{}) error
//	CreateDocuments(ctx context.Context, colName string, docs interface{}) error
//	Query(ctx context.Context, query string, binVars map[string]interface{}) (interface{}, error)
//}

// <------------------------------------mongo----------------------------------->

// DB interface wraps the database

//type DB interface {
//	Ping(ctx context.Context) error
//	Disconnect(ctx context.Context) error
//	EnsureIndices(ctx context.Context, tab string, inds []DbIndex) error
//	DropIndices(ctx context.Context, tab string, inds []DbIndex) error
//	Insert(ctx context.Context, tab string, v interface{}) error
//	InsertMany(ctx context.Context, tab string, v []interface{}) error
//	List(ctx context.Context, tab string, filter DbQuery, skip, limit int64, v interface{}, sort ...interface{}) error
//	FindOne(ctx context.Context, tab string, filter DbQuery, v interface{}, sort ...interface{}) error
//	PartialUpdateMany(ctx context.Context, col string, filter DbQuery, data interface{}) error
//	PartialUpdateManyByQuery(ctx context.Context, col string, filter DbQuery, query UnorderedDbQuery) error
//	BulkUpdate(ctx context.Context, col string, models []mongo.WriteModel) error
//	Aggregate(ctx context.Context, col string, q []DbQuery, v interface{}) error
//	AggregateWithDiskUse(ctx context.Context, col string, q []DbQuery, v interface{}) error
//	Distinct(ctx context.Context, col, field string, q DbQuery, v interface{}) error
//	DeleteMany(ctx context.Context, col string, filter interface{}) error
//}

// DbIndex holds database index
type DbIndex struct {
	Name   string
	Keys   []DbIndexKey
	Unique *bool
	Sparse *bool

	// If ExpireAfter is defined the server will periodically delete
	// documents with indexed time.Time older than the provided delta.
	ExpireAfter *time.Duration
}

type DbIndexKey struct {
	Key string
	Asc interface{}
}

// DbQuery holds a database query
type DbQuery bson.D
type UnorderedDbQuery bson.M

type BulkWriteModel mongo.WriteModel
>>>>>>> format code
