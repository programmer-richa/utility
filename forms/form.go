// Package Form specifies the functions required to create and validate forms.
// This is helpful to add server-side validations.
package forms

import (
	"net/url"
)

// The below constant block contains the list of form field types.
// These are used in FieldBuilder and Field struct.
const (
	TEXT     = "text"
	TEXTAREA = "textarea"
	PASSWORD = "password"
	EMAIL    = "email"
	CHECKBOX = "checkbox"
	BUTTON   = "button"
)

// FormValues represents a map of form field values
type FormValues map[string]interface{}

// FormValidator represents a form validator function
type FormValidator func(formValues *FormValues) bool

// Form struct represents an HTML form
type Form struct {
	Fields          map[string]*Field
	FieldNames      []string
	Errors          map[string]error
	Values          FormValues
	validator       FormValidator
	PrefixFieldName string
	EncryptType     string
	Action          string
	Method          string
}

// New returns a pointer to Form
// It accepts prefix value for each field,
// encryption to specify if the form accepts multipart data,
// action specifies the url of form submission,
// and method (GET/POST) supported by the form
func New(prefix string, encryption string, action string, method string) *Form {
	return &Form{
		Fields:          make(map[string]*Field),
		FieldNames:      make([]string, 0),
		PrefixFieldName: prefix,
		EncryptType:     encryption,
		Action:          action,
		Method:          method,
	}
}

// WithField adds the Field produced by the FieldBuilder to the Form under the given name.
func (f *Form) WithField(name string, fb *FieldBuilder) *Form {
	field := fb.Build()
	f.Fields[name] = field
	f.FieldNames = append(f.FieldNames, name)
	return f
}

// WithValidator adds the FormValidator to the form.
func (f *Form) WithValidator(validator FormValidator) *Form {
	f.validator = validator
	return f
}

// Valid validates every field followed by the form's validator if provided.
func (f *Form) Valid(postForm url.Values) bool {
	valid := true

	f.Errors = nil
	f.Values = nil

	formValues := make(FormValues)
	formErrors := make(map[string]error)

	// validate fields
	for _, fname := range f.FieldNames {
		fieldValue, fieldError := f.Fields[fname].Validate(postForm.Get(fname))
		if fieldError != nil {
			valid = false
			formErrors[fname] = fieldError
		} else {
			formValues[fname] = fieldValue
		}
	}

	// validate form
	if valid && f.validator != nil {
		valid = f.validator(&formValues)
	}

	// if its valid, make the values available
	// otherwise make the errors available
	if valid {
		f.Values = formValues
	} else {
		f.Errors = formErrors
	}

	return valid
}
