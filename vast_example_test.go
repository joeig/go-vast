package vast_test

import (
	"fmt"
	"go.eigsys.de/go-vast"
	"log"
)

func ExampleNew() {
	example := vast.New()
	example.Version = vast.VAST42Version
	// Output:
}

var exampleReader = mustOpenFixture("iab/Inline_Simple.xml")

func ExampleRead() {
	example, err := vast.Read(exampleReader)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%s", example.Version)
	// Output: 4.2
}

func ExampleVAST_Bytes() {
	example := vast.New()

	exampleBytes, err := example.Bytes()
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Printf("%s", exampleBytes)
	// Output: <?xml version="1.0" encoding="UTF-8"?>
	// <VAST version="4.2" xmlns="http://www.iab.com/VAST"></VAST>
}
