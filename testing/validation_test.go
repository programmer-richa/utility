package testing

import (
	"fmt"
	"github.com/programmer-richa/utility/validators"
	"testing"
)

// TestPassword runs several test cases to check the correctness of
//the Password function defined in validators package.
func TestPassword(t *testing.T) {
	tests := []struct {
		name  string
		pass  string
		valid bool
	}{
		{
			"NoCharacterAtAll",
			"",
			false,
		},
		{
			"JustEmptyStringAndWhitespace",
			" \n\t\r\v\f ",
			false,
		},
		{
			"MixtureOfEmptyStringAndWhitespace",
			"U u\n1\t?\r1\v2\f34",
			false,
		},
		{
			"MissingUpperCaseString",
			"uu1?1234",
			false,
		},
		{
			"MissingLowerCaseString",
			"UU1?1234",
			false,
		},
		{
			"MissingNumber",
			"Uua?aaaa",
			false,
		},
		{
			"MissingSymbol",
			"Uu101234",
			false,
		},
		{
			"LessThanRequiredMinimumLength",
			"Uu1?123",
			false,
		},
		{
			"ValidPassword",
			"Uu1?1234",
			true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if c.valid != validators.Password(c.pass) {
				t.Fatal("Password Validator Failed", c.name)
			} else {
				fmt.Println("Password Validator-", c.name, "Pass")
			}
		})
	}
}

// TestEmail runs several test cases to check the correctness of
//the Email function defined in validators package.
func TestEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		valid bool
	}{
		{
			name:  "Empty Email",
			email: "",
			valid: false,
		},
		{
			name:  "Without @ Symbol",
			email: "abcdgmail.yahoo",
			valid: false,
		},
		{
			name:  "Without . Symbol",
			email: "abcd@gmailyahoo",
			valid: false,
		},
		{
			name:  "With Special Characters",
			email: "ç$€§/az@gmail.com",
			valid: false,
		},
		{
			name:  "Ending with .",
			email: "abc@gmail.com.",
			valid: false,
		},
		{
			name:  "With _ character",
			email: "abc_one.a@gmail_yahoo.com",
			valid: true,
		},
		{
			name:  "With - character",
			email: "abc-one.a@gmail-yahoo.com",
			valid: true,
		},
		{
			name:  "Valid Email Address",
			email: "abc@gmail.com",
			valid: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if c.valid != validators.Email(c.email) {
				t.Fatal("Email Validator Failed", c.name)
			} else {
				fmt.Println("Email Validator-", c.name, "Pass")
			}
		})
	}
}

// TestEmpty runs several test cases to check the correctness of
//the Empty function defined in validators package.
func TestEmpty(t *testing.T) {
	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "Empty String",
			value: "",
			valid: true,
		},
		{
			name:  "Non Empty String",
			value: "abc",
			valid: false,
		},
		{
			name:  "String with only tab spaces",
			value: "\t",
			valid: true,
		},
		{
			name:  "String with new line character only",
			value: "\n",
			valid: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if c.valid != validators.Empty(c.value) {
				t.Fatal("Empty Validator Failed-", c.name)
			} else {
				fmt.Println("Empty Validator-", c.name, "Pass")
			}
		})
	}
}

// TestMinLength runs several test cases to check the correctness of
//the MinLength function defined in validators package.
func TestMinLength(t *testing.T) {
	tests := []struct {
		name  string
		value string
		min   int
		valid bool
	}{
		{
			name:  "Empty string (Min Length 5)",
			value: "",
			min:   5,
			valid: false,
		},
		{
			name:  "Non-Empty string of length 3 (Min Length 5)",
			value: "abc",
			min:   5,
			valid: false,
		},
		{
			name:  "String with only tab spaces  (Min Length 4)",
			value: "\t",
			min:   4,
			valid: false,
		},
		{
			name:  "String with only a new line character  (Min Length 5)",
			value: "\n",
			min:   5,
			valid: false,
		},
		{
			name:  "String with only one character  (Min Length 1)",
			value: "a",
			min:   1,
			valid: true,
		},
		{
			name:  "String with only two characters  (Min Length 1)",
			value: "ab",
			min:   1,
			valid: true,
		},
		{
			name:  "String with only one character  (Min Length 0)",
			value: "a",
			min:   0,
			valid: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if c.valid != validators.MinLength(c.value, c.min) {
				t.Fatal("MinLength Validator Failed-", c.name, c.value, c.min)
			} else {
				fmt.Println("MinLength Validator-", c.name, "Pass")
			}
		})
	}
}

// TestMaxLength runs several test cases to check the correctness of
//the MaxLength function defined in validators package.
func TestMaxLength(t *testing.T) {
	tests := []struct {
		name  string
		value string
		max   int
		valid bool
	}{
		{
			name:  "Empty string (Max Length 5)",
			value: "",
			max:   5,
			valid: true,
		},
		{
			name:  "Non-Empty string of length 3 (Max Length 5)",
			value: "abc",
			max:   5,
			valid: true,
		},
		{
			name:  "String with only tab spaces  (Max Length 4)",
			value: "\t",
			max:   4,
			valid: true,
		},
		{
			name:  "String with only a new line character  (Max Length 5)",
			value: "\n",
			max:   5,
			valid: true,
		},
		{
			name:  "String with only one character  (Max Length 1)",
			value: "a",
			max:   1,
			valid: true,
		},
		{
			name:  "String with only two characters  (Max Length 1)",
			value: "ab",
			max:   1,
			valid: false,
		},
		{
			name:  "String with zero character  (Max Length 0)",
			value: "",
			max:   0,
			valid: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if c.valid != validators.MaxLength(c.value, c.max) {
				t.Fatal("MaxLength Validator Failed-", c.name, c.value, c.max)
			} else {
				fmt.Println("MaxLength Validator-", c.name, "Pass")
			}
		})
	}
}

