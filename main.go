package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
)

func main() {

	options := "luns"
	length := 8

	if len(os.Args) == 3 {
		l, err := strconv.Atoi(os.Args[1])
		if err != nil {
			panic("First argument must be a number")
		}
		length = l
		options = os.Args[2]
	} else if len(os.Args) == 2 {
		arg1 := os.Args[1]
		
		conv, err := strconv.Atoi(arg1)

		if (err != nil) {
			options = arg1
		} else {
			length = conv
		}
	}
	fmt.Println(Generate(length, options))

}

func Generate(length int, options string) string {

	pass := ""

	chars := []string {}

	for i := 0; i < len(options); i++ {
		switch options[i] {
		case 'u':
			chars = append(chars, "ABCDEFGHJKMNPQRSTUVWXYZ")
		case 'l':
			chars = append(chars, "abcdefghijklmnopqrstuvwxyz")
		case 'n':
			chars = append(chars, "23456789")
		case 's':
			chars = append(chars, "!@#$%^&*()=+[]{};:,.<>?/`~")
		default:
			panic("Invalid option " + string(options[i]))
		}
	}

	for i := 0; i < length; i++ {
		charType := rand.IntN(len(chars))

		nextRand := rand.IntN(len(chars[charType]))

		char := chars[charType][nextRand]

		pass = pass + string(char)

	}

	return pass
}
