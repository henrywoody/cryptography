package simple

// NewAtbashKey creates and returns the key for the Atbash cipher (reversed
// alphabet).
func NewAtbashKey() Key {
	return NewKeyFromAlphabet("ZYXWVUTSRQPONMLKJIHGFEDCBA")
}
