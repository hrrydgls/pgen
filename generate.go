package main 

import (
	"math/rand/v2"
)


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