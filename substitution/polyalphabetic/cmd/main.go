package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/henrywoody/cryptography/substitution/polyalphabetic"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cmd := &Command{}
	cmd.parseArgs()
	cmd.run()
}

type Command struct {
	randomKeyLength  int
	vigenereKeyword  string
	gronsfeldKeyword string
	keywords         string
	message          string
}

func (c *Command) run() {
	if len(os.Args) < 2 {
		log.Fatal("Must specify a command argument")
	}

	switch cmdName := os.Args[1]; cmdName {
	case "key":
		c.key()
	case "enc":
		c.enc()
	case "dec":
		c.dec()
	default:
		log.Fatalf("Command '%s' not recognized\n", cmdName)
	}
}

func (c *Command) parseArgs() {
	keyType := ""
	if len(os.Args) >= 3 {
		keyType = os.Args[2]
	}

	keyArg := ""
	if len(os.Args) >= 4 {
		keyArg = os.Args[3]
	}

	c.message = ""
	if len(os.Args) >= 5 {
		c.message = os.Args[4]
	}

	switch keyType {
	case "l":
		c.randomKeyLength, _ = strconv.Atoi(keyArg)
	case "v":
		c.vigenereKeyword = keyArg
	case "g":
		c.gronsfeldKeyword = keyArg
	case "k":
		c.keywords = keyArg
	}
}

func (c *Command) key() {
	key := c.keyFromArgs()
	fmt.Println(key)
}

func (c *Command) enc() {
	key := c.keyFromArgs()

	ciphertext := polyalphabetic.Substitute(key, c.message)
	fmt.Println(ciphertext)
}

func (c *Command) dec() {
	key := c.keyFromArgs().Inverse()

	plaintext := polyalphabetic.Substitute(key, c.message)
	fmt.Println(plaintext)
}

func (c *Command) keyFromArgs() polyalphabetic.Key {
	if c.randomKeyLength > 0 {
		return polyalphabetic.NewRandomKey(c.randomKeyLength)
	}
	if c.vigenereKeyword != "" {
		return polyalphabetic.NewVigenereKey(c.vigenereKeyword)
	}
	if c.gronsfeldKeyword != "" {
		return polyalphabetic.NewGronsfeldKey(c.gronsfeldKeyword)
	}
	if c.keywords != "" {
		return polyalphabetic.NewKeyFromKeywords(strings.Split(c.keywords, ",")...)
	}
	return polyalphabetic.NewRandomKey(3)
}
