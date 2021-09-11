package main

import (
	"fmt"
	"os"
	"time"
)

func GenerateGoCode(args []string) (dir, file, source string) {
	numberOfArgs := len(args)
	if numberOfArgs < 4 {
		fmt.Println(`
Usage: esg go output_dir pkg_name error_code formatted_message [name_of_arguments..]
Example: esg go . errors InvalidPhone "%s is not valid phone number." Phone
`)
		os.Exit(0)
	}

	dir = args[0]
	pkg := args[1]
	errCode := args[2]
	msg := args[3]
	var formatArgs []string
	haveArgs := numberOfArgs > 4
	if haveArgs {
		formatArgs = args[4:]
	}

	// write struct
	lastIdx := len(formatArgs) - 1
	source = fmt.Sprintf(`// Package errors
// Generated by ESG at %s. github.com/simplefelix/esg
package %s

import "fmt"

type %s struct {
`, time.Now().Format("2006-01-02 15:04:05"),
		pkg, errCode)
	if haveArgs {
		for _, arg := range formatArgs {
			source += fmt.Sprintf("	%s interface{}", arg)
			source += "\n"
		}
	}
	source += "}\n"

	// write Code() function
	source += fmt.Sprintf(`
func (e %s)Code() string {
	return "%s"
}
`, errCode, errCode)

	// write Error() function
	source += fmt.Sprintf("\nfunc (e %s)Error() string {\n	return fmt.Sprintf(\"%s\"", errCode, msg)
	if haveArgs {
		for _, arg := range formatArgs {
			source += fmt.Sprintf(", e.%s", arg)
		}
	}
	source += ")\n}\n"

	// write New() function
	source += fmt.Sprintf(`
func New%s(`, errCode)
	if haveArgs {
		for idx, arg := range formatArgs {
			source += fmt.Sprintf("%s interface{}", arg)
			if idx != lastIdx {
				source += ", "
			}
		}
	}
	source += fmt.Sprintf(`) %s {
	return %s{
`, errCode, errCode)
	if haveArgs {
		for _, arg := range formatArgs {
			source += fmt.Sprintf("		%s: %s,\n", arg, arg)
		}
	}
	source += "	}\n}\n"

	return dir, errCode + ".go", source
}