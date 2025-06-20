package pkg

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/rlaope/ainit/pkg/ai"
)

func GenerateProejct(description, projectName string) {
	println("프로젝트 생성 - README.md:", description)
	fmt.Printf("프로젝트 이름: %s\n", projectName)

	createReadme(description)
}

func createReadme(description string) error {
	content, err := ai.GenerateReadme(description)

	if err != nil {
		return err
	}

	readmePath := filepath.Join(".", "README.md") // TODO add filepath on cli
	if err := os.WriteFile(readmePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("README 작성 실패: %w", err)
	}

	fmt.Println("✅ README.md 생성 완료")
	return nil
}

// TODO Add go project structure