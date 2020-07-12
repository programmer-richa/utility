// Package validators specifies that the variables, constants and functions
// to validate form data.
package validators

// EmailPattern stores a regular expression to validate the email address.
const EmailPattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-_]{0,61}[a-zA-Z0-9])+\\.[a-zA-Z0-9](?:[a-zA-Z0-9-_]{0,61}[a-zA-Z0-9])*$"

// NamePattern stores a regular expression to validate the name.
const NamePattern = "^[A-Za-z ]{5,}$"
