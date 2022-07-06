package polyalphabetic

import (
	"strings"

	"github.com/henrywoody/cryptography/substitution/simple"
)

const (
	alphabet       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetLength = 26
)

// Key for polyalphabetic substitution ciphers.
type Key []simple.Key

// NewKeyFromAlphabets creates and returns a key with the same length as the
// number of alphabets using each alphabet for each underlying simple
// substitution key. The given alphabets should be a string containing every
// letter of the alphabet in upper case and each only once (i.e. a valid
// shuffling of the uppercase alphabet).
func NewKeyFromAlphabets(alphabets ...string) Key {
	key := make(Key, len(alphabets))
	for i, alphabet := range alphabets {
		key[i] = simple.NewKeyFromAlphabet(alphabet)
	}
	return key
}

// NewRandomKey creates and returns a key with the given length using a
// randomized alphabet order for each ciphertext alphabet.
func NewRandomKey(length int) Key {
	key := make(Key, length)
	for i := 0; i < length; i++ {
		key[i] = simple.NewRandomKey()
	}
	return key
}

// NewKeyFromKeywords creates and returns a key using the given keywords to
// generate a ciphertext alphabet that uses the keyword and fills in the
// remaning characters in alphabetical order. The keywords must not contain
// duplicate characters.
func NewKeyFromKeywords(keywords ...string) Key {
	key := make(Key, len(keywords))
	for i, keyword := range keywords {
		key[i] = simple.NewKeyFromKeyword(keyword)
	}
	return key
}

// Inverse returns the inverse key for the key. The inverse key can be used to
// decrypt messages encrypted with this key.
func (k Key) Inverse() Key {
	inverse := make(Key, len(k))
	for i, simpleKey := range k {
		inverse[i] = simpleKey.Inverse()
	}
	return inverse
}

func (k Key) String() string {
	simpleStrings := make([]string, len(k))
	for i, simpleKey := range k {
		simpleStrings[i] = simpleKey.String()
	}
	return strings.Join(simpleStrings, ",")
}

// Substitute each of the characters in the given message according to the given
// key. If a plaintext message is given, the result is ciphertext. If a
// ciphertext message is given, the key must be the inverse of the encryption
// and the result is plaintext.
//
// Capitalization and punctuation are preserved. For improved security, it is
// recommended to use a single casing (upper or lower), and remove all spaces
// and punctuation.
func Substitute(key Key, message string) string {
	messageRunes := []rune(message)
	keyIndex := 0
	result := make([]rune, 0, len(messageRunes))
	for _, c := range messageRunes {
		upperC := []rune(strings.ToUpper(string(c)))[0]
		isLower := c != upperC

		if nextC, ok := key[keyIndex][upperC]; ok {
			keyIndex = mod(keyIndex+1, len(key))
			if isLower {
				result = append(result, []rune(strings.ToLower(string(nextC)))[0])
			} else {
				result = append(result, nextC)
			}
			continue
		}
		result = append(result, c)
	}
	return string(result)
}

func mod(n, m int) int {
	return ((n % m) + m) % m
}
