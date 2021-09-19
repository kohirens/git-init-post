package main

import (
	"fmt"
	"log"
	"os"
)

const (
	PS       = string(os.PathSeparator)
)

func init() {

}

func main() {
	var mainErr error

	defer func() {
		if mainErr != nil {
			fmt.Print("\nfatal error detected: ")
			log.Fatalln(mainErr)
		}
		os.Exit(0)
	}()
}