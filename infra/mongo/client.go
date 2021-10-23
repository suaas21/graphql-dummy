package mongo

import (
	"context"
	"encoding/json"
	"github.com/suaas21/graphql-dummy/infra"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/utils"
	"log"
	"time"

	//"bitbucket.org/evaly/go-boilerplate/infra"
	//"bitbucket.org/evaly/go-boilerplate/logger"
	//"bitbucket.org/evaly/go-boilerplate/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Mongo holds necessery fields and
// mongo database session to connect
type Mongo struct {
	*mongo.Client
	database *mongo.Database
	name     string
	lgr      logger.Logger
}

// New returns a new instance of mongodb using session s
func New(ctx context.Context, uri, name string, timeout time.Duration, opts ...Option) (*Mongo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	minPoolSize := uint64(10)
	maxPoolSize := uint64(100)
	connectionOption := &options.ClientOptions{
		SocketTimeout:          &timeout,
		ConnectTimeout:         &timeout,
		MaxPoolSize:            &maxPoolSize,
		MinPoolSize:            &minPoolSize,
		ServerSelectionTimeout: &timeout,
		RetryWrites:            utils.BoolP(true),
		ReadPreference:         readpref.Secondary(),
		//ReplicaSet:             nil,
		//Direct:                 nil,
	}
	log.Println("hitting mongo connect...")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri), connectionOption)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}
	log.Println("mongo connected...")

	db := &Mongo{
		Client:   client,
		database: client.Database(name),
		name:     name,
	}
	for _, opt := range opts {
		opt.apply(db)
	}
	return db, nil
}

// Option is mongo db option
type Option interface {
	apply(*Mongo)
}

// OptionFunc implements Option interface
type OptionFunc func(db *Mongo)

func (f OptionFunc) apply(db *Mongo) {
	f(db)
}

// SetLogger sets logger
func SetLogger(lgr logger.Logger) Option {
	return OptionFunc(func(db *Mongo) {
		db.lgr = lgr
	})
}

func (d *Mongo) println(args ...interface{}) {
	if d.lgr != nil {
		d.lgr.Println(args...)
	}
}

func (d *Mongo) Ping(ctx context.Context) error {
	return d.Client.Ping(ctx, readpref.Primary())
}

func (d *Mongo) Close(ctx context.Context) error {
	return d.Client.Disconnect(ctx)
}

// EnsureIndices creates indices for collection col
func (d *Mongo) EnsureIndices(ctx context.Context, col string, inds []infra.DbIndex) error {
	log.Println("creating indices for", col)
	db := d.database
	indexModels := []mongo.IndexModel{}
	for _, ind := range inds {
		keys := bson.D{}
		for _, k := range ind.Keys {
			keys = append(keys, bson.E{k.Key, k.Asc})
		}
		opts := options.Index()
		if ind.Unique != nil {
			opts.SetUnique(*ind.Unique)
		}
		if ind.Sparse != nil {
			opts.SetSparse(*ind.Sparse)
		}
		if ind.Name != "" {
			opts.SetName(ind.Name)
		}
		if ind.ExpireAfter != nil {
			opts.SetExpireAfterSeconds(int32(ind.ExpireAfter.Seconds()))
		}
		im := mongo.IndexModel{
			Keys:    keys,
			Options: opts,
		}
		indexModels = append(indexModels, im)
	}
	if _, err := db.Collection(col).Indexes().CreateMany(ctx, indexModels); err != nil {
		return err
	}
	return nil
}

// DropIndices drops indices from collection col
func (d *Mongo) DropIndices(ctx context.Context, col string, inds []infra.DbIndex) error {
	d.println("dropping indices from", col)
	if _, err := d.database.Collection(col).Indexes().DropAll(ctx); err != nil {
		return err
	}
	return nil
}

// Insert inserts doc into collection
func (d *Mongo) Insert(ctx context.Context, col string, doc interface{}) error {
	d.println("insert into", col)
	if _, err := d.database.Collection(col).InsertOne(ctx, doc); err != nil {

		return err
	}
	return nil
}

// Insert inserts doc into collection
func (d *Mongo) InsertMany(ctx context.Context, col string, docs []interface{}) error {
	d.println("insert many into", col)
	if _, err := d.database.Collection(col).InsertMany(ctx, docs); err != nil {
		return err
	}
	return nil
}

// FindOne finds a doc by query
func (d *Mongo) FindOne(ctx context.Context, col string, q infra.DbQuery, v interface{}, sort ...interface{}) error {
	d.println("find", q, "from", col)
	findOneOpts := options.FindOne()
	if len(sort) > 0 {
		findOneOpts = findOneOpts.SetSort(sort[0])
	}
	err := d.database.Collection(col).FindOne(ctx, q, findOneOpts).Decode(v)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return infra.ErrNotFound
		}
		return err
	}
	return nil
}

// List finds list of docs that matches query with skip and limit
func (d *Mongo) List(ctx context.Context, col string, filter infra.DbQuery, skip, limit int64, v interface{}, sort ...interface{}) error {
	d.println("list", filter, "from", col)
	findOpts := options.Find().SetSkip(skip).SetLimit(limit)
	if len(sort) > 0 {
		findOpts = findOpts.SetSort(sort[0])
	}
	cursor, err := d.database.Collection(col).Find(ctx, filter, findOpts)
	if err != nil {
		return err
	}
	if err := cursor.All(ctx, v); err != nil {
		return err
	}

	return nil
}

// Aggregate runs aggregation q on docs and store the result on v
func (d *Mongo) Aggregate(ctx context.Context, col string, q []infra.DbQuery, v interface{}) error {
	d.println("aggregate", q, "from", col)
	cursor, err := d.database.Collection(col).Aggregate(ctx, q)
	if err != nil {
		return err
	}
	if err := cursor.All(ctx, v); err != nil {
		return err
	}
	return nil
}

func (d *Mongo) AggregateWithDiskUse(ctx context.Context, col string, q []infra.DbQuery, v interface{}) error {
	d.println("aggregate", q, "from", col)
	opt := options.Aggregate().SetAllowDiskUse(true)
	cursor, err := d.database.Collection(col).Aggregate(ctx, q, opt)
	if err != nil {
		return err
	}
	if err := cursor.All(ctx, v); err != nil {
		return err
	}
	return nil
}

func (d *Mongo) Distinct(ctx context.Context, col, field string, q infra.DbQuery, v interface{}) error {
	d.println("aggregate", q, "from", col)
	interfaces, err := d.database.Collection(col).Distinct(ctx, field, q)
	if err != nil {
		return err
	}
	data, err := json.Marshal(interfaces)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (d *Mongo) PartialUpdateMany(ctx context.Context, col string, filter infra.DbQuery, data interface{}) error {
	_, err := d.database.Collection(col).UpdateMany(ctx, filter, bson.M{"$set": data})
	if err != nil {
		return err
	}
	return nil
}

func (d *Mongo) PartialUpdateManyByQuery(ctx context.Context, col string, filter infra.DbQuery, query infra.UnorderedDbQuery) error {
	_, err := d.database.Collection(col).UpdateMany(ctx, filter, query)
	if err != nil {
		return err
	}
	return nil
}

func (d *Mongo) BulkUpdate(ctx context.Context, col string, models []mongo.WriteModel) error {
	_, err := d.database.Collection(col).BulkWrite(ctx, models)
	return err
}

func (d *Mongo) DeleteMany(ctx context.Context, col string, filter interface{}) error {
	_, err := d.database.Collection(col).DeleteMany(ctx, filter)
	return err
}
