package forms

import (
	"github.com/programmer-richa/utility/constants"
	"github.com/programmer-richa/utility/validators"
	"net/http"
	"net/url"
	"testing"
	//"time"
)

func testForm() *Form {
	form := New("registration", "", "/registration", http.MethodPost)

	// username
	form.WithField(form.PrefixFieldName+"Name", new(FieldBuilder).
		Required().
		Label("Name").
		FieldType(TEXT).
		Loader(StringLoader).
		WithValidators(ReValidator(validators.NamePattern, constants.InvalidName)))

	// email
	form.WithField(form.PrefixFieldName+"Email", new(FieldBuilder).
		Required().
		Label("Email address").
		FieldType(EMAIL).
		Loader(StringLoader).
		WithValidators(ReValidator(validators.EmailPattern, constants.InvalidEmail)))

	// password
	form.WithField(form.PrefixFieldName+"Password", new(FieldBuilder).
		Required().
		Label("Password").
		FieldType(PASSWORD).
		Loader(StringLoader).
		WithValidators(PasswordValidator(constants.InvalidPassword)))

	// password
	form.WithField(form.PrefixFieldName+"CPassword", new(FieldBuilder).
		Required().
		Label("Confirm Password").
		FieldType(PASSWORD).
		Loader(StringLoader).
		WithValidators(PasswordValidator(constants.InvalidPassword)))

	// Subscribe
	form.WithField(form.PrefixFieldName+"Subscribe", new(FieldBuilder).
		Label("Subscribe to email alerts.").
		FieldType(CHECKBOX).
		Loader(StringLoader))

	// Submit
	form.WithField(form.PrefixFieldName+"Submit", new(FieldBuilder).
		Label("Register Now").
		FieldType(BUTTON).
		Loader(StringLoader))

	return form
}

func TestValidForm(t *testing.T) {
	formValues := url.Values{}
	formValues.Add("registrationName", "Richa Chawla")
	formValues.Add("registrationEmail", "programmer.richa@gmail.com")
	formValues.Add("registrationPassword", "Uu1?1234")
	formValues.Add("registrationCPassword", "Uu1?1234")
	formValues.Add("registrationSubscribe", "on")
	form := testForm()
	valid := form.Valid(formValues)
	//fmt.Println(form.Values)
	if !valid {
		t.Error("Form should be valid")
	}
}

func TestInvalidForm(t *testing.T) {
	formValues := url.Values{}
	formValues.Add("registrationName", "ABC")
	formValues.Add("registrationEmail", "programmer.richa@gmailcom")
	formValues.Add("registrationPassword", "Uu1?123")
	formValues.Add("registrationCPassword", "Uu1?123")
	formValues.Add("registrationSubscribe", "on")
	form := testForm()
	valid := form.Valid(formValues)
	//fmt.Println(form.Values)
	if valid {
		t.Error("Form should be invalid")
	}

	if form.Values != nil {
		t.Error("Form should not have values")
	}

	if form.Errors == nil {
		t.Error("Form should have errors")
	}

	for _, fieldName := range []string{"registrationName", "registrationEmail", "registrationPassword"} {
		if _, ok := form.Errors[fieldName]; !ok {
			t.Errorf("Form should have errors for field %s", fieldName)
		}
	}
}
