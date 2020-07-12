package validators

import (
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Password validates s password against the rules defined below.
//
// upp: at least one upper case letter.
// low: at least one lower case letter.
// num: at least one digit.
// sym: at least one special character.
// tot: at least eight characters long.
// No empty string or whitespace.
func Password(s string) bool {
	s = strings.TrimSpace(s)
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range s {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}

// Email validates s email against the EMAIL_PATTERN set in the constant.
func Email(s string) bool {
	s = strings.TrimSpace(s)
	re := regexp.MustCompile(EmailPattern)
	return re.MatchString(s)
}

// Empty validates s value by checking if it is blank.
func Empty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// MinLength validates if the length of s is greater than or equal to the specified min value.
func MinLength(s string, min int) bool {
	s = strings.TrimSpace(s)
	l := utf8.RuneCountInString(strings.TrimSpace(s))
	return l >= min
}

// MaxLength validates if the length s  is smaller than or equal to the specified max value.
func MaxLength(s string, max int) bool {
	s = strings.TrimSpace(s)
	l := utf8.RuneCountInString(strings.TrimSpace(s))
	return l <= max
}

// LengthInRange validates if the length s is between specified min and max value.
func LengthInRange(s string, min int, max int) bool {
	s = strings.TrimSpace(s)
	l := utf8.RuneCountInString(strings.TrimSpace(s))
	return l >= min && l <= max
}

// IsInteger validates if s contains integral value.
func IsInteger(s string) bool {
	s = strings.TrimSpace(s)
	_, err := strconv.Atoi(s)
	return err == nil
}

// IntegerInRange validates if s contains integral value,
// and its value lies between min and max values.
// This function returns error if s is non-integral.
func IntegerInRange(s string, min int, max int) (error, bool) {
	s = strings.TrimSpace(s)
	n, err := strconv.Atoi(s)
	if err != nil {
		// non-numeric value
		return err, false
	} else if n < min || n > max {
		return nil, false
	}
	return nil, true
}

// NotBlank is the validation function for validating if the current field
// has a value or length greater than zero, or is not a space only string.
func NotBlank(data interface{}) bool {
	field := reflect.ValueOf(data)
	switch field.Kind() {
	case reflect.String:
		return len(strings.TrimSpace(field.String())) > 0
	case reflect.Chan, reflect.Map, reflect.Slice, reflect.Array:
		return field.Len() > 0
	case reflect.Ptr, reflect.Interface, reflect.Func:
		return !field.IsNil()
	default:
		return field.IsValid() && field.Interface() != reflect.Zero(field.Type()).Interface()
	}
}
