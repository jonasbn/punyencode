package main

import (
	"fmt"
	"os"

	"golang.org/x/net/idna"
)

func main() {
	argStr := os.Args[1] // we only take a single parameter and

	asciiStr, err := idna.ToASCII(argStr)

	if err == nil {
		fmt.Printf("%s\n", asciiStr)
	}
}
