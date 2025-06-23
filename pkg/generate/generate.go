package pkg

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rlaope/ainit/pkg/ai"
)

const gitignoreTemplate = `
/bin/
/vendor/
/.env
*.log
`

const gitattributesTemplate = `
* text=auto
`

const makefileTemplate = `
APP_NAME = app

build:
	go build -o bin/$(APP_NAME) ./cmd

run:
	go run ./cmd/main.go

test:
	go test ./...

tidy:
	go mod tidy

clean:
	rm -rf bin/
`

func GenerateProject(description, projectName string) {
	fmt.Println("Creating project - README.md from:", description)
	fmt.Printf("Project name: %s\n", projectName)

	projectDir, err := getProjectPath(projectName)
	if err != nil {
		fmt.Printf("Failed to get project path: %v\n", err)
		os.Exit(1)
	}

	if err := os.MkdirAll(filepath.Join(projectDir, "cmd"), 0755); err != nil {
		fmt.Printf("Failed to create project directory: %v\n", err)
		os.Exit(1)
	}

	if err := generateReadmeFile(description, projectDir); err != nil {
		fmt.Printf("Failed to generate README.md: %v\n", err)
		os.Exit(1)
	}

	if err := generateTemplateFiles(projectDir, projectName); err != nil {
		fmt.Printf("Failed to generate template files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Project created successfully:", projectName)
}

func generateReadmeFile(description, projectDir string) error {
	content, err := ai.GenerateReadme(description)
	if err != nil {
		return err
	}
	readmePath := filepath.Join(projectDir, "README.md")
	return os.WriteFile(readmePath, []byte(content), 0644)
}

func generateTemplateFiles(projectDir, projectName string) error {
	if err := os.WriteFile(filepath.Join(projectDir, ".gitignore"), []byte(gitignoreTemplate), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(projectDir, ".gitattributes"), []byte(gitattributesTemplate), 0644); err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(projectDir, "Makefile"), []byte(makefileTemplate), 0644); err != nil {
		return err
	}

	mainContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, ` + projectName + `!")
}`
	if err := os.WriteFile(filepath.Join(projectDir, "cmd", "main.go"), []byte(mainContent), 0644); err != nil {
		return err
	}

	cmd := exec.Command("go", "mod", "init", projectName)
	cmd.Dir = projectDir
	return cmd.Run()
}

func getProjectPath(projectName string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Downloads", projectName), nil
}
