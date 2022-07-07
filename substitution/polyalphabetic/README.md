# Polyalphabetic Substitution

Polyalphabetic substitution ciphers build upon the simple, monoalphabetic ciphers but use multiple ciphertext alphabets for encryption with improved security.

In a polyalphabetic substitution cipher, a series of ciphertext alphabets are selected and used to encipher each character in the plaintext message. The ciphertext alphabet used to encrypt a given character is selected according to the position, modulo the number of ciphertext alphabets, of the character. For example, when using three ciphertext alphabets, the first character is enciphered using the first alphabet, the second using the second, the third using the third, the fourth using the first, and so on.

## Example

Plaintext alphabet: ABCDEFGHIJKLMNOPQRSTUVWXYZ

Ciphertext alphabets:

1. MXPVEUFBKYHQDTRJICOSZGAWNL
2. EYLJBFRPOADVTXHMNIWKCZQSGU
3. CQTHGMRYIXOLBFSKPUJNZEWDAV

Plaintext message: HELLO WORLD

Ciphertext message: BBLQH WRILV

## Named Ciphers

Some cases of polyalphabetic substitution have special names depending on the key or the process for creating/deriving the key.

A Vigenère cipher is a polyalphabetic substitution cipher in which the key is a set of Caesar ciphertext alphabets (each shifted by some distance).

A Gronsfeld cipher is similar to a Vigenère cipher but uses a numeric key instead of an alphabetic key. This restricts the ciphertext alphabet space from containing 26 possibilities to 10, but encourages harder to guess keys since the values are entirely numeric (instead of alphabetic, which encourages words as keys, which are less random).

## Keys

The key for a polyalphabetic substitution cipher can be represented by a list [monoalphabetic substitution keys](../simple/README.md#Keys).

The key for a Vigenère cipher are usually encoded as a keyword where each letter of the keyword is used to denote the shift distance (determined by the letter's distance to 'A') for the ciphertext alphabet. For example, the keyword KEY represents three ciphertext alphabets where the first is shifted by 10, the second is shifted by 4, and the last by 24, and would correspond to the full key: KLMNOPQRSTUVWXYZABCDEFGHIJ, EFGHIJKLMNOPQRSTUVWXYZABCD, YZABCDEFGHIJKLMNOPQRSTUVWX.

## Security

Polyalphabetic substitution ciphers are considered much more secure than monoalphabetic substitution ciphers. Polyalphabetic substitution ciphers can be weakened if the sender does not follow the restrictions of the algorithm or uses phrases that an attacker can guess will be included in the message (e.g. "Heil Hitler"). Polyalphabetic substitution ciphers become more suseptible to analysis when the message becomes much longer than the key (longer than 27 times the length of the key).

Vigenère ciphers reduce the entropy of the key by reducing the number of possible ciphertext alphabets from 26! (or !26) to 26 (or 25) and are thus less secure relative to polyalphabetic substitution ciphers that use mixed alphabets (randomly generated).

## Further Reading

- https://en.wikipedia.org/wiki/Substitution_cipher
- [https://en.wikipedia.org/wiki/Vigenère_cipher](https://en.wikipedia.org/wiki/Vigenère_cipher)

