package forms

import (
	"github.com/programmer-richa/utility/constants"
	"github.com/programmer-richa/utility/validators"
	"errors"
	"regexp"
	"strings"
)

// ValidatorFunc specifies the criteria of testing the field data.
type ValidatorFunc func(interface{}) error

// PasswordValidator validates data using valid password criteria in package validators
func PasswordValidator(errorMsg string) ValidatorFunc {
	return ValidatorFunc(func(value interface{}) error {
		strValue, ok := value.(string)
		// check type assertion
		if !ok {
			return errors.New(constants.InvalidString)
		}
		// check password constraint
		if !validators.Password(strValue) {
			return errors.New(errorMsg)
		}
		return nil
	})
}

// ReValidator validates data using given regular expression pattern
func ReValidator(pattern string, errorMsg string) ValidatorFunc {
	validRe := regexp.MustCompile(pattern)
	return ValidatorFunc(func(value interface{}) error {
		strValue, ok := value.(string)
		// check type assertion
		if !ok {
			return errors.New(constants.InvalidString)
		}
		// check regular expression
		if !validRe.MatchString(strings.TrimSpace(strValue)) {
			return errors.New(errorMsg)
		}
		return nil
	})
}

// RangeValidator validates data using given range of values
func RangeValidator(min, max int, errorMsg string) ValidatorFunc {
	return ValidatorFunc(func(value interface{}) error {
		intValue, ok := value.(int)
		// check type assertion
		if !ok {
			return errors.New(constants.InvalidInteger)
		}
		// check range
		if intValue < min || intValue > max {
			return errors.New(errorMsg)
		}
		return nil
	})
}
