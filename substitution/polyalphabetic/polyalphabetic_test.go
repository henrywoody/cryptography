package polyalphabetic_test

import (
	"strings"
	"testing"

	"github.com/henrywoody/cryptography/substitution/polyalphabetic"
)

func TestSubstitute(t *testing.T) {
	testTable := []struct {
		Keywords       []string
		Message        string
		ExpectedResult string
	}{
		{
			Keywords:       []string{""},
			Message:        "Hello",
			ExpectedResult: "Hello",
		},
		{
			Keywords: []string{
				"MXPVEUFBKYHQDTRJICOSZGAWNL",
				"EYLJBFRPOADVTXHMNIWKCZQSGU",
				"CQTHGMRYIXOLBFSKPUJNZEWDAV",
			},
			Message:        "HELLO WORLD",
			ExpectedResult: "BBLQH WRILV",
		},
		{
			Keywords:       []string{"ZEBRAS"},
			Message:        "FLEE AT ONCE. WE ARE DISCOVERED!",
			ExpectedResult: "SIAA ZQ LKBA. VA ZOA RFPBLUAOAR!",
		},
		{
			Keywords:       []string{"ZEBRAS"},
			Message:        "Flee at once. We are discovered!",
			ExpectedResult: "Siaa zq lkba. Va zoa rfpbluaoar!",
		},
		{
			Keywords:       []string{"DEFGHIJKLMNOPQRSTUVWXYZABC"},
			Message:        "The quick brown fox jumps over the lazy dog",
			ExpectedResult: "Wkh txlfn eurzq ira mxpsv ryhu wkh odcb grj",
		},
		{
			// example from: https://en.wikipedia.org/wiki/Vigenère_cipher
			Keywords: []string{
				"LMNOPQRSTUVWXYZABCDEFGHIJK",
				"EFGHIJKLMNOPQRSTUVWXYZABCD",
				"MNOPQRSTUVWXYZABCDEFGHIJKL",
				"OPQRSTUVWXYZABCDEFGHIJKLMN",
				"NOPQRSTUVWXYZABCDEFGHIJKLM",
			},
			Message:        "attack at dawn",
			ExpectedResult: "lxfopv ef rnhr",
		},
	}

	for _, testCase := range testTable {
		key := polyalphabetic.NewKeyFromKeywords(testCase.Keywords...)
		result := polyalphabetic.Substitute(key, testCase.Message)
		if result != testCase.ExpectedResult {
			t.Errorf(
				"Expected result for keywords '%s' and message '%s' to be '%s', received '%s'",
				strings.Join(testCase.Keywords, ","), testCase.Message, testCase.ExpectedResult, result,
			)
		}
	}
}
