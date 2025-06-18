package pkg

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateProejct(description, projectName string) {
	println("프로젝트 생성 - README.md:", description)
	fmt.Printf("프로젝트 이름: %s\n", projectName)

	// createReadme(input)
	
}

func createReadme(description string) error {
	filename := "README.md"

	outputDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	outputPath := filepath.Join(outputDir, "Downloads", filename)

	// TODO
	// README content by generating ai

	if err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm); err != nil {
		return err
	}

	return os.WriteFile(outputPath, []byte(description), 0644)
}

// TODO Add go project structure