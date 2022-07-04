# Monoalphabetic Substitution

Monoalphabetic substitution ciphers use two alphabets, a plaintext alphabet and a ciphertext alphabet. Each individual character from the plaintext alphabet is mapped to a character in the ciphertext alphabet. The two alphabets are typically composed of the same characters (e.g. the Latin alphabet) but in differing orders.

## Example

Plaintext alphabet: ABCDEFGHIJKLMNOPQRSTUVWXYZ

Ciphertext alphabet: MXPVEUFBKYHQDTRJICOSZGAWNL

Plaintext message: HELLO WORLD

Ciphertext message: BEQQR ARCQV

## Named Ciphers

Some cases of monoalphabetic substitution have special names depending on the key.

An Affine cipher is a monoalphabetic substitution cipher in which the ciphertext alphabet contains the same characters as the plaintext alphabet where the plaintext characters are first assigned sequential numbers starting with zero and then mapped to a ciphertext character according to the function `ax + b` where `x` is the numerical value (or index) of the character and `a` and `b` are parameters that form the encryption key.

A Caesar cipher is a monoalphabetic substitution cipher in which the ciphertext alphabet contains the same characters as the plaintext alphabet in the same order but shifted by a fixed distance (e.g. DEFGHIJKLMNOPQRSTUVWXYZABC), there are 26 possible Caesar cipher alphabets (including the trivial case). A Caesar cipher is a special case of the Affine cipher where the value of `a` is 1.

ROT13 (rotate by 13 places) is a Caesar cipher with a shift of 13. ROT13 is commonly used due to the convenient property that the encryption key and decription key are the same, which is a unique property among the Caesar ciphers.

The Atbash cipher is a monoalphabetic substituion cipher in which the chipertext alphabet contains the same characters as the plaintext alphabet in reverse order (e.g. ZYXWVUTSRQPONMLKJIHGFEDCBA). The Atbash cipher is a special case of the Affine cipher where `a` is -1 and `b` is -1.

## Keys

The key for a monoalphabetic substitution cipher can be represented by pairs of characters where each pair contains one character from the plaintext alphabet and one character from the ciphertext alphabet such that all characters from each alphabet are used once and only once. For example, using the example above, the key can be represented as the pairs: (A, M), (B, X), (C, P), ..., (Z, L). To generate a random key, use the `NewRandomKey` function.

For alphabets with an order (like the latin alphabet), the key can be represented as just the ciphertext alphabet alone, where the character position indicates the character from the plaintext alphabet.

Keywords can be used to generate a key in order to simplify expression of the key (e.g. to make it easier to remember). To generate a ciphertext alphabet from a keyword, start the alphabet with the keyword and then add the remaining letters of the alphabet in order. For example, the keyword CLOUD produces the ciphertext alphabet CLOUDABEFGHIJKMNPQRSTVWXYZ. Note that all letters after the last letter (alphabetically) of the keyword map to the same letter in the ciphertext alphabet, so including Z in the keyword is recommended so that all letters are encrypted. Use the `NewKeyFromKeyword` function to generate a key in this fashion.

Keys for Affine ciphers can be expressed more simply by a pair if integers for the values `a` and `b` used in the substitution function. For example a key of (3, 1) would represent the ciphertext alphabet BEHKNQTWZCFILORUXADGJMPSVY. Note that in order to be decodable, the value of `a` must be relatively prime (coprime) with the length of the alphabet. In the case of the Latin alphabet (length 26) the value of `a` cannot be even or 13 (mod 26).

Keys for the Caesar cipher can be expressed even more simply with a single integer between 0 and 25 to indicate the distance of the shift (or the `b` value of an Affine cipher). For example, a key of 3 would represent the ciphertext alphabet DEFGHIJKLMNOPQRSTUVWXYZABC. Use the `NewCaesarKey` function to generate a Caesar cipher key or `NewROT13Key` to generate the ROT13 key.

The Atbash cipher uses a single key for encryption so agreeing on the Atbash cipher is enough to communicate the key. Use the `NewAtbashKey` function to generate the Atbash cipher key.

### Encryption vs Decryption

Encryption and decryption use the same process (and same `Substitute` function in the code) but with different keys. For decryption, the alphabets used for encryption are swapped so that the encryption's plaintext alphabet becomes the decription's ciphertext alphabet and vice versa. To obtain the inverse key for a key in the code, use the `Inverse` method on the key.

## Security

Monoalphabetic substitution ciphers are easy to crack and are no longer used for serious encryption needs. When encrypting using a simple substitution cipher, it is recommended to use a single case and omit spacing and punctuation to offer fewer hints to an attacker.

The Atbash cipher can be cracked instantly if the attacker knows it is being used as there is only one key. Considering [Kerckhoffs's principle](https://en.wikipedia.org/wiki/Kerckhoffs%27s_principle), which states that a cryptosystem should be secure even when everything about the system (other than the key) is public knowledge, the Atbash cipher is totally insecure since the key is known automatically.

The Caesar cipher has 25 (non-trivial) possible keys and can easily be enumerated. The ROT13 cipher, similar to the Atbash cipher, has only one key and can therefore be cracked instantly if ROT13 is known to be used. Caesar ciphers, and ROT13 in particular, are often used to obfuscate text so that it cannot be understood at a glance but can easily be understood if desired, similar to writing upside down. This is useful for hiding spoilers or answers (e.g. to a riddle) on a page and letting the reader decide if they want to read the message.

Affine ciphers have 311 (non-trivial) possible keys and can still be enumerated quite easily, especially with a computer.

Using a keyword to generate a ciphertext alphabet reduces the entropy of the key, making the system less secure than randomly ordering the alphabet. With a fully randomly ordered ciphertext alphabet, there are 26! (about 2^88) possible keys or !26 (about 2^87) where no characters remain in the same position, making enumeration infeasible. Instead character frequency or pattern analysis can be used to crack the message. For example, in English, the most common letters are E and T. Also certain patterns are common in english such as the pairs CH, SH, SS, and EA.

## Further Reading

- http://abstract.ups.edu/aata/section-private-key-crypt.html
- https://en.wikipedia.org/wiki/Substitution_cipher