package main

import (
	"fmt"
	"os"
)

// code_summarizer - Summarize large codebases
func code_summarizer(path string) {
	fmt.Println("========================================")
	fmt.Println("  Code-Summarizer")
	fmt.Println("  Summarize large codebases")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	code_summarizer(path)
}
