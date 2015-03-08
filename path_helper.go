package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Println("Usage: path_helper <file1>[, <file2>...]")
		return
	}

	var errors []error
	var output []string

	for i, arg := range os.Args {
		if i != 0 {
			content, err := ioutil.ReadFile(arg)
			if err != nil {
				errors = append(errors, err)
				continue
			}

			lines := strings.Split(string(content), "\n")

			for _, line := range lines {
				line := strings.TrimSpace(line)

				if len(line) == 0 {
					continue
				} else {
					output = append(output, line)
				}
			}
		}
	}

	if len(errors) > 0 {
		for _, err := range errors {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}
		os.Exit(1)
	} else {
		fmt.Print("PATH=\"")
		fmt.Print(strings.Join(output, ":"))
		fmt.Print("\"; export PATH\n")
	}
}