// TestLengthInRange runs several test cases to check the correctness of
//the LengthInRange function defined in validators package.
func TestLengthInRange(t *testing.T) {
	tests := []struct {
		name  string
		value string
		min   int
		max   int
		valid bool
	}{
		{
			name:  "Empty string (Min Length 5,Max Length 5)",
			value: "",
			min:   5,
			max:   5,
			valid: false,
		},
		{
			name:  "Non-Empty string of length 3 (Min Length 5,Max Length 5)",
			value: "abc",
			min:   5,
			max:   5,
			valid: false,
		},
		{
			name:  "String with only tab spaces  (Min Length 4,Max Length 5)",
			value: "\t",
			min:   4,
			max:   5,
			valid: false,
		},
		{
			name:  "String with only a new line character  (Min Length 5,Max Length 5)",
			value: "\n",
			min:   5,
			max:   5,
			valid: false,
		},
		{
			name:  "String with only one character  (Min Length 1,Max Length 5)",
			value: "a",
			min:   1,
			max:   5,
			valid: true,
		},
		{
			name:  "String with only two characters  (Min Length 1,Max Length 5)",
			value: "ab",
			min:   1,
			max:   5,
			valid: true,
		},
		{
			name:  "String with only one character  (Min Length 0,Max Length 5)",
			value: "a",
			min:   0,
			max:   5,
			valid: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if c.valid != validators.LengthInRange(c.value, c.min, c.max) {
				t.Fatal("LengthInRange Validator Failed-", c.name, c.value, c.min, c.max)
			} else {
				fmt.Println("LengthInRange Validator-", c.name, "Pass")
			}
		})
	}
}

// TestIsInteger runs several test cases to check the correctness of
//the TestIsInteger function defined in validators package.
func TestIsInteger(t *testing.T) {
	tests := []struct {
		name  string
		value string
		valid bool
	}{
		{
			name:  "Empty string",
			value: "",
			valid: false,
		},
		{
			name:  "Non-Empty string of characters",
			value: "abc",
			valid: false,
		},
		{
			name:  "String with only tab spaces",
			value: "\t",
			valid: false,
		},
		{
			name:  "String with only a new line character",
			value: "\n",
			valid: false,
		},
		{
			name:  "String with only one character",
			value: "a",
			valid: false,
		},
		{
			name:  "String with characters and tab",
			value: "ab\t",
			valid: false,
		},
		{
			name:  "String with float value",
			value: "10.56",
			valid: false,
		},
		{
			name:  "String with positive integral values",
			value: "10",
			valid: true,
		},
		{
			name:  "String with negative integral values",
			value: "-10",
			valid: true,
		},
		{
			name:  "String with zero as integral values",
			value: "0",
			valid: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			if c.valid != validators.IsInteger(c.value) {
				t.Fatal("IsInteger Validator Failed-", c.name)
			} else {
				fmt.Println("IsInteger Validator-", c.name, "Pass")
			}
		})
	}
}

// TestIntegerInRange runs several test cases to check the correctness of
//the IntegerInRange function defined in validators package.
func TestIntegerInRange(t *testing.T) {
	tests := []struct {
		name  string
		value string
		min   int
		max   int
		valid bool
	}{
		{
			name:  "Empty string",
			value: "",
			min:   1,
			max:   5,
			valid: false,
		},
		{
			name:  "Non-Empty string of characters",
			value: "abc",
			min:   1,
			max:   5,
			valid: false,
		},
		{
			name:  "String with only tab spaces",
			value: "\t",
			min:   1,
			max:   5,
			valid: false,
		},
		{
			name:  "String with only a new line character",
			value: "\n",
			min:   1,
			max:   5,
			valid: false,
		},
		{
			name:  "String with only one character",
			value: "a",
			min:   1,
			max:   5,
			valid: false,
		},
		{
			name:  "String with characters and tab",
			value: "ab\t",
			min:   1,
			max:   5,
			valid: false,
		},
		{
			name:  "String with float value",
			value: "10.56",
			min:   1,
			max:   5,
			valid: false,
		},
		{
			name:  "String with positive integral values",
			value: "10",
			min:   1,
			max:   10,
			valid: true,
		},
		{
			name:  "String with negative integral values",
			value: "-10",
			min:   -1,
			max:   5,
			valid: false,
		},
		{
			name:  "String with zero as integral values",
			value: "0",
			min:   0,
			max:   10,
			valid: true,
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {
			_, res := validators.IntegerInRange(c.value, c.min, c.max)
			if c.valid != res {
				t.Fatal("IsInteger Validator Failed-", c.name)
			} else {
				fmt.Println("IsInteger Validator-", c.name, "Pass")
			}
		})
	}
}
