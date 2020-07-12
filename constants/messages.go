// Package constants provides list of standard messages.
package constants

// Error messages
const (
	NilTpl               = "Template pointer is not initialised."
	NilMongoClient       = "DB Client is not initialised."
	DatabaseNotConnected = "Unable to connect database."
	CollectionNotFound   = "Collection not found."
	NilMongoHelper       = "Mongo helper is not initialised."
	InvalidString        = "Enter a string value."
	InvalidInteger       = "Enter an integer value."
	InvalidName          = "Name must be at least 5 characters."
	InvalidPassword      = "Password must be at least 8 characters."
	InvalidEmail         = "Invalid email address."
)
