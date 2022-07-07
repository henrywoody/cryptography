package polyalphabetic

import (
	"strconv"

	"github.com/henrywoody/cryptography/substitution/simple"
)

// NewGronsfeldKey creates and returns a key for a Gronsfeld cipher. A Gronsfeld
// cipher is a polyalphabetic substitution cipher where each of the ciphertext
// alphabets is a Caesar cipher alphabet (shifted).
//
// The keyword is used to pick the number of ciphertext alphabets (by the length
// of the keyword) and the offset for each (by the numeral value of each
// character). The keyword is composed of characters from the decimal characters
// 0-9.
func NewGronsfeldKey(keyword string) Key {
	keywordRunes := []rune(keyword)
	key := make(Key, len(keywordRunes))
	for i, keywordChar := range keywordRunes {
		keywordInt, err := strconv.Atoi(string(keywordChar))
		if err != nil {
			continue
		}
		key[i] = simple.NewCaesarKey(keywordInt)
	}
	return key
}
