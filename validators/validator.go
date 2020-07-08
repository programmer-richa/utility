package validators

import (
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
func Email(s string) bool{
	re := regexp.MustCompile(EMAIL_PATTERN)
	return re.MatchString(s)
}


// Empty validates s value by checking if it is blank.
func Empty(s string) bool{
	return strings.TrimSpace(s)==""
}

// MinLength validates if the length of s is greater than or equal to the specified min value.
func MinLength(s string,min int) bool{
	l := utf8.RuneCountInString(strings.TrimSpace(s))
	return  l >= min
}

// MaxLength validates if the length s  is smaller than or equal to the specified max value.
func MaxLength(s string,max int) bool{
	l := utf8.RuneCountInString(strings.TrimSpace(s))
	return  l <= max
}

// LengthInRange validates if the length s is between specified min and max value.
func LengthInRange(s string, min int, max int) bool{
	l := utf8.RuneCountInString(strings.TrimSpace(s))
	return  l >= min && l <= max
}

// IsInteger validates if s contains integral value.
func IsInteger(s string) bool{
	_, err := strconv.Atoi(s)
	return err==nil
}

// IntegerInRange validates if s contains integral value,
// and its value lies between min and max values.
// This function returns error if s is non-integral.
func IntegerInRange(s string,min int,max int) (error,bool){
	n, err := strconv.Atoi(s)
	if err!=nil{
		// non-numeric value
		return err,false
	}else if n<min || n>max{
		return nil,false
	}
	return nil,true
}