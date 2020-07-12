package database

import (
	"github.com/programmer-richa/utility/constants"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

//User holds basic information regarding site user
//MongoDB represents JSON documents in binary-encoded format called BSON behind the scenes.
//BSON extends the JSON model to provide additional data types
//and to be efficient for encoding and decoding within different languages
type User struct {
	UserName         string             `json:"username" bson:"username"`
	FirstName        string             `json:"first_name" bson:"first_name"`
	LastName         string             `json:"last_name" bson:"last_name"`
	Name             string             `json:"name" bson:"name"`
	Email            string             `json:"email" bson:"email"`
	Password         string             `json:"password" bson:"password"`
	SubscribeToEmail bool               `json:"subscribe" bson:"subscribe"`
	IsActiveAccount  bool               `json:"is_active_account" bson:"is_active_account"`
	Id               primitive.ObjectID `json:"id" bson:"_id"` // Data type  bson.ObjectId is for mongodb bson format
}

// NewUser returns a variable of type User
func NewUser(name string, email string, password string, subscribeToEmail bool,
	isActiveAccount bool, id primitive.ObjectID) User {
	return User{
		Name:             name,
		Email:            email,
		Password:         password,
		SubscribeToEmail: subscribeToEmail,
		IsActiveAccount:  isActiveAccount,
		Id:               id,
	}
}

// TestMongoGetSession runs several test cases to check the correctness of
// the session connectivity functionality defined in database package.
func TestMongoGetSession(t *testing.T) {
	tests := []struct {
		name   string
		helper *MongoHelper
		valid  bool
	}{
		{
			"Nil value",
			nil,
			false,
		}, {
			"Invalid  database credentials",
			NewMongoHelper(constants.DbUser, "invalid", constants.DbHost,
				constants.DbPort, constants.Database),
			false,
		}, {
			"Correct database credentials",
			NewMongoHelper(constants.DbUser, constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			client, err := c.helper.GetSession()
			if err != nil && c.valid {
				t.Fatal("Mongo DB GetSession Function ", c.name, client, err)
			} else {
				fmt.Println("Mongo DB GetSession Function-", c.name, "Pass")
			}
		})
	}
}

// TestMongoGetDatabase runs several test cases to check the correctness of
// the database connectivity functionality defined in database package.
func TestMongoGetDatabase(t *testing.T) {

	tests := []struct {
		name   string
		helper *MongoHelper
		valid  bool
	}{
		{
			"Nil value",
			nil,
			false,
		}, {
			"Invalid  database credentials",
			NewMongoHelper(constants.DbUser, "invalid", constants.DbHost,
				constants.DbPort, constants.Database),
			false,
		}, {
			"Correct database credentials",
			NewMongoHelper(constants.DbUser, constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			client, err := c.helper.GetSession()
			db, err := c.helper.GetDatabase(client)
			if err != nil && c.valid {
				t.Fatal("Mongo DB GetDatabase Function ", c.name, db, err)
			} else {
				fmt.Println("Mongo DB GetDatabase Function-", c.name, "Pass")
			}
		})
	}
}

// TestMongoInsertDocument runs several test cases to check the correctness of
// the insert document(s) in a collection functionality defined in database package.
func TestMongoInsertDocument(t *testing.T) {

	tests := []struct {
		name           string
		helper         *MongoHelper
		collectionName string
		entry          User
		valid          bool
	}{
		{
			"Nil value",
			nil,
			"test",
			NewUser("Richa", "programmer.richa@gmail.com", "123456",
				false, false, primitive.NewObjectID()),
			false,
		}, {
			"Invalid  database credentials",
			NewMongoHelper(constants.DbUser+"invalid", constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			"test",
			NewUser("Richa", "programmer.richa@gmail.com", "123456",
				false, false, primitive.NewObjectID()),
			false,
		}, {
			"Correct database credentials",
			NewMongoHelper(constants.DbUser, constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			"test",
			NewUser("Richa", "programmer.richa@gmail.com", "123456",
				false, false, primitive.NewObjectID()),
			true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			_, err := c.helper.InsertDocument(c.collectionName, c.entry)
			if err != nil && c.valid {
				t.Fatal("Mongo DB InsertDocument Function ", c.name, err)
			} else {
				fmt.Println("Mongo DB InsertDocument Function-", c.name, "Pass")
			}
		})
	}
}

// TestMongoIsExistingDocument runs several test cases to check the correctness of
// the availability of a document in a collection functionality defined in database package.
func TestMongoIsExistingDocument(t *testing.T) {

	tests := []struct {
		name           string
		helper         *MongoHelper
		collectionName string
		entry          User
		valid          bool
	}{
		{
			"Nil value",
			nil,
			"test",
			NewUser("Richa", "programmer.richa@gmail.com", "123456",
				false, false, primitive.NewObjectID()),
			false,
		}, {
			"Invalid  database credentials",
			NewMongoHelper(constants.DbUser+"invalid", constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			"test",
			NewUser("Richa", "programmer.richa@gmail.com", "123456",
				false, false, primitive.NewObjectID()),
			false,
		}, {
			"Correct database credentials",
			NewMongoHelper(constants.DbUser, constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			"test",
			NewUser("Richa", "programmer.richa@gmail.com", "123456",
				false, false, primitive.NewObjectID()),
			true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			//c.helper.InsertDocument(c.collectionName, c.entry)
			_, err := c.helper.IsExistingDocument(c.collectionName, bson.M{"_id": c.entry.Id.Hex()})
			if err != nil && c.valid {
				t.Fatal("Mongo DB IsExistingDocument Function ", c.name, err)
			} else {
				fmt.Println("Mongo DB IsExistingDocument Function-", c.name, "Pass")
			}
		})
	}
}

// TestMongoGetaRecord runs several test cases to check the correctness of
// the retrieving an existing document in a collection functionality defined in database package.
func TestMongoGetaRecord(t *testing.T) {

	tests := []struct {
		name           string
		helper         *MongoHelper
		collectionName string
		valid          bool
	}{
		{
			"Nil value",
			nil,
			"test",
			false,
		}, {
			"Invalid  database credentials",
			NewMongoHelper(constants.DbUser+"invalid", constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			"test",
			false,
		}, {
			"Correct database credentials",
			NewMongoHelper(constants.DbUser, constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			"test",
			true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			//c.helper.InsertDocument(c.collectionName, c.entry)
			_, err := c.helper.GetaRecord(c.collectionName, bson.M{"email": "programmer.richa@gmail.com"})
			if err != nil && c.valid {
				t.Fatal("Mongo DB GetaRecord Function ", c.name, err)
			} else {
				fmt.Println("Mongo DB GetaRecord Function-", c.name, "Pass")
			}
		})
	}
}

// TestMongoFindDocument runs several test cases to check the correctness of
// the retrieving an existing document in a collection functionality defined in database package.
func TestMongoFindDocument(t *testing.T) {

	tests := []struct {
		name           string
		helper         *MongoHelper
		collectionName string
		valid          bool
	}{
		{
			"Nil value",
			nil,
			"test",
			false,
		}, {
			"Invalid  database credentials",
			NewMongoHelper(constants.DbUser+"invalid", constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			"test",
			false,
		}, {
			"Correct database credentials",
			NewMongoHelper(constants.DbUser, constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			"test",
			true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			//c.helper.InsertDocument(c.collectionName, c.entry)
			data, err := c.helper.GetaRecord(c.collectionName, bson.M{"email": "programmer.richa@gmail.com"})
			u := User{}
			bsonBytes, _ := bson.Marshal(data)
			bson.Unmarshal(bsonBytes, &u)
			_, err = c.helper.FindDocument(c.collectionName, u.Id.Hex())
			if err != nil && c.valid {
				t.Fatal("Mongo DB FindDocument Function ", c.name, err)
			} else {
				fmt.Println("Mongo DB FindDocument Function-", c.name, "Pass")
			}
		})
	}
}

// TestMongoUpdateDocument runs several test cases to check the correctness of
// the updating an existing document in a collection functionality defined in database package.
func TestMongoUpdateDocument(t *testing.T) {

	tests := []struct {
		name           string
		helper         *MongoHelper
		collectionName string
		valid          bool
	}{
		{
			"Nil value",
			nil,
			"test",
			false,
		}, {
			"Invalid  database credentials",
			NewMongoHelper(constants.DbUser+"invalid", constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			"test",
			false,
		}, {
			"Correct database credentials",
			NewMongoHelper(constants.DbUser, constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			"test",
			true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			//c.helper.InsertDocument(c.collectionName, c.entry)
			data, err := c.helper.GetaRecord(c.collectionName, bson.M{"email": "programmer.richa@gmail.com"})
			u := User{}
			bsonBytes, _ := bson.Marshal(data)
			bson.Unmarshal(bsonBytes, &u)
			u.Password = "newPass"
			err = c.helper.UpdateDocument(c.collectionName, u.Id.Hex(), u)
			if err != nil && c.valid {
				t.Fatal("Mongo DB FindDocument Function ", c.name, err)
			} else {
				fmt.Println("Mongo DB FindDocument Function-", c.name, "Pass")
			}
		})
	}
}

// TestMongoRemoveCollection runs several test cases to check the correctness of
// the remove collection functionality defined in database package.
func TestMongoRemoveCollection(t *testing.T) {

	tests := []struct {
		name           string
		helper         *MongoHelper
		collectionName string
		valid          bool
	}{
		{
			"Nil value",
			nil,
			"test",
			false,
		}, {
			"Invalid  database credentials",
			NewMongoHelper(constants.DbUser, "Invalid",
				constants.DbHost, constants.DbPort, constants.Database),
			"test",
			false,
		}, {
			"Correct database credentials",
			NewMongoHelper(constants.DbUser, constants.DbPassword, constants.DbHost,
				constants.DbPort, constants.Database),
			"test",
			true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			err := c.helper.RemoveCollection(c.collectionName)
			if err != nil && c.valid {
				t.Fatal("Mongo DB RemoveCollection Function ", c.name, err)
			} else {
				fmt.Println("Mongo DB RemoveCollection Function-", c.name, "Pass")
			}
		})
	}
}
