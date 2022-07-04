package monoalpha

import (
	"math/rand"
	"strings"
)

const (
	alphabet       = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetLength = 26
)

// Key for subsitution ciphers.
type Key map[rune]rune

// NewKeyFromAlphabet creates and returns a key with the given alphabet order.
// The given alphabet should be a string containing every letter of the alphabet
// in upper case and each only once (i.e. a valid shuffling of the uppercase
// alphabet).
func NewKeyFromAlphabet(alphabet string) Key {
	key := make(Key, alphabetLength)
	for i, c := range []rune(alphabet) {
		key['A'+rune(i)] = c
	}
	return key
}

// NewRandomKey creates and returns a key with a randomized alphabet order.
func NewRandomKey() Key {
	alphabet := make([]rune, 0, alphabetLength)
	for c := 'A'; c <= 'Z'; c++ {
		alphabet = append(alphabet, c)
	}
	rand.Shuffle(len(alphabet), func(i, j int) {
		alphabet[i], alphabet[j] = alphabet[j], alphabet[i]
	})
	return NewKeyFromAlphabet(string(alphabet))
}

// NewKeyFromKeyword creates and returns a key using the given keyword and
// filling in the remaning characters in alphabetical order. The key must not
// contain duplicate characters.
func NewKeyFromKeyword(keyword string) Key {
	keywordChars := []rune(strings.ToUpper(keyword))
	key := make(Key, alphabetLength)
	usedChars := make(map[rune]bool, len(keywordChars))

	nextC := 'A'
	for i := 0; i < alphabetLength; i++ {
		if i < len(keywordChars) {
			key['A'+rune(i)] = keywordChars[i]
			usedChars[keywordChars[i]] = true
			continue
		}
		for usedChars[nextC] {
			nextC++
		}
		key['A'+rune(i)] = nextC
		nextC++
	}

	return key
}

// Inverse returns the inverse key for the key. The inverse key can be used to
// decrypt messages encrypted with this key.
func (k Key) Inverse() Key {
	inverse := make(map[rune]rune, len(k))
	for i, o := range k {
		inverse[o] = i
	}
	return inverse
}

func (k Key) String() string {
	chars := make([]rune, 0, len(k))
	for c := 'A'; c <= 'Z'; c++ {
		chars = append(chars, k[c])
	}
	return string(chars)
}

// Substitute each of the characters in the given message according to the given
// key. If a plaintext message is given, the result is ciphertext. If a
// ciphertext message is given, the key must be the inverse of the encryption
// key and the result is plaintext.
//
// Capitalization and punctuation are preserved. For additional security, it is
// recommended to use a single casing (upper or lower), and remove all spaces
// and punctuation.
func Substitute(key Key, message string) string {
	result := make([]rune, 0, len(message))
	for _, c := range []rune(message) {
		upperC := []rune(strings.ToUpper(string(c)))[0]
		isLower := c != upperC

		if nextC, ok := key[upperC]; ok {
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
