package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/henrywoody/cryptography/substitution/monoalpha"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Must specify a command argument")
		return
	}

	switch c := os.Args[1]; c {
	case "key":
		key()
	case "enc":
		enc()
	case "dec":
		dec()
	default:
		log.Fatalf("Command '%s' not recognized\n", c)
	}
}

func key() {
	input := ""
	if len(os.Args) >= 3 {
		input = os.Args[2]
	}
	key := keyFromInput(input)
	fmt.Println(key)
}

func enc() {
	keyInput := ""
	if len(os.Args) >= 3 {
		keyInput = os.Args[2]
	}
	key := keyFromInput(keyInput)

	message := ""
	if len(os.Args) >= 4 {
		message = os.Args[3]
	}

	ciphertext := monoalpha.Substitute(key, message)
	fmt.Println(ciphertext)
}

func dec() {
	keyInput := ""
	if len(os.Args) >= 3 {
		keyInput = os.Args[2]
	}
	key := keyFromInput(keyInput).Inverse()

	message := ""
	if len(os.Args) >= 4 {
		message = os.Args[3]
	}

	plaintext := monoalpha.Substitute(key, message)
	fmt.Println(plaintext)
}

var affineKeyRe = regexp.MustCompile(`^(-?\d+),(-?\d+)$`)

func keyFromInput(input string) monoalpha.Key {
	if input == "" {
		return monoalpha.NewRandomKey()
	}
	if strings.ToLower(input) == "atbash" {
		return monoalpha.NewAtbashKey()
	}
	if shift, err := strconv.Atoi(input); err == nil {
		return monoalpha.NewCaesarKey(shift)
	}
	if match := affineKeyRe.FindStringSubmatch(input); len(match) > 0 {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		return monoalpha.NewAffineKey(a, b)
	}
	return monoalpha.NewKeyFromKeyword(input)
}
