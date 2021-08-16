package main

import (
	"fmt"
	"os"

	"github.com/aerogear/charmil-plugin-example/internal/docs"
)

func main() {
	err := docs.CreateModularDocs()
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
