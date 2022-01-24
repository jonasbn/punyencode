package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/idna"
)

func main() {
	argsWithoutProg := os.Args[1:]

	var regularString string

	if len(argsWithoutProg) <= 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			regularString = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
			os.Exit(1)
		}
	} else {
		regularString = os.Args[1] // we only take a single parameter, the string to decode
	}

	punycodeString, err := idna.ToASCII(regularString)

	if err == nil {
		fmt.Printf("%s\n", punycodeString)
	}
}
