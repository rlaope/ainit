package main

import (
	"fmt"
	"os"

	"github.com/rlaope/ainit/pkg"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <description> <project_name>")
		os.Exit(1)
	}

	description := os.Args[1]
	projectName := os.Args[2]

	pkg.GenerateProejct(description, projectName)

	fmt.Printf("Project '%s' generated successfully.\n", projectName)
}