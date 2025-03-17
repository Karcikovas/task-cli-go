package main

import (
	"fmt"
	"os"
)

func main() {
	app := NewApp()

	err := app.Start()

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
