package simple

// NewAffineKey creates and returns a key for an Affine cipher. Note that the
// value of a must be coprime with the length of the alphabet (26), and
// therefore cannot be even (multiple of 2) or 13 (mod 26).
func NewAffineKey(a, b int) Key {
	alphabetRunes := []rune(alphabet)
	cipherAlphabet := make([]rune, len(alphabetRunes))
	for i := 0; i < len(alphabetRunes); i++ {
		cipherAlphabet[i] = alphabetRunes[mod(a*i+b, len(alphabetRunes))]
	}
	return NewKeyFromAlphabet(string(cipherAlphabet))
}

func mod(n, m int) int {
	return ((n % m) + m) % m
}
