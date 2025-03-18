package main

import (
	"fmt"
	"os"
)

func main() {
	app := NewApp()

	err := app.Start()

	if err != nil {
		_, err = fmt.Fprintln(os.Stderr, err)

		if err != nil {
			os.Exit(1)
		}
		os.Exit(1)
	}

	os.Exit(0)
}
