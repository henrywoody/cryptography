package monoalpha

// NewAtbashKey creates and returns the key for the ROT13 Cipher (Caesar with a
// shift of 13).
func NewROT13Key() Key {
	return NewCaesarKey(13)
}

// NewCaesarKey creates and returns a key for a Caesar Cipher with a ciphertext
// alphabet formed by shifting the plaintext alphabet by the given shift
// distance.
func NewCaesarKey(shift int) Key {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	shift = mod(shift, len(alphabet))
	return NewKeyFromAlphabet(alphabet[shift:] + alphabet[:shift])
}

func mod(n, m int) int {
	return ((n % m) + m) % m
}
