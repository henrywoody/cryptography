package simple_test

import (
	"testing"

	"github.com/henrywoody/cryptography/substitution/simple"
)

// ZEBRAS keyword examples from Wikipedia: https://en.wikipedia.org/wiki/Substitution_cipher

func TestNewKeyFromKeyword(t *testing.T) {
	testTable := []struct {
		Keyword           string
		ExpectedKeyString string
	}{
		{
			Keyword:           "",
			ExpectedKeyString: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			Keyword:           "ABC",
			ExpectedKeyString: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		{
			Keyword:           "ZEBRAS",
			ExpectedKeyString: "ZEBRASCDFGHIJKLMNOPQTUVWXY",
		},
		{
			Keyword:           "QUICKREDFOX",
			ExpectedKeyString: "QUICKREDFOXABGHJLMNPSTVWYZ",
		},
	}

	for _, testCase := range testTable {
		key := simple.NewKeyFromKeyword(testCase.Keyword)
		keyString := key.String()
		if keyString != testCase.ExpectedKeyString {
			t.Errorf(
				"Expected key for keyword '%s' to be '%s', received '%s'",
				testCase.Keyword, testCase.ExpectedKeyString, keyString,
			)
		}
	}
}

func TestSubstitute(t *testing.T) {
	testTable := []struct {
		Keyword        string
		Message        string
		ExpectedResult string
	}{
		{
			Keyword:        "",
			Message:        "Hello",
			ExpectedResult: "Hello",
		},
		{
			Keyword:        "MXPVEUFBKYHQDTRJICOSZGAWNL",
			Message:        "HELLO WORLD",
			ExpectedResult: "BEQQR ARCQV",
		},
		{
			Keyword:        "ZEBRAS",
			Message:        "FLEE AT ONCE. WE ARE DISCOVERED!",
			ExpectedResult: "SIAA ZQ LKBA. VA ZOA RFPBLUAOAR!",
		},
		{
			Keyword:        "ZEBRAS",
			Message:        "Flee at once. We are discovered!",
			ExpectedResult: "Siaa zq lkba. Va zoa rfpbluaoar!",
		},
		{
			Keyword:        "DEFGHIJKLMNOPQRSTUVWXYZABC",
			Message:        "The quick brown fox jumps over the lazy dog",
			ExpectedResult: "Wkh txlfn eurzq ira mxpsv ryhu wkh odcb grj",
		},
	}

	for _, testCase := range testTable {
		key := simple.NewKeyFromKeyword(testCase.Keyword)
		result := simple.Substitute(key, testCase.Message)
		if result != testCase.ExpectedResult {
			t.Errorf(
				"Expected result for keyword '%s' and message '%s' to be '%s', received '%s'",
				testCase.Keyword, testCase.Message, testCase.ExpectedResult, result,
			)
		}
	}
}
