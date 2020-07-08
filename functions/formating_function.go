package functions

import (
	"fmt"
	"math"
	"strings"
)

// TitleCase makes first letter of words uppercase in a string
func TitleCase(text string) string{
	return strings.Title(strings.ToLower(text))
}

// Round2DecimalPlaces rounds to a floating-point value to up by 2 decimal places
func Round2DecimalPlaces(value float64) float64{
	return math.Ceil(value *100)/100
}

// CurrencyFormat returns string representation of value
// rounded to a floating-point value to up by 2 decimal places
func CurrencyFormat(value float64) string{
	return fmt.Sprintf("%.2f",math.Ceil(value *100)/100)
}