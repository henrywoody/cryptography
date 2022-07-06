package polyalphabetic

import (
	"strings"

	"github.com/henrywoody/cryptography/substitution/simple"
)

// NewVigenereKey creates and returns a key for a Vigenere cipher. A Vigenere
// cipher is a polyalphabetic substitution cipher where each of the ciphertext
// alphabets is a Caesar cipher alphabet (shifted).
//
// The keyword is used to pick the number of ciphertext alphabets (by the length
// of the keyword) and the offset for each (by the distance of each character in
// the keyword from 'A').
func NewVigenereKey(keyword string) Key {
	upperKeywordRunes := []rune(strings.ToUpper(keyword))
	key := make(Key, len(upperKeywordRunes))
	for i, keywordChar := range upperKeywordRunes {
		key[i] = simple.NewCaesarKey(int(keywordChar - 'A'))
	}
	return key
}
