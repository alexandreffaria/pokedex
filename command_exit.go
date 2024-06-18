package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Bye bye!")
	os.Exit(0)
	return nil
}
