package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"golang.org/x/net/idna"
)

// Main function, wrapping the realMain function for unit-test capability
func main() {
	os.Exit(realMain())
}

// realMain function, wrapped by the Main function, processes inpput STDIN and CLI and returns an integer indicating success/failure
func realMain() int {
	argsWithoutProg := os.Args[1:]

	var regularString string

	if len(argsWithoutProg) <= 0 {
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			regularString = scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Println(err)
			return 2
		}
	} else {
		regularString = os.Args[1] // we only take a single parameter, the string to decode
	}

	if regularString != "" {
		punycodeString, err := idna.ToASCII(regularString)

		if err == nil {
			fmt.Printf("%s\n", punycodeString)
			return 0
		}
	}

	return 1
}
