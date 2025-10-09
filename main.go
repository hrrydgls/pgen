package main

import (
	"fmt"
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

