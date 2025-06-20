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

func GenerateProejct(description, projectName string) {
	println("프로젝트 생성 - README.md:", description)
	fmt.Printf("프로젝트 이름: %s\n", projectName)

	err := createReadme(description, projectName)
	if err != nil {
		fmt.Printf("README.md 생성 실패: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("프로젝트 생성 완료:", projectName)
}

func createReadme(description string, name string) error {
	content, err := ai.GenerateReadme(description)

	if err != nil {
		return err
	}

	projectDir, err := getProjectPath(name)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Join(projectDir, "cmd"), 0755); err != nil {
		return fmt.Errorf("디렉토리 생성 실패: %w", err)
	}
	_ = os.WriteFile(projectDir, []byte(content), 0644)
	_ = os.WriteFile(filepath.Join(name, ".gitignore"), []byte(gitignoreTemplate), 0644)
	_ = os.WriteFile(filepath.Join(name, ".gitattributes"), []byte(gitattributesTemplate), 0644)

	_ = os.WriteFile(filepath.Join(name, "Makefile"), []byte(makefileTemplate), 0644)

	mainContent := `package main

import "fmt"

func main() {
	fmt.Println("Hello, ` + name + `!")
}`
	_ = os.WriteFile(filepath.Join(name, "cmd", "main.go"), []byte(mainContent), 0644)

	cmd := exec.Command("go", "mod", "init", name)
	cmd.Dir = name
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("go mod init 실패: %w", err)
	}

	fmt.Println("✅ 프로젝트 생성 완료:", name)
	return nil
}

func getProjectPath(projectName string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	projectPath := filepath.Join(home, "Downloads", projectName)
	return projectPath, nil
}