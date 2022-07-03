package monoalpha

// NewAtbashKey creates and returns the key for the Atbash Cipher (reversed
// alphabet).
func NewAtbashKey() Key {
	return NewKeyFromAlphabet("ZYXWVUTSRQPONMLKJIHGFEDCBA")
}
