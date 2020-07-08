package testing

import (
	"fmt"
	"github.com/programmer-richa/utility/functions"
	"testing"
)

// TestTitleCase runs several test cases to check the correctness of
//the TitleCase function defined in functions package.
func TestTitleCase(t *testing.T) {
	tests := []struct {
		name string
		value   string
		correctOutput  string
	}{
		{
			"All Lower case values",
			"sahil chawla",
			"Sahil Chawla",
		},
		{
			"All Upper case values",
			"SAHIL CHAWLA",
			"Sahil Chawla",
		},
		{
			"Mixed case values",
			"sAhIl cHaWlA",
			"Sahil Chawla",
		},
		{
			"Title case values",
			"Sahil Chawla",
			"Sahil Chawla",
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {

			if  functions.TitleCase(c.value) != c.correctOutput {
				t.Fatal("TitleCase", c.name)
			} else {
				fmt.Println("TitleCase Validator-", c.name, "Pass")
			}
		})
	}
}

// TestCurrencyFormat runs several test cases to check the correctness of
//the CurrencyFormat function defined in functions package.
func TestCurrencyFormat(t *testing.T) {
	tests := []struct {
		name string
		value   float64
		correctOutput  string
	}{
		{
			"Float with no decimal places",
			10,
			"10.00",
		},
		{
			"Float with negative values",
			-10,
			"-10.00",
		},
		{
			"Float with more than 2 decimal places",
			12.4567,
			"12.46",
		},
		{
			"Float with more than 2 decimal places near 5",
			12.245,
			"12.25",
		},
	}

	for _, c := range tests {
		t.Run(c.name, func(t *testing.T) {

			if  functions.CurrencyFormat(c.value) != c.correctOutput {
				t.Fatal("CurrencyFormat", c.name)
			} else {
				fmt.Println("CurrencyFormat Validator-", c.name, "Pass")
			}
		})
	}
}