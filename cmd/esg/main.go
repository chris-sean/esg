package main

import (
	"fmt"
	"os"

	"github.com/SimpleFelix/esg/internal"
)

func main() {
	argsWithProg := os.Args
	if len(argsWithProg) < 2 {
		showHelp()
	}

	lang := argsWithProg[1]
	args := argsWithProg[2:]

	// generate and save file
	switch lang {
	case "go":
		save(internal.GenerateGoCode(args))
	default:
		showHelp()
	}
}

func save(outputDir, filename, source string) {
	fmt.Printf("Generate Code:\n\n%s\n", source)
	_ = os.MkdirAll(outputDir, 0666)
	filepath := fmt.Sprintf("%s/%s", outputDir, filename)
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(source))
	if err != nil {
		panic(err)
	}
	if err = f.Close(); err != nil {
		panic(err)
	}
	fmt.Printf("\nFile Path: %s", filepath)
}

func showHelp() {
	fmt.Println(`Error Struct Generator

Usage: esg language arguments

Command:
go		Generate Go code
`)
	os.Exit(0)
}
