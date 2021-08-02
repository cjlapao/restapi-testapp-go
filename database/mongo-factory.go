package database

import (
	"context"
	"time"

	"github.com/cjlapao/restapi-testapp-go/executioncontext"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoFactory MongoFactory Entity
type MongoFactory struct {
	Client   *mongo.Client
	Database *mongo.Database
	Context  *executioncontext.Context
}

// NewFactory Instantiate a new MongoDb Factory
func NewFactory() MongoFactory {
	factory := MongoFactory{}
	factory.GetClient()
	factory.GetContext()
	factory.GetDatabase()
	return factory
}

// GetContext Gets the Execution context
func (f *MongoFactory) GetContext() *executioncontext.Context {
	if f.Context != nil {
		return f.Context
	}

	context := executioncontext.Context{
		Database: executioncontext.Get().Database,
	}

	f.Context = &context

	return &context
}

// GetClient returns mongodb client
func (f *MongoFactory) GetClient() *mongo.Client {

	if f.Client != nil {
		return f.Client
	}

	connectionString := executioncontext.Get().MongoConnectionString
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		logger.LogError(err)
		return nil
	}

	// Check the connection
	err = client.Ping(ctx, nil)

	if err != nil {
		logger.LogError(err)
		return nil
	}

	f.Client = client

	logger.Debug("Client connection created successfully")
	return client
}

// GetDatabase returns MongoDb database
func (f *MongoFactory) GetDatabase() *mongo.Database {
	if f.Client == nil {
		f.Client = f.GetClient()
	}

	database := f.Client.Database(f.Context.Database)

	if database == nil {
		logger.Error("There was an error getting the database " + f.Context.Database)
		return nil
	}

	f.Database = database

	logger.Debug("Database was retrieved successfully")
	return database
}

// Find Finds documents in the database
func (f *MongoFactory) Find(collectionName string, filter bson.D) []*bson.M {
	if f.Database == nil {
		f.GetDatabase()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	collection := f.Database.Collection(collectionName)

	if collection == nil {
		return nil
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		logger.LogError(err)
		return nil
	}
	var elements []*bson.M
	for cur.Next(ctx) {
		var element bson.M
		err := cur.Decode(&element)
		if err != nil {
			logger.LogError(err)
			return nil
		}
		elements = append(elements, &element)
	}

	return elements
}

// InsertOne Inserts one document into the selected collection
func (f *MongoFactory) InsertOne(collectionName string, element interface{}) *mongo.InsertOneResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := f.Database.Collection(collectionName)

	insertResult, err := collection.InsertOne(ctx, element)

	if err != nil {
		logger.LogError(err)
		return nil
	}

	return insertResult
}

// InsertMany Inserts one document into the selected collection
func (f *MongoFactory) InsertMany(collectionName string, elements []interface{}) *mongo.InsertManyResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := f.Database.Collection(collectionName)

	insertResult, err := collection.InsertMany(ctx, elements)

	if err != nil {
		logger.LogError(err)
		return nil
	}

	return insertResult
}
