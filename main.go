package main

import (
	"fmt"
	"os"
)

func main() {
	app, err := NewApp()

	if err != nil {
		panic(err)
	}

	err = app.Start()

	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}
