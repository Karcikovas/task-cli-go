package main

import (
	"fmt"
	"os"
)

func main() {
	app := NewApp()

	err := app.Start()

	if err != nil {
		_, err = fmt.Fprintf(os.Stderr, err.Error())

		if err != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}

	os.Exit(0)
}
