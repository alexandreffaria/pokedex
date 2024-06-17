package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Command EXIT")
	os.Exit(0)
	return nil
}
