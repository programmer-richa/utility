// Package database provides functions to connect and manipulate database data.
package database

import (
	"github.com/programmer-richa/utility/constants"
	"github.com/programmer-richa/utility/functions"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoHelper struct {
	dbUser     string
	dbPassword string
	dbHost     string
	dbPort     string
	database   string
}

func NewMongoHelper(dbUser string, dbPassword string, dbHost string, dbPort string, database string) *MongoHelper {
	return &MongoHelper{
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		database,
	}
}

// GetSession creates a db session and returns its reference.
func (m *MongoHelper) GetSession() (*mongo.Client, error) {
	if m == nil {
		return nil, errors.New(constants.NilMongoHelper)
	}
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://" + m.dbUser + ":" + m.dbPassword + "@" + m.dbHost + ":" + m.dbPort)
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	// Return successful connection
	return client, err

}

// GetDatabase returns database reference
func (m *MongoHelper) GetDatabase(client *mongo.Client) (*mongo.Database, error) {
	if m == nil {
		return nil, errors.New(constants.NilMongoHelper)
	}
	if client == nil {
		return nil, errors.New(constants.NilMongoClient)
	}
	db := client.Database(m.database)
	if db == nil {
		return nil, errors.New(constants.DatabaseNotConnected)
	}
	return db, nil
}

// GetCollection returns collection reference
func (m *MongoHelper) GetCollection(client *mongo.Client, collectionName string) (*mongo.Collection, error) {
	if m == nil {
		return nil, errors.New(constants.NilMongoHelper)
	}
	if client == nil {
		return nil, errors.New(constants.NilMongoClient)
	}
	db, err := m.GetDatabase(client)
	if err != nil {
		return nil, err
	}
	collection := db.Collection(collectionName)
	if err != nil {
		return nil, errors.New(constants.CollectionNotFound)
	}
	return collection, err
}

// InsertDocument inserts an entry in the specified collection name using the provided db session
func (m *MongoHelper) InsertDocument(collectionName string, entry interface{}) (interface{}, error) {
	if m == nil {
		return nil, errors.New(constants.NilMongoHelper)
	}
	// Create DB Session
	client, err := m.GetSession()
	if err != nil {
		return 0, err
	}
	// Close DB connection after this method is executed.
	defer client.Disconnect(context.TODO())

	// Initialize DB Collection here
	collection, err := m.GetCollection(client, collectionName)
	if err != nil {
		return nil, errors.New(constants.CollectionNotFound)
	}
	// Insert data into collection
	res, err := collection.InsertOne(context.TODO(), entry)
	if err != nil {
		return 0, err
	}
	// Return ID of inserted document
	return res.InsertedID, nil
}

// UpdateDocument updates an entry in the specified collection name
func (m *MongoHelper) UpdateDocument(collectionName string, id string, entry interface{}) error {
	if m == nil {
		return errors.New(constants.NilMongoHelper)
	}
	// Create DB Session
	client, err := m.GetSession()
	if err != nil {
		return err
	}
	// Close DB connection after this method is executed.
	defer client.Disconnect(context.TODO())
	//Initialize DB Collection here
	collection, err := m.GetCollection(client, collectionName)
	if err != nil {
		return errors.New(constants.CollectionNotFound)
	}
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	// Covert struct data to bson map
	doc := functions.ConvertToMap(entry, true, true)
	document := bson.M{"$set": doc}
	updateResult := collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": docID}, document)
	if updateResult.Err() != nil {
		return updateResult.Err()
	}
	return nil
}

// IsExistingDocument tests the existence of a record with the provided conditions.
// It returns true if the record is available in the collection.
func (m *MongoHelper) IsExistingDocument(collectionName string, condition bson.M) (found bool, err error) {
	if m == nil {
		return found, errors.New(constants.NilMongoClient)
	}
	// Create DB Session
	client, err := m.GetSession()
	if err != nil {
		return found, err
	}
	// Close DB connection after this method is executed.
	defer client.Disconnect(context.TODO())
	//Initialize DB Collection here
	collection, err := m.GetCollection(client, collectionName)
	if err != nil {
		return found, errors.New(constants.CollectionNotFound)
	}
	count, err := collection.CountDocuments(context.TODO(), condition)
	found = count > 0
	return found, err
}

// GetaRecord returns a record that satisfies the provided conditions.
// It returns nil if the record is unavailable in the collection.
func (m *MongoHelper) GetaRecord(collectionName string, condition bson.M) (record interface{}, err error) {
	if m == nil {
		return nil, errors.New(constants.NilMongoHelper)
	}
	// Create DB Session
	client, err := m.GetSession()
	if err != nil {
		return nil, err
	}
	// Close DB connection after this method is executed.
	defer client.Disconnect(context.TODO())
	//Initialize DB Collection here
	collection, err := m.GetCollection(client, collectionName)
	if err != nil {
		return record, errors.New(constants.CollectionNotFound)
	}
	err = collection.FindOne(context.TODO(), condition).Decode(&record)
	return record, err
}

// FindDocument returns a record that satisfies the provided conditions.
// It returns nil if the record is unavailable in the collection.
func (m *MongoHelper) FindDocument(collectionName string, id string) (record interface{}, err error) {
	if m == nil {
		return nil, errors.New(constants.NilMongoHelper)
	}
	// Create DB Session
	client, err := m.GetSession()
	if err != nil {
		return nil, err
	}
	// Close DB connection after this method is executed.
	defer client.Disconnect(context.TODO())
	//Initialize DB Collection here
	collection, err := m.GetCollection(client, collectionName)
	if err != nil {
		return record, errors.New(constants.CollectionNotFound)
	}
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(context.TODO(), bson.M{"_id": docID}).Decode(&record)
	return record, err
}

// RemoveCollection deletes the collection from the database.
// It returns nil if the collection successfully deleted.
func (m *MongoHelper) RemoveCollection(collectionName string) (err error) {
	if m == nil {
		return errors.New(constants.NilMongoHelper)
	}
	// Create DB Session
	client, err := m.GetSession()
	if err != nil {
		return err
	}
	// Close DB connection after this method is executed.
	defer client.Disconnect(context.TODO())
	//Initialize DB Collection here
	collection, err := m.GetCollection(client, collectionName)
	if err != nil {
		return errors.New(constants.CollectionNotFound)
	}
	err = collection.Drop(context.TODO())
	return err
}
