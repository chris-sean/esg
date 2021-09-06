package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	argsWithProg := os.Args
	numberOfArgs := len(argsWithProg) - 1
	if numberOfArgs < 4 {
		log.Fatalln(`
Usage: esg output_dir pkg_name error_code formatted_message [name_of_arguments..]
name_of_argument must not be ErrorCode.
Example: esg /src/myproj errors InvalidPhone "%s is not valid phone number." Phone
`)
	}

	outputDir := argsWithProg[1]
	pkg := argsWithProg[2]
	errCode := argsWithProg[3]
	msg := argsWithProg[4]
	args := []string{}
	haveArgs := numberOfArgs > 4
	if haveArgs {
		args = argsWithProg[5:]
	}

	goCode := fmt.Sprintf(`package %s

import "fmt"

type %s struct {
	ErrorCode string
`, pkg, errCode)
	if haveArgs {
		lastIdx := len(args) - 1
		for idx, arg := range args {
			goCode += fmt.Sprintf("	%s interface{}", arg)
			if idx != lastIdx {
				goCode += "\n"
			}
		}
	}
	goCode += `
}

`

	goCode += fmt.Sprintf("func (e %s)Error() string {\n	return fmt.Sprintf(\"%s\"", errCode, msg)

	if haveArgs {
		for _, arg := range args {
			goCode += fmt.Sprintf(", e.%s", arg)
		}
	}
	goCode += ")\n}"

	fmt.Printf("Generate Code:\n\n%s\n", goCode)

	// save file
	save(outputDir, pkg, errCode, goCode)
}

func save(outputDir, pkg, errCode, goCode string) {
	path := fmt.Sprintf("%s/%s", outputDir, pkg)
	_ = os.MkdirAll(path, 0666)
	filepath := fmt.Sprintf("%s/%s.go", path, errCode)
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		panic(err)
	}
	_, err = f.Write([]byte(goCode))
	if err != nil {
		panic(err)
	}
	if err = f.Close(); err != nil {
		panic(err)
	}
	fmt.Printf("\nFile P/ath: %s", filepath)
}


